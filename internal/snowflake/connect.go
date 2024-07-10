package snowflake

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/snowflakedb/gosnowflake"
)

func Connect() (*sql.DB, error) {
	loadEnv()
	connStr, err := connectionString()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("snowflake", *connStr)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("use database SNOWFLAKE;")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Debug(fmt.Sprintf("no .env file found: %s", err.Error()))
	}
}

func connectionString() (*string, error) {
	user, ok := os.LookupEnv("SNOWFLAKE_USER")
	if !ok {
		return nil, fmt.Errorf("environment variable snowflake_user not set")
	}

	password, ok := os.LookupEnv("SNOWFLAKE_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("environment variable snowflake_password not set")
	}

	account, ok := os.LookupEnv("SNOWFLAKE_ACCOUNT")
	if !ok {
		return nil, fmt.Errorf("environment variable snowflake_account not set")
	}

	warehouse, ok := os.LookupEnv("SNOWFLAKE_WAREHOUSE")
	if !ok {
		return nil, fmt.Errorf("environment variable snowflake_warehouse not set")
	}

	role, ok := os.LookupEnv("SNOWFLAKE_ROLE")
	if !ok {
		return nil, fmt.Errorf("environment variable snowflake_role not set")
	}

	connStr := fmt.Sprintf(
		"%s:%s@%s/%s?warehouse=%s&role=%s",
		user, password, account, "SNOWFLAKE", warehouse, role,
	)

	return &connStr, nil
}
