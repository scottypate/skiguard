package users

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
)

func Update(db *sql.DB) error {

	err := Delete()

	if err != nil {
		return err
	}

	fileByte, err := os.ReadFile("db/sql/snowflake_users.sql")
	if err != nil {
		slog.Error(fmt.Sprintf("Error reading file snowflake_users: %s", err.Error()))
	}
	err = executeQuery(db, string(fileByte))

	if err != nil {
		return err
	}
	return nil
}
