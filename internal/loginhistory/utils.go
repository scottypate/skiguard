package loginhistory

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/scalecraft/skiguard/internal/duckdb"
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

	appender, err := duckdb.NewAppender(connector, "main", "snowflake_login_history")
	if err != nil {
		return err
	}

	for rows.Next() {
		l := LoginHistory{}
		err = rows.Scan(
			&l.EventId,
			&l.EventTimestamp,
			&l.EventType,
			&l.UserName,
			&l.ClientIp,
			&l.ReportedClientType,
			&l.ReportedClientVersion,
			&l.FirstAuthenticationFactor,
			&l.SecondAuthenticationFactor,
			&l.IsSuccess,
			&l.ErrorCode,
			&l.ErrorMessage,
			&l.RelatedEventId,
			&l.Connection,
		)

		if err != nil {
			return err
		}

		duckDbRow := []driver.Value{
			l.EventId,
			l.EventTimestamp,
			l.EventType,
			l.UserName,
			l.ClientIp,
			l.ReportedClientType,
			l.ReportedClientVersion,
			l.FirstAuthenticationFactor,
			l.SecondAuthenticationFactor,
			l.IsSuccess,
			l.ErrorCode,
			l.ErrorMessage,
			l.RelatedEventId,
			l.Connection,
		}
		err = appender.AppendRow(duckDbRow...)
		if err != nil {
			return err
		}
	}
	err = appender.Close()
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Error closing appender for snowflake_login_history: %v", err,
		))
	}

	return nil
}
