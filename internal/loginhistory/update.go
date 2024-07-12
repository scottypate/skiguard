package loginhistory

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/scalecraft/snowguard/internal/duckdb"
)

func Update(db *sql.DB) error {
	latestDate, err := getLatestDate()
	if err != nil {
		return err
	}

	fileByte, err := os.ReadFile("db/sql/snowflake_login_history.sql")
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("%v and event_timestamp >= '%s'::timestamp + interval '1 second'", string(fileByte), *latestDate)

	err = executeQuery(db, sql)

	if err != nil {
		return err
	}
	return nil
}

func getLatestDate() (*string, error) {
	var iso8601 string
	sql := "select max(event_timestamp) from main.snowflake_login_history"
	row, err := duckdb.Query(sql)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&iso8601)
		if err != nil {
			return nil, err
		}
	}

	t, err := time.Parse("2006-01-02T15:04:05Z", iso8601)

	if err != nil {
		return nil, err
	}

	result := t.Format("2006-01-02 15:04:05")

	return &result, nil

}
