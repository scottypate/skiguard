select
    coalesce(event_id::integer, 0) as event_id,
    event_timestamp,
    coalesce(event_type, 'NONE') as event_type,
    coalesce(user_name, 'NONE') as user_name,
    coalesce(client_ip, 'NONE') as client_ip,
    coalesce(reported_client_type, 'NONE') as reported_client_type,
    coalesce(reported_client_version, 'NONE') as reported_client_version,
    coalesce(first_authentication_factor, 'NONE') as first_authentication_factor,
    coalesce(second_authentication_factor, 'NONE') as second_authentication_factor,
    coalesce(is_success, 'NONE') as is_success,
    coalesce(error_code::integer, 0) as error_code,
    coalesce(error_message, 'NONE') as error_message,
    coalesce(related_event_id::integer, 0) as related_event_id,
    coalesce(connection, 'NONE') as connection
from
    snowflake.account_usage.login_history
where
    1=1
    and
    event_timestamp >= current_date() - interval '60 day'