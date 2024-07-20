package users

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/scalecraft/snowguard/internal/duckdb"
)

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
			&u.UserId,
			&u.LoginName,
			&u.CreatedOn,
			&u.DeletedOn,
			&u.Email,
			&u.HasPassword,
			&u.ExtAuthnDuo,
			&u.ExtAuthnUid,
			&u.Disabled,
			&u.LastSuccessLogin,
			&u.PasswordLastSetTime,
		)

		if err != nil {
			return err
		}

		duckDbRow := []driver.Value{
			u.UserId,
			u.LoginName,
			u.CreatedOn,
			u.DeletedOn,
			u.Email,
			u.HasPassword,
			u.ExtAuthnDuo,
			u.ExtAuthnUid,
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
		return errors.New(fmt.Sprintf(
			"Error closing appender for snowflake_users: %v", err,
		))
	}

	return nil
}
