package loginhistory

import "github.com/scalecraft/skiguard/internal/duckdb"

func Truncate() error {
	err := duckdb.Execute("delete from snowflake_login_history where event_timestamp < current_date - to_days(60)")

	if err != nil {
		return err
	}

	return nil
}
