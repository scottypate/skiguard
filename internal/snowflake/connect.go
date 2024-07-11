package snowflake

import (
	"database/sql"

	_ "github.com/snowflakedb/gosnowflake"
)

func Connect(snowflakeConnStr string) (*sql.DB, error) {
	db, err := sql.Open("snowflake", snowflakeConnStr)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("use database SNOWFLAKE;")

	if err != nil {
		return nil, err
	}

	return db, nil
}
