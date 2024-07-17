#!/bin/bash

/sbin/service cron start

# Start Apache Superset
nohup sh -c /usr/bin/run-server.sh &

# Start SnowGuard
/home/snowguard/bin/snowguard-linux-amd64
