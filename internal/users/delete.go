package users

import "github.com/scalecraft/skiguard/internal/duckdb"

func Delete() error {
	err := duckdb.Execute("delete from snowflake_users")

	if err != nil {
		return err
	}

	return nil
}
