#!/bin/bash

# Login and get the JWT token

ACCESS_TOKEN=$(curl -X POST -H "Content-Type: application/json" -d '{
    "username": "admin",
    "password": "admin",
    "provider": "db"
}' http://localhost:8088/api/v1/security/login | jq -r .access_token)

# This script is used to build the Superset components (databases, datasets, etc.)

DATABASE_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database_name": "snowguard", 
    "engine": "duckdb", 
    "expose_in_sqllab": true,
    "sqlalchemy_uri": "duckdb:////home/snowguard/snowguard.db?access_mode=read_only"
}' http://localhost:8088/api/v1/database/ | jq -r .id)

LOGIN_DATASET_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database": "1",
    "schema": "main",
    "table_name": "snowflake_login_history"
}' http://localhost:8088/api/v1/dataset/ | jq -r .id)

COPY_DATASET_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database": "1",
    "schema": "main",
    "table_name": "snowflake_copy_history"
}' http://localhost:8088/api/v1/dataset/ | jq -r .id)

LOGINS_BY_SUCCESS_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "logins_by_success",
    "viz_type": "echarts_timeseries_bar",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"echarts_timeseries_bar\",\"slice_id\":1,\"x_axis\":\"event_timestamp\",\"time_grain_sqla\":\"P1D\",\"x_axis_sort_asc\":true,\"x_axis_sort_series\":\"name\",\"x_axis_sort_series_ascending\":true,\"metrics\":[{\"aggregate\":\"COUNT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_id\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":1,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"BIGINT\",\"type_generic\":0,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":false,\"label\":\"COUNT(event_id)\",\"optionName\":\"metric_xcvxqkdyq7n_ssmraeky6pq\",\"sqlExpression\":null}],\"groupby\":[\"is_success\"],\"adhoc_filters\":[{\"clause\":\"WHERE\",\"comparator\":\"No filter\",\"expressionType\":\"SIMPLE\",\"operator\":\"TEMPORAL_RANGE\",\"subject\":\"event_timestamp\"}],\"order_desc\":true,\"row_limit\":10000,\"truncate_metric\":true,\"show_empty_columns\":true,\"comparison_type\":\"values\",\"annotation_layers\":[],\"forecastPeriods\":10,\"forecastInterval\":0.8,\"orientation\":\"vertical\",\"x_axis_title_margin\":15,\"y_axis_title_margin\":15,\"y_axis_title_position\":\"Left\",\"sort_series_type\":\"sum\",\"color_scheme\":\"supersetColors\",\"only_total\":true,\"show_legend\":true,\"legendType\":\"scroll\",\"legendOrientation\":\"top\",\"x_axis_time_format\":\"smart_date\",\"y_axis_format\":\"SMART_NUMBER\",\"truncateXAxis\":true,\"y_axis_bounds\":[null,null],\"rich_tooltip\":true,\"tooltipTimeFormat\":\"smart_date\",\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8088/api/v1/chart/ | jq -r .id)

FIRST_LOGIN_PAST_30_DAYS=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "first_login_past_30_days",
    "viz_type": "table",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"table\",\"query_mode\":\"aggregate\",\"groupby\":[\"user_name\",\"client_ip\"],\"time_grain_sqla\":\"P1D\",\"temporal_columns_lookup\":{\"event_timestamp\":true},\"metrics\":[{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_timestamp\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":2,\"is_certified\":false,\"is_dttm\":true,\"python_date_format\":null,\"type\":\"TIMESTAMP WITHOUT TIME ZONE\",\"type_generic\":2,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"MIN\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":true,\"label\":\"first_login_time\",\"optionName\":\"metric_ozs73um0o8_5mmkqnoukah\"}],\"all_columns\":[],\"percent_metrics\":[],\"adhoc_filters\":[{\"expressionType\":\"SQL\",\"sqlExpression\":\"first_login_time >= max(current_date) - INTERVAL 30 DAY\",\"clause\":\"HAVING\",\"subject\":\"event_timestamp\",\"operator\":\"TEMPORAL_RANGE\",\"comparator\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_fr1dt5cd7d6_nn69s55zq1r\"}],\"order_by_cols\":[],\"row_limit\":1000,\"server_page_length\":10,\"order_desc\":true,\"table_timestamp_format\":\"smart_date\",\"show_cell_bars\":true,\"color_pn\":true,\"conditional_formatting\":[],\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8088/api/v1/chart/ | jq -r .id)

LOGINS_IP_BUBBLE=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "logins_ip_bubble",
    "viz_type": "bubble_v2",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"bubble_v2\",\"slice_id\":1,\"entity\":\"user_name\",\"x\":{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_id\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":1,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"BIGINT\",\"type_generic\":0,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"COUNT_DISTINCT\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":true,\"label\":\"number_logins\",\"optionName\":\"metric_3rztpspmthz_xljjuo6mvuq\"},\"y\":{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"client_ip\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":5,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"VARCHAR\",\"type_generic\":1,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"COUNT_DISTINCT\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":true,\"label\":\"distinct_client_ip\",\"optionName\":\"metric_w1hbiiyb2z_2qknpw41cup\"},\"adhoc_filters\":[{\"clause\":\"WHERE\",\"comparator\":\"Last month\",\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"filterOptionName\":\"filter_qsvon2ch777_p351xte1x5\",\"isExtra\":false,\"isNew\":false,\"operator\":\"TEMPORAL_RANGE\",\"sqlExpression\":null,\"subject\":\"event_timestamp\"}],\"size\":{\"aggregate\":\"COUNT_DISTINCT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"client_ip\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":5,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"VARCHAR\",\"type_generic\":1,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":false,\"label\":\"COUNT_DISTINCT(client_ip)\",\"optionName\":\"metric_4u83u526moj_7c9ms8bsay3\",\"sqlExpression\":null},\"order_desc\":true,\"row_limit\":10000,\"color_scheme\":\"supersetColors\",\"show_legend\":true,\"legendType\":\"scroll\",\"legendOrientation\":\"top\",\"max_bubble_size\":\"25\",\"tooltipSizeFormat\":\"SMART_NUMBER\",\"opacity\":0.6,\"x_axis_title_margin\":30,\"xAxisFormat\":\"SMART_NUMBER\",\"y_axis_title_margin\":30,\"y_axis_format\":\"SMART_NUMBER\",\"truncateXAxis\":true,\"y_axis_bounds\":[null,null],\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8088/api/v1/chart/ | jq -r .id)
