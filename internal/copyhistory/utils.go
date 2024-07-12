package copyhistory

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/scalecraft/snowguard/internal/duckdb"
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

	appender, err := duckdb.NewAppender(connector, "main", "snowflake_copy_history")
	if err != nil {
		return err
	}

	for rows.Next() {
		c := CopyHistory{}
		err = rows.Scan(
			&c.Filename,
			&c.StageLocation,
			&c.LastLoadTime,
			&c.RowCount,
			&c.RowParsed,
			&c.FileSize,
			&c.FirstErrorMessage,
			&c.FirstErrorLineNumber,
			&c.FirstErrorCharacterPos,
			&c.FirstErrorColumnName,
			&c.ErrorCount,
			&c.ErrorLimit,
			&c.Status,
			&c.TableId,
			&c.TableName,
			&c.TableSchemaId,
			&c.TableSchemaName,
			&c.TableCatalogId,
			&c.TableCatalogName,
			&c.PipeCatalogName,
			&c.PipeSchemaName,
			&c.PipeName,
			&c.PipeReceivedTime,
			&c.FirstCommitTime,
		)

		if err != nil {
			return err
		}

		duckDbRow := []driver.Value{
			c.Filename,
			c.StageLocation,
			c.LastLoadTime,
			c.RowCount,
			c.RowParsed,
			c.FileSize,
			c.FirstErrorMessage,
			c.FirstErrorLineNumber,
			c.FirstErrorCharacterPos,
			c.FirstErrorColumnName,
			c.ErrorCount,
			c.ErrorLimit,
			c.Status,
			c.TableId,
			c.TableName,
			c.TableSchemaId,
			c.TableSchemaName,
			c.TableCatalogId,
			c.TableCatalogName,
			c.PipeCatalogName,
			c.PipeSchemaName,
			c.PipeName,
			c.PipeReceivedTime,
			c.FirstCommitTime,
		}
		err = appender.AppendRow(duckDbRow...)
		if err != nil {
			return err
		}
	}
	err = appender.Close()
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Error closing appender for snowflake_copy_history: %v", err,
		))
	}
	return nil
}
