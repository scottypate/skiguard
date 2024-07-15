select
    coalesce(file_name, 'NONE') as file_name, 
    coalesce(stage_location, 'NONE') as stage_location, 
    last_load_time,
    coalesce(row_count, 0) as row_count, 
    coalesce(row_parsed, 0) as row_parsed, 
    coalesce(file_size, 0) as file_size, 
    coalesce(first_error_message, 'NONE') as first_error_message,
    coalesce(first_error_line_number, 0) as first_error_line_number,
    coalesce(first_error_character_pos, 0) as first_error_character_pos,
    coalesce(first_error_column_name, 'NONE') as first_error_column_name,
    coalesce(error_count, 0) as error_count,
    coalesce(error_limit, 0) as error_limit,
    coalesce(status, 'NONE') as status,
    coalesce(table_id, 0) as table_id,
    coalesce(table_name, 'NONE') as table_name,
    coalesce(table_schema_id, 0) as table_schema_id,
    coalesce(table_schema_name, 'NONE') as table_schema_name,
    coalesce(table_catalog_id, 0) as table_catalog_id,
    coalesce(table_catalog_name, 'NONE') as table_catalog_name,
    coalesce(pipe_catalog_name, 'NONE') as pipe_catalog_name,
    coalesce(pipe_schema_name, 'NONE') as pipe_schema_name,
    coalesce(pipe_name, 'NONE') as pipe_name,
    coalesce(pipe_received_time, '1970-01-01') as pipe_received_time,
    coalesce(first_commit_time, '1970-01-01') as first_commit_time
from
   snowflake.account_usage.copy_history
where
   1=1
   and
   last_load_time >= current_date() - interval '60 day'
