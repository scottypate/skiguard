package snowflake

import "database/sql"

type Querier interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Execute(query string, args ...any) (sql.Result, error)
}
