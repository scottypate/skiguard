package users

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log/slog"
	"os"

	"github.com/scalecraft/snowguard/internal/duckdb"
)

func Load(db *sql.DB) error {
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

func executeQuery(db *sql.DB, query string) error {
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = insertResults(rows)
	if err != nil {
		return err
	}
	return nil
}

func insertResults(rows *sql.Rows) error {
	connector, err := duckdb.CreateConnector()
	if err != nil {
		return err
	}

	appender, err := duckdb.NewAppender(connector, "main", "snowflake_users")
	if err != nil {
		return err
	}

	for rows.Next() {
		u := Users{}
		err = rows.Scan(
			&u.LoginName,
			&u.CreatedOn,
			&u.Email,
			&u.HasPassword,
			&u.Disabled,
			&u.LastSuccessLogin,
			&u.PasswordLastSetTime,
		)

		if err != nil {
			return err
		}

		duckDbRow := []driver.Value{
			u.LoginName,
			u.CreatedOn,
			u.Email,
			u.HasPassword,
			u.Disabled,
			u.LastSuccessLogin,
			u.PasswordLastSetTime,
		}
		err = appender.AppendRow(duckDbRow...)
		if err != nil {
			return err
		}
	}
	err = appender.Close()
	if err != nil {
		return err
	}

	return nil
}
