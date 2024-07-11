#!/bin/bash

# Login and get the JWT token

ACCESS_TOKEN=$(curl -X POST -H "Content-Type: application/json" -d '{
    "username": "admin",
    "password": "admin",
    "provider": "db"
}' http://localhost:8080/api/v1/security/login | jq -r .access_token)

# This script is used to build the Superset components (databases, datasets, etc.)

DATABASE_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database_name": "snowguard", 
    "engine": "duckdb", 
    "expose_in_sqllab": true,
    "sqlalchemy_uri": "duckdb:////home/snowguard/snowguard.db?access_mode=read_only"
}' http://localhost:8080/api/v1/database/ | jq -r .id)

LOGIN_DATASET_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database": "1",
    "schema": "main",
    "table_name": "snowflake_login_history"
}' http://localhost:8080/api/v1/dataset/ | jq -r .id)

COPY_DATASET_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database": "1",
    "schema": "main",
    "table_name": "snowflake_copy_history"
}' http://localhost:8080/api/v1/dataset/ | jq -r .id)

USER_DATASET_ID=$(curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "database": "1",
    "schema": "main",
    "table_name": "snowflake_users"
}' http://localhost:8080/api/v1/dataset/ | jq -r .id)

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "Logins by Success",
    "description": "Count of all logins by success status. Shown for the previous 30 days by default.",
    "viz_type": "echarts_area",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"echarts_area\",\"slice_id\":1,\"x_axis\":\"event_timestamp\",\"time_grain_sqla\":\"P1D\",\"x_axis_sort_asc\":true,\"x_axis_sort_series\":\"name\",\"x_axis_sort_series_ascending\":true,\"metrics\":[{\"aggregate\":\"COUNT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_id\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":1,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"BIGINT\",\"type_generic\":0,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":false,\"label\":\"COUNT(event_id)\",\"optionName\":\"metric_xcvxqkdyq7n_ssmraeky6pq\",\"sqlExpression\":null}],\"groupby\":[\"is_success\"],\"adhoc_filters\":[{\"clause\":\"WHERE\",\"comparator\":\"Last month\",\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"filterOptionName\":\"filter_jmel5uqatbb_67rteny3ct\",\"isExtra\":false,\"isNew\":false,\"operator\":\"TEMPORAL_RANGE\",\"sqlExpression\":null,\"subject\":\"event_timestamp\"}],\"order_desc\":true,\"row_limit\":10000,\"truncate_metric\":true,\"show_empty_columns\":true,\"comparison_type\":\"values\",\"annotation_layers\":[],\"forecastPeriods\":10,\"forecastInterval\":0.8,\"x_axis_title_margin\":15,\"y_axis_title\":\"Count Logins\",\"y_axis_title_margin\":15,\"y_axis_title_position\":\"Left\",\"sort_series_type\":\"sum\",\"color_scheme\":\"supersetColors\",\"seriesType\":\"line\",\"opacity\":0.6,\"show_value\":false,\"stack\":null,\"only_total\":true,\"markerEnabled\":false,\"markerSize\":6,\"minorTicks\":false,\"zoomable\":false,\"show_legend\":true,\"legendType\":\"scroll\",\"legendOrientation\":\"top\",\"x_axis_time_format\":\"%Y-%m-%d\",\"rich_tooltip\":true,\"tooltipTimeFormat\":\"smart_date\",\"y_axis_format\":\"SMART_NUMBER\",\"truncateXAxis\":true,\"y_axis_bounds\":[null,null],\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "New IP Address Logins",
    "description": "Record of the first login time for each user and IP address. Shown for the previous 30 days by default.",
    "viz_type": "table",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"table\",\"query_mode\":\"aggregate\",\"groupby\":[\"user_name\",\"client_ip\"],\"time_grain_sqla\":\"P1D\",\"temporal_columns_lookup\":{\"event_timestamp\":true},\"metrics\":[{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_timestamp\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":2,\"is_certified\":false,\"is_dttm\":true,\"python_date_format\":null,\"type\":\"TIMESTAMP WITHOUT TIME ZONE\",\"type_generic\":2,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"MIN\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":true,\"label\":\"first_login_time\",\"optionName\":\"metric_ozs73um0o8_5mmkqnoukah\"}],\"all_columns\":[],\"percent_metrics\":[],\"adhoc_filters\":[{\"expressionType\":\"SQL\",\"sqlExpression\":\"first_login_time >= max(current_date) - INTERVAL 30 DAY\",\"clause\":\"HAVING\",\"subject\":\"event_timestamp\",\"operator\":\"TEMPORAL_RANGE\",\"comparator\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_fr1dt5cd7d6_nn69s55zq1r\"}],\"order_by_cols\":[],\"row_limit\":1000,\"server_page_length\":10,\"order_desc\":true,\"table_timestamp_format\":\"smart_date\",\"show_cell_bars\":true,\"color_pn\":true,\"conditional_formatting\":[],\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "Distinct IP Addresses by Logins",
    "description": "Bubble size represents the number of distinct IP addresses for that Snowflake user. Consider using network policies to restrict password access to Snowflake.",
    "viz_type": "bubble_v2",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"bubble_v2\",\"slice_id\":3,\"entity\":\"user_name\",\"x\":{\"aggregate\":\"COUNT_DISTINCT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_id\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":1,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"BIGINT\",\"type_generic\":0,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":true,\"label\":\"number_logins\",\"optionName\":\"metric_3rztpspmthz_xljjuo6mvuq\",\"sqlExpression\":null},\"y\":{\"aggregate\":\"COUNT_DISTINCT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"client_ip\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":5,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"VARCHAR\",\"type_generic\":1,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":true,\"label\":\"distinct_client_ip\",\"optionName\":\"metric_w1hbiiyb2z_2qknpw41cup\",\"sqlExpression\":null},\"adhoc_filters\":[{\"clause\":\"WHERE\",\"comparator\":\"Last month\",\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"filterOptionName\":\"filter_qsvon2ch777_p351xte1x5\",\"isExtra\":false,\"isNew\":false,\"operator\":\"TEMPORAL_RANGE\",\"sqlExpression\":null,\"subject\":\"event_timestamp\"}],\"size\":{\"aggregate\":\"COUNT_DISTINCT\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"client_ip\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":5,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"VARCHAR\",\"type_generic\":1,\"verbose_name\":null,\"warning_markdown\":null},\"datasourceWarning\":false,\"expressionType\":\"SIMPLE\",\"hasCustomLabel\":false,\"label\":\"COUNT_DISTINCT(client_ip)\",\"optionName\":\"metric_4u83u526moj_7c9ms8bsay3\",\"sqlExpression\":null},\"order_desc\":true,\"row_limit\":10000,\"color_scheme\":\"supersetColors\",\"show_legend\":false,\"legendType\":\"scroll\",\"legendOrientation\":\"top\",\"legendMargin\":null,\"max_bubble_size\":\"25\",\"tooltipSizeFormat\":\"SMART_NUMBER\",\"opacity\":0.9,\"x_axis_label\":\"Total Number of Logins\",\"x_axis_title_margin\":30,\"xAxisFormat\":\"SMART_NUMBER\",\"logXAxis\":false,\"y_axis_label\":\"Distinct IP Addresses\",\"y_axis_title_margin\":30,\"y_axis_format\":\"SMART_NUMBER\",\"logYAxis\":false,\"truncateXAxis\":false,\"y_axis_bounds\":[null,null],\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${COPY_DATASET_ID}',
    "datasource_type": "table",
    "description": "The sum of all rows copied via `COPY INTO` statements and loaded via Snowpipe. Shown for the previous 30 days by default.",
    "slice_name": "Total Rows Copied",
    "viz_type": "echarts_timeseries_bar",
    "params": "{\"datasource\":\"2__table\",\"viz_type\":\"echarts_timeseries_bar\",\"x_axis\":\"last_load_time\",\"time_grain_sqla\":\"P1D\",\"x_axis_sort_asc\":true,\"x_axis_sort_series\":\"name\",\"x_axis_sort_series_ascending\":true,\"metrics\":[{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"row_count\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":18,\"is_certified\":false,\"is_dttm\":false,\"python_date_format\":null,\"type\":\"BIGINT\",\"type_generic\":0,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"SUM\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":false,\"label\":\"SUM(row_count)\",\"optionName\":\"metric_b9e1y2hfyxh_wil1vf5dyl\"}],\"groupby\":[],\"adhoc_filters\":[{\"expressionType\":\"SIMPLE\",\"subject\":\"last_load_time\",\"operator\":\"TEMPORAL_RANGE\",\"comparator\":\"Last month\",\"clause\":\"WHERE\",\"sqlExpression\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_sj1tu4gmfoc_fm3sdpxzfj\"}],\"order_desc\":true,\"row_limit\":10000,\"truncate_metric\":true,\"show_empty_columns\":true,\"comparison_type\":\"values\",\"annotation_layers\":[],\"forecastPeriods\":10,\"forecastInterval\":0.8,\"orientation\":\"vertical\",\"x_axis_title_margin\":15,\"y_axis_title\":\"Sum of Rows Copied\",\"y_axis_title_margin\":50,\"y_axis_title_position\":\"Left\",\"sort_series_type\":\"sum\",\"color_scheme\":\"supersetColors\",\"only_total\":true,\"show_legend\":true,\"legendType\":\"scroll\",\"legendOrientation\":\"top\",\"x_axis_time_format\":\"%Y-%m-%d\",\"y_axis_format\":\"SMART_NUMBER\",\"truncateXAxis\":true,\"y_axis_bounds\":[null,null],\"rich_tooltip\":true,\"tooltipTimeFormat\":\"smart_date\",\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${LOGIN_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "Logins without MFA",
    "description": "Snowflake logins without Multi-Factor Authentication (MFA). Note this does not account for external MFA via an IdP. Shown for the previous 30 days by default.",
    "viz_type": "table",
    "params": "{\"datasource\":\"1__table\",\"viz_type\":\"table\",\"query_mode\":\"aggregate\",\"groupby\":[\"event_timestamp\",\"user_name\",\"client_ip\",\"reported_client_type\",\"first_authentication_factor\",\"second_authentication_factor\",\"is_success\"],\"time_grain_sqla\":\"PT1S\",\"temporal_columns_lookup\":{\"event_timestamp\":true},\"all_columns\":[],\"percent_metrics\":[],\"adhoc_filters\":[{\"expressionType\":\"SIMPLE\",\"subject\":\"event_timestamp\",\"operator\":\"TEMPORAL_RANGE\",\"comparator\":\"Last month\",\"clause\":\"WHERE\",\"sqlExpression\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_inrowdvn27d_pouflbo8ute\"},{\"expressionType\":\"SIMPLE\",\"subject\":\"second_authentication_factor\",\"operator\":\"==\",\"operatorId\":\"EQUALS\",\"comparator\":\"NONE\",\"clause\":\"WHERE\",\"sqlExpression\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_894yx5fs0zh_7c9s6ymo5jd\"}],\"timeseries_limit_metric\":{\"expressionType\":\"SIMPLE\",\"column\":{\"advanced_data_type\":null,\"certification_details\":null,\"certified_by\":null,\"column_name\":\"event_timestamp\",\"description\":null,\"expression\":null,\"filterable\":true,\"groupby\":true,\"id\":2,\"is_certified\":false,\"is_dttm\":true,\"python_date_format\":null,\"type\":\"TIMESTAMP WITHOUT TIME ZONE\",\"type_generic\":2,\"verbose_name\":null,\"warning_markdown\":null},\"aggregate\":\"MAX\",\"sqlExpression\":null,\"datasourceWarning\":false,\"hasCustomLabel\":false,\"label\":\"MAX(event_timestamp)\",\"optionName\":\"metric_w7dpp8l35rr_0cgv3w936m6\"},\"order_by_cols\":[],\"row_limit\":1000,\"server_page_length\":10,\"order_desc\":true,\"table_timestamp_format\":\"smart_date\",\"show_cell_bars\":true,\"color_pn\":true,\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${USER_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "Accounts with Aged Passwords",
    "description": "Snowflake users with passwords that have not been changed in the last 6 months. Consider rotating these user passwords.",
    "viz_type": "table",
    "params": "{\"datasource\":\"3__table\",\"viz_type\":\"table\",\"query_mode\":\"aggregate\",\"groupby\":[\"login_name\",\"email\",\"has_password\",\"disabled\",\"password_last_set_time\"],\"time_grain_sqla\":\"P1D\",\"temporal_columns_lookup\":{\"created_on\":true,\"last_success_login\":true,\"password_last_set_time\":true},\"all_columns\":[],\"percent_metrics\":[],\"adhoc_filters\":[{\"expressionType\":\"SQL\",\"sqlExpression\":\"password_last_set_time < current_date - INTERVAL 6 MONTH\",\"clause\":\"WHERE\",\"subject\":\"password_last_set_time\",\"operator\":\"TEMPORAL_RANGE\",\"operatorId\":\"TEMPORAL_RANGE\",\"comparator\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_t4s0c66f6s_jbtrgiykx3c\"},{\"expressionType\":\"SIMPLE\",\"subject\":\"has_password\",\"operator\":\"IN\",\"operatorId\":\"IN\",\"comparator\":[true],\"clause\":\"WHERE\",\"sqlExpression\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_hf76uhutatk_ihcds63ul6\"}],\"order_by_cols\":[],\"row_limit\":1000,\"server_page_length\":10,\"order_desc\":true,\"table_timestamp_format\":\"smart_date\",\"show_cell_bars\":true,\"color_pn\":true,\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "datasource_id": '${USER_DATASET_ID}',
    "datasource_type": "table",
    "slice_name": "Accounts without Logins",
    "description": "Snowflake users who have never logged in. These user accounts may be stale or unused.",
    "viz_type": "table",
    "params": "{\"datasource\":\"3__table\",\"viz_type\":\"table\",\"query_mode\":\"aggregate\",\"groupby\":[\"login_name\",\"email\",\"created_on\",{\"label\":\"last_success_login\",\"sqlExpression\":\"case \\n  when last_success_login = '"'"'1970-01-01'"'"'\\n  then NULL\\nend\",\"expressionType\":\"SQL\"}],\"time_grain_sqla\":\"P1D\",\"temporal_columns_lookup\":{\"created_on\":true,\"last_success_login\":true,\"password_last_set_time\":true},\"all_columns\":[],\"percent_metrics\":[],\"adhoc_filters\":[{\"expressionType\":\"SQL\",\"sqlExpression\":\"last_success_login = '"'"'1970-01-01'"'"'\",\"clause\":\"WHERE\",\"subject\":\"last_success_login\",\"operator\":\"TEMPORAL_RANGE\",\"operatorId\":\"TEMPORAL_RANGE\",\"comparator\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_o5zyearnb6_3qlsdad6hz4\"},{\"expressionType\":\"SIMPLE\",\"subject\":\"login_name\",\"operator\":\"!=\",\"operatorId\":\"NOT_EQUALS\",\"comparator\":\"SNOWFLAKE\",\"clause\":\"WHERE\",\"sqlExpression\":null,\"isExtra\":false,\"isNew\":false,\"datasourceWarning\":false,\"filterOptionName\":\"filter_5dgbdkp2soc_03wl89dz8ovl\"}],\"order_by_cols\":[],\"row_limit\":1000,\"server_page_length\":10,\"order_desc\":true,\"table_timestamp_format\":\"smart_date\",\"show_cell_bars\":true,\"color_pn\":true,\"extra_form_data\":{},\"dashboards\":[]}"
}' http://localhost:8080/api/v1/chart/

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{
    "dashboard_title": "SnowGuard Security Dashboard",
    "css": ".grid-content .dragdroppable-row .dragdroppable-column .dashboard-component-chart-holder {\n  \n  border-radius:5px;\n}",
    "json_metadata": "{\"chart_configuration\": {\"1\": {\"id\": 1, \"crossFilters\": {\"scope\": \"global\", \"chartsInScope\": [2, 3, 4, 5, 6]}}, \"2\": {\"id\": 2, \"crossFilters\": {\"scope\": \"global\", \"chartsInScope\": [1, 3, 4, 5, 6]}}, \"3\": {\"id\": 3, \"crossFilters\": {\"scope\": \"global\", \"chartsInScope\": [1, 2, 4, 5, 6]}}, \"4\": {\"id\": 4, \"crossFilters\": {\"scope\": {\"rootPath\": [], \"excluded\": [4]}, \"chartsInScope\": []}}, \"5\": {\"id\": 5, \"crossFilters\": {\"scope\": \"global\", \"chartsInScope\": [1, 2, 3, 4, 6]}}, \"6\": {\"id\": 6, \"crossFilters\": {\"scope\": \"global\", \"chartsInScope\": [1, 2, 3, 4, 5]}}}, \"global_chart_configuration\": {\"scope\": {\"rootPath\": [\"ROOT_ID\"], \"excluded\": []}, \"chartsInScope\": [1, 2, 3, 4, 5, 6]}, \"color_scheme\": \"\", \"refresh_frequency\": 0, \"shared_label_colors\": {\"SUM(row_count)\": \"#1FA8C9\", \"YES\": \"#454E7C\", \"NO\": \"#FF7F44\"}, \"color_scheme_domain\": [], \"expanded_slices\": {}, \"label_colors\": {}, \"timed_refresh_immune_slices\": [], \"cross_filters_enabled\": true, \"default_filters\": \"{}\"}",
    "published": "true"
}' http://localhost:8080/api/v1/dashboard/

curl -X GET -H "Authorization: Bearer ${ACCESS_TOKEN}" http://localhost:8080/api/v1/dashboard/ | jq