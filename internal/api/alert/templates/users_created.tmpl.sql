with hours as (
    select
        date_trunc('hour', current_timestamp::timestamp + to_days(1)) - to_hours(i.generate_series) as generated_hour
    from
        generate_series(0, 744) as i
),

users_created as (
    select
        hours.generated_hour,
        count(login_name) as users_created
    from
        hours
    left join
        main.snowflake_users
    on
        date_trunc('hour', snowflake_users.created_on) = hours.generated_hour
    where
        (created_on > current_date() - to_days(30) or created_on is null)
        and
        hours.generated_hour <= '{{ .NowUTC }}'
    group by 1 order by 1 desc
)
select
    generated_hour,
    case when users_created is null then 0 else users_created end as metric_value,
    row_number() over (order by generated_hour desc) as rn,
    max(users_created) over () as p100
from
    users_created
qualify
    rn = 1
    and
    users_created > (p100 * {{ .AlertThreshold }}) ;