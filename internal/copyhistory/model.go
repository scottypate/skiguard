package copyhistory

import "time"

type CopyHistory struct {
	Filename               string    `json:"file_name"`
	StageLocation          string    `json:"stage_location"`
	LastLoadTime           time.Time `json:"last_load_time"`
	RowCount               int64     `json:"row_count"`
	RowParsed              int64     `json:"row_parsed"`
	FileSize               int64     `json:"file_size"`
	FirstErrorMessage      string    `json:"first_error_message"`
	FirstErrorLineNumber   int64     `json:"first_error_line_number"`
	FirstErrorCharacterPos int64     `json:"first_error_character_pos"`
	FirstErrorColumnName   string    `json:"first_error_column_name"`
	ErrorCount             int64     `json:"error_count"`
	ErrorLimit             int64     `json:"error_limit"`
	Status                 string    `json:"status"`
	TableId                int64     `json:"table_id"`
	TableName              string    `json:"table_name"`
	TableSchemaId          int64     `json:"table_schema_id"`
	TableSchemaName        string    `json:"table_schema_name"`
	TableCatalogId         int64     `json:"table_catalog_id"`
	TableCatalogName       string    `json:"table_catalog_name"`
	PipeCatalogName        string    `json:"pipe_catalog_name"`
	PipeSchemaName         string    `json:"pipe_schema_name"`
	PipeName               string    `json:"pipe_name"`
	PipeReceivedTime       time.Time `json:"pipe_received_time"`
	FirstCommitTime        time.Time `json:"first_commit_time"`
}
