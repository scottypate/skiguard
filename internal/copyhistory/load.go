package copyhistory

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
)

func Load(db *sql.DB) error {
	fileByte, err := os.ReadFile("db/sql/snowflake_copy_history.sql")
	if err != nil {
		slog.Error(fmt.Sprintf("Error reading file snowflake_copy_history: %s", err.Error()))
	}
	err = executeQuery(db, string(fileByte))

	if err != nil {
		return err
	}
	return nil
}
