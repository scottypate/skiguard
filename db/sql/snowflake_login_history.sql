select
    event_id,
    event_timestamp,
    event_type,
    user_name,
    client_ip,
    reported_client_type,
    reported_client_version,
    first_authentication_factor,
    coalesce(second_authentication_factor, 'NONE') as second_authentication_factor,
    is_success,
    coalesce(error_code, 0) as error_code,
    coalesce(error_message, 'NONE') as error_message,
    related_event_id,
    coalesce(connection, 'NONE') as connection,
from
    snowflake.account_usage.login_history;
