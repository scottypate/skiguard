create table if not exists snowflake_login_history (
    event_id bigint primary key,
    event_timestamp timestamp,
    event_type string,
    user_name string,
    client_ip string,
    reported_client_type string,
    reported_client_version string,
    first_authentication_factor string,
    second_authentication_factor string,
    is_success string,
    error_code bigint,
    error_message string,
    related_event_id bigint,
    connection string
);