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
	"github.com/scalecraft/snowguard/internal/api/alert"
	"github.com/scalecraft/snowguard/internal/api/delete"
	"github.com/scalecraft/snowguard/internal/api/health"
	"github.com/scalecraft/snowguard/internal/api/load"
	"github.com/scalecraft/snowguard/internal/api/truncate"
	"github.com/scalecraft/snowguard/internal/api/update"
	"github.com/scalecraft/snowguard/internal/config"
	"github.com/scalecraft/snowguard/internal/duckdb"
	"github.com/scalecraft/snowguard/internal/slack"
	"github.com/scalecraft/snowguard/internal/snowflake"
)

func dbMigration() {
	duckdb.RunMigrations("db/schema/up")
}

func initialLoad(cfg *config.Config) {
	// Delete all data from duckdb on startup
	slog.Debug("deleting all data from duckdb")
	_, err := delete.Delete()
	if err != nil {
		slog.Error(fmt.Sprintf("error deleting all data from duckdb: %v", err))
		panic(err)
	}

	// Load data from snowflake on startup and fail if unable to load
	slog.Info("loading data from snowflake")
	_, err = load.DataLoad(load.PostHandlerRequest{Cfg: cfg})

	if err != nil {
		slog.Error(fmt.Sprintf("error loading data from snowflake: %v", err))
		panic(err)
	}
}

func main() {
	dbMigration()
	err := godotenv.Load()
	if err != nil {
		slog.Debug("no .env file found. proceeding with existing environment variables")
	}
	cfg := config.GetConfig()
	gin.SetMode(cfg.GinMode)

	_, err = snowflake.Connect(cfg.SnowflakeDSN)

	if err != nil {
		panic(err)
	}

	if cfg.SlackToken == "" {
		slog.Debug("slack token is not set. proceeding without slack integration")
	} else {
		if cfg.SlackChannelId == "" {
			panic("slack channel is required when slack token is set")
		}
		err := slack.AuthVerity(cfg.SlackToken)

		if err != nil {
			panic(fmt.Sprintf("error sending welcome message to slack channel: %v", err))
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HttpPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	slog.Info(fmt.Sprintf("Server started on port %d", cfg.HttpPort))

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
