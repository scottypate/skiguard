create user skiguard;
alter user skiguard set password = '<password_here>';
create role skiguard;

grant role skiguard to user skiguard;
grant usage on warehouse compute_wh to role skiguard;

use snowflake;

grant database role usage_viewer to role skiguard;
grant database role security_viewer to role skiguard;

-- All password accounts should use a network policy to restrict access.
create network rule skiguard_ingress
  mode = INGRESS
  type = IPV4
  value_list = ('<skiguard server IP>');

create network policy skiguard
  allowed_network_rule_list = ('skiguard_ingress');

alter user skiguard set network_policy = skiguard;