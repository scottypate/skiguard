with hours as (
    select
        date_trunc('hour', current_timestamp::timestamp + to_days(1)) - to_hours(i.generate_series) as generated_hour
    from
        generate_series(0, 744) as i
),

total_logins as (
    select
        hours.generated_hour,
        count(event_id) as total_logins
    from
        hours
    left join
        main.snowflake_login_history
    on
        date_trunc('hour', snowflake_login_history.event_timestamp) = hours.generated_hour
    where
        (event_timestamp > current_date() - to_days(30) or event_timestamp is null)
        and
        hours.generated_hour <= '{{ .NowUTC }}'
    group by 1 order by 1 desc
)
select
    generated_hour,
    total_logins as metric_value,
    row_number() over (order by generated_hour desc) as rn,
    max(total_logins) over () as p100
from
    total_logins
qualify
    rn = 1
    and
    total_logins > (p100 * {{ .AlertThreshold }}) ;