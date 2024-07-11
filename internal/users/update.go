package users

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/scalecraft/snowguard/internal/duckdb"
)

func Update(db *sql.DB) error {
	latestDate, err := getLatestDate()
	fileByte, err := os.ReadFile("db/sql/snowflake_users.sql")
	if err != nil {
		slog.Error(fmt.Sprintf("Error reading file snowflake_users: %s", err.Error()))
	}

	err = executeQuery(db, string(fileByte)+" and created_on > "+*latestDate)

	if err != nil {
		return err
	}
	return nil
}

func getLatestDate() (*string, error) {
	var results string
	sql := "select max(created_on) from main.snowflake_users"
	row, err := duckdb.Query(sql)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&results)
		if err != nil {
			return nil, err
		}
	}
	return &results, nil

}
