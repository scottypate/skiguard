package duckdb

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func Execute(queryString string) error {
	connector, err := CreateConnector()
	if err != nil {
		return err
	}

	db, err := OpenDatabase(connector)
	if err != nil {
		return err
	}

	defer db.Close()
	slog.Debug(fmt.Sprintf("querying duckdb database with query: %v", queryString))
	_, err = db.Exec(queryString)

	if err != nil {
		return err
	}

	return nil
}

func Query(queryString string) (*sql.Rows, error) {
	connector, err := CreateConnector()
	if err != nil {
		return nil, err
	}

	db, err := OpenDatabase(connector)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	slog.Debug(fmt.Sprintf("querying duckdb database with query: %v", queryString))
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
