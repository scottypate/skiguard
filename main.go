package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/scalecraft/skiguard/internal/api/alert"
	"github.com/scalecraft/skiguard/internal/api/delete"
	"github.com/scalecraft/skiguard/internal/api/health"
	"github.com/scalecraft/skiguard/internal/api/load"
	"github.com/scalecraft/skiguard/internal/api/truncate"
	"github.com/scalecraft/skiguard/internal/api/update"
	"github.com/scalecraft/skiguard/internal/api/validate"
	"github.com/scalecraft/skiguard/internal/config"
	"github.com/scalecraft/skiguard/internal/duckdb"
	"github.com/scalecraft/skiguard/internal/slack"
	"github.com/scalecraft/skiguard/internal/snowflake"
)

func dbMigration() {
	duckdb.RunMigrations("db/schema/up")
}

func initialLoad(cfg *config.Config) {
	// Delete all data from duckdb on startup
	slog.Debug("deleting all data from duckdb")
	_, err := delete.Delete()
	if err != nil {
		log.Fatalf("error deleting all data from duckdb: %v", err)
	}

	// Load data from snowflake on startup and fail if unable to load
	slog.Info("loading data from snowflake")
	_, err = load.DataLoad(load.PostHandlerRequest{Cfg: cfg})

	if err != nil {
		log.Fatalf("error loading data from snowflake: %v", err)
	}
}

func main() {
	dbMigration()
	err := godotenv.Load()
	if err != nil {
		slog.Debug("no .env file found. proceeding with existing environment variables")
	}
	cfg := config.GetConfig()
	if err := config.ValidateLicenseKey(cfg.LicenseKey); err != nil {
		log.Fatalf("error validating license key: %v", err)
	}

	gin.SetMode(cfg.GinMode)

	_, err = snowflake.Connect(cfg.SnowflakeDSN)

	if err != nil {
		log.Fatalf("error connecting to snowflake: %v", err)
	}

	if cfg.SlackToken == "" {
		slog.Debug("slack token is not set. proceeding without slack integration")
	} else {
		if cfg.SlackChannelId == "" {
			log.Fatal("slack channel is required when slack token is set")
		}
		err := slack.AuthVerity(cfg.SlackToken)

		if err != nil {
			log.Fatalf("error authenticating to slack: %v", err)
		}
	}

	initialLoad(cfg)

	r := gin.Default()
	r.Use(gin.Recovery())
	err = r.SetTrustedProxies(nil)

	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	r.GET("/health", health.GetHandler)
	r.POST("/load", load.PostHandler(cfg))
	r.POST("/update", update.PostHandler(cfg))
	r.POST("/alert", alert.PostHandler(cfg))
	r.POST("/truncate", truncate.PostHandler())
	r.DELETE("/delete", delete.DeleteHandler())
	r.GET("/validate", validate.GetHandler(cfg))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HttpPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	slog.Info("skiguard application started successfully.")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	signal := <-shutdown

	slog.Info(fmt.Sprintf("Received signal %v, shutting down server", signal))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error(fmt.Sprintf("Error shutting down server: %v", err))
	}
}
