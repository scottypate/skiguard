package copyhistory

import "github.com/scalecraft/skiguard/internal/duckdb"

func Truncate() error {
	err := duckdb.Execute("delete from snowflake_copy_history where last_load_time < current_date - to_days(60)")

	if err != nil {
		return err
	}

	return nil
}
