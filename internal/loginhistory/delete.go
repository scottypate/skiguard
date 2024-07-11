package loginhistory

import "github.com/scalecraft/snowguard/internal/duckdb"

func Delete() error {
	err := duckdb.Execute("delete from snowflake_login_history")

	if err != nil {
		return err
	}

	return nil
}
