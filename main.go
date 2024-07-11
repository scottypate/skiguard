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
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/scalecraft/snowguard/internal/api/delete"
	"github.com/scalecraft/snowguard/internal/api/health"
	"github.com/scalecraft/snowguard/internal/api/load"
	"github.com/scalecraft/snowguard/internal/api/update"
	"github.com/scalecraft/snowguard/internal/duckdb"
	"github.com/scalecraft/snowguard/internal/snowflake"
)

type config struct {
	snowflakeAccount   string
	snowflakeUser      string
	snowflakePassword  string
	snowflakeWarehouse string
	snowflakeRole      string
	slackToken         string
	snowflakeDSN       string
	httpPort           int
	shutdownTimeout    time.Duration
}

func dbMigration() {
	duckdb.RunMigrations("db/schema/down")
	duckdb.RunMigrations("db/schema/up")
}

func main() {
	dbMigration()
	err := godotenv.Load()
	if err != nil {
		slog.Debug("no .env file found. proceeding with existing environment variables")
	}
	cfg := getConfig()

	snowflakeDB, err := snowflake.Connect(cfg.snowflakeDSN)

	if err != nil {
		panic(err)
	}

	if cfg.slackToken == "" {
		slog.Debug("slack token is not set. proceeding without slack integration")
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	err = r.SetTrustedProxies(nil)

	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	r.GET("/health", health.GetHandler)
	r.POST("/load", load.PostHandler(snowflakeDB))
	r.POST("/update", update.PostHandler(snowflakeDB))
	r.DELETE("/delete", delete.DeleteHandler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.httpPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	slog.Info(fmt.Sprintf("Server started on port %d", cfg.httpPort))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	signal := <-shutdown

	slog.Info(fmt.Sprintf("Received signal %v, shutting down server", signal))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error(fmt.Sprintf("Error shutting down server: %v", err))
	}

}

func getConfig() config {
	c := config{
		snowflakeAccount:   getEnv("SNOWFLAKE_ACCOUNT", "", true),
		snowflakeUser:      getEnv("SNOWFLAKE_USER", "", true),
		snowflakePassword:  getEnv("SNOWFLAKE_PASSWORD", "", true),
		snowflakeWarehouse: getEnv("SNOWFLAKE_WAREHOUSE", "", true),
		snowflakeRole:      getEnv("SNOWFLAKE_ROLE", "", true),
		httpPort:           getEnv("HTTP_PORT", 50051, false),
		slackToken:         getEnv("SLACK_TOKEN", "", false),
		shutdownTimeout:    getEnv("SHUTDOWN_TIMEOUT", 10*time.Second, false),
	}

	c.snowflakeDSN = fmt.Sprintf(
		"%s:%s@%s/%s?warehouse=%s&role=%s",
		c.snowflakeUser,
		c.snowflakePassword,
		c.snowflakeAccount,
		"SNOWFLAKE",
		c.snowflakeWarehouse,
		c.snowflakeRole,
	)

	return c
}

func getEnv[T string | int | bool | time.Duration](key string, defaultVal T, required bool) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		if !required {
			return defaultVal
		} else {
			panic(fmt.Sprintf("missing required environment variable %s", key))
		}
	}

	var out T
	switch ptr := any(&out).(type) {
	case *string:
		{
			*ptr = val
		}
	case *int:
		{
			v, err := strconv.Atoi(val)
			if err != nil {
				return defaultVal
			}
			*ptr = v
		}
	case *bool:
		{
			v, err := strconv.ParseBool(val)
			if err != nil {
				return defaultVal
			}
			*ptr = v
		}
	case *time.Duration:
		{
			v, err := time.ParseDuration(val)
			if err != nil {
				return defaultVal
			}
			*ptr = v
		}
	default:
		{
			panic(fmt.Sprintf("unsupported type %T", out))
		}
	}

	return out
}
