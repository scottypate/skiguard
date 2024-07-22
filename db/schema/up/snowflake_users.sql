create table if not exists snowflake_users (
    user_id bigint primary key,
    login_name string,
    created_on timestamp,
    deleted_on timestamp,
    email string,
    has_password boolean,
    ext_authn_duo boolean,
    ext_authn_uid string,
    disabled boolean,
    last_success_login timestamp,
    password_last_set_time timestamp
);
