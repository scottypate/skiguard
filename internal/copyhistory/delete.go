package copyhistory

import "github.com/scalecraft/skiguard/internal/duckdb"

func Delete() error {
	err := duckdb.Execute("delete from snowflake_copy_history")

	if err != nil {
		return err
	}

	return nil
}
