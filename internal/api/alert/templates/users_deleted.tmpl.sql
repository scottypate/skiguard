with hours as (
    select
        date_trunc('hour', current_timestamp::timestamp + to_days(1)) - to_hours(i.generate_series) as generated_hour
    from
        generate_series(0, 744) as i
),

users_deleted as (
    select
        hours.generated_hour,
        count(login_name) as users_deleted
    from
        hours
    left join
        main.snowflake_users
    on
        date_trunc('hour', snowflake_users.created_on) = hours.generated_hour
    where
        (deleted_on > current_date() - to_days(30) or created_on is null)
        and
        hours.generated_hour <= '{{ .NowUTC }}'
    group by 1 order by 1 desc
)
select
    generated_hour,
    case when users_deleted is null then 0 else users_deleted end as metric_value,
    row_number() over (order by generated_hour desc) as rn,
    max(users_deleted) over () as p100
from
    users_deleted
qualify
    rn = 1
    and
    users_deleted > (p100 * {{ .AlertThreshold }}) ;