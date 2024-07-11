create table if not exists snowflake_users (
    login_name string primary key,
    created_on timestamp,
    email string,
    has_password boolean,
    disabled boolean,
    last_success_login timestamp,
    password_last_set_time timestamp
);
