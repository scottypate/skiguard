package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/scalecraft/snowguard/internal/copyhistory"
	"github.com/scalecraft/snowguard/internal/duckdb"
	"github.com/scalecraft/snowguard/internal/loginhistory"
	"github.com/scalecraft/snowguard/internal/snowflake"
	"github.com/scalecraft/snowguard/internal/users"
)

func main() {
	runMigrations("db/schema/down")
	runMigrations("db/schema/up")
	snowflakeDb, err := snowflake.Connect()

	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to Snowflake: %s", err.Error()))
	}
	defer snowflakeDb.Close()

	err = loginhistory.Load(snowflakeDb)
	if err != nil {
		slog.Error(fmt.Sprintf("Error loading login_history: %s", err.Error()))
	}

	err = copyhistory.Load(snowflakeDb)
	if err != nil {
		slog.Error(fmt.Sprintf("Error loading copy_history: %s", err.Error()))
	}

	err = users.Load(snowflakeDb)
	if err != nil {
		slog.Error(fmt.Sprintf("Error loading users: %s", err.Error()))
	}

	duckdb.Execute("checkpoint;")
}

func runMigrations(dir string) {
	files, err := os.ReadDir(dir)

	if err != nil {
		slog.Error(fmt.Sprintf("Error reading directory: %s", err.Error()))
	}

	for _, file := range files {
		fileByte, err := os.ReadFile(dir + "/" + file.Name())
		if err != nil {
			slog.Error(fmt.Sprintf("Error reading file: %s", err.Error()))
		}
		duckdb.Execute(string(fileByte))
	}
}
