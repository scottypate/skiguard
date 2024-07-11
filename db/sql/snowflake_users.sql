select 
    login_name,
    created_on,
    coalesce(email, 'NONE') as email,
    has_password,
    disabled,
    coalesce(last_success_login, '1970-01-01') as last_success_login,
    coalesce(password_last_set_time, '1970-01-01') as password_last_set_time
from
    snowflake.account_usage.users
where
    deleted_on is null
