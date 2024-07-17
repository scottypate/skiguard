package config

import (
	"fmt"
	"time"
)

type Config struct {
	SnowflakeAccount   string
	SnowflakeUser      string
	SnowflakePassword  string
	SnowflakeWarehouse string
	SnowflakeRole      string
	SlackToken         string
	SlackChannelId     string
	SnowflakeDSN       string
	HttpPort           int
	ShutdownTimeout    time.Duration
	AlertThreshold     float64
	NowUTC             string
	GinMode            string
}

func GetConfig() *Config {
	c := Config{
		SnowflakeAccount:   getEnv("SNOWFLAKE_ACCOUNT", "", true),
		SnowflakeUser:      getEnv("SNOWFLAKE_USER", "", true),
		SnowflakePassword:  getEnv("SNOWFLAKE_PASSWORD", "", true),
		SnowflakeWarehouse: getEnv("SNOWFLAKE_WAREHOUSE", "", true),
		SnowflakeRole:      getEnv("SNOWFLAKE_ROLE", "", true),
		HttpPort:           getEnv("HTTP_PORT", 50051, false),
		SlackToken:         getEnv("SLACK_TOKEN", "", false),
		SlackChannelId:     getEnv("SLACK_CHANNEL_ID", "", false),
		ShutdownTimeout:    getEnv("SHUTDOWN_TIMEOUT", 10*time.Second, false),
		AlertThreshold:     getEnv("ALERT_THRESHOLD", 0.90, false),
		GinMode:            getEnv("GIN_MODE", "release", false),
	}

	c.SnowflakeDSN = fmt.Sprintf(
		"%s:%s@%s/%s?warehouse=%s&role=%s",
		c.SnowflakeUser,
		c.SnowflakePassword,
		c.SnowflakeAccount,
		"SNOWFLAKE",
		c.SnowflakeWarehouse,
		c.SnowflakeRole,
	)

	return &c
}
