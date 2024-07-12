#!/bin/bash

# Start Apache Superset
nohup sh -c /usr/bin/run-server.sh &

# Start SnowGuard
/home/snowguard/bin/snowguard-linux-arm64
