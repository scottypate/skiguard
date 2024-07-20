select
    user_id,
    login_name,
    created_on,
    coalesce(deleted_on, '1970-01-01') as deleted_on,
    coalesce(email, 'NONE') as email,
    has_password,
    ext_authn_duo,
    coalesce(ext_authn_uid, false) as ext_authn_uid,
    disabled,
    coalesce(last_success_login, '1970-01-01') as last_success_login,
    coalesce(password_last_set_time, '1970-01-01') as password_last_set_time
from
    snowflake.account_usage.users
where
    1=1