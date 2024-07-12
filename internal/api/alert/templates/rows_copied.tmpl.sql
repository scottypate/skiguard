with hours as (
    select
        date_trunc('hour', current_timestamp::timestamp + to_days(1)) - to_hours(i.generate_series) as generated_hour
    from
        generate_series(0, 744) as i
),

rows_copied as (
    select
        hours.generated_hour,
        sum(row_count) as rows_copied
    from
        hours
    left join
        main.snowflake_copy_history
    on
        date_trunc('hour', snowflake_copy_history.last_load_time) = hours.generated_hour
    where
        (last_load_time > current_date() - to_days(30) or last_load_time is null)
        and
        hours.generated_hour <= '{{ .NowUTC }}'
    group by 1 order by 1 desc
)
select
    generated_hour,
    case when rows_copied is null then 0 else rows_copied end as metric_value,
    row_number() over (order by generated_hour desc) as rn,
    max(rows_copied) over () as p100
from
    rows_copied
qualify
    rn = 1
    and
    rows_copied > (p100 * {{ .AlertThreshold }}) ;