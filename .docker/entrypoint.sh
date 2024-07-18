#!/bin/bash

/sbin/service cron start

# Start Apache Superset
nohup sh -c /usr/bin/run-server.sh &

# Start SnowGuard
arch=$(uname -m)
machine=$(uname)
if [[ $arch == "arm64" ]] && [[ $machine == "Darwin" ]]; then
  /home/snowguard/bin/snowguard-darwin-arm64
elif [[ $arch == "x86_64" ]]; then
  /home/snowguard/bin/snowguard-linux-amd64
elif [[ $arch == "arm*" ]] || [[ $arch == "aarch64" ]]; then
  /home/snowguard/bin/snowguard-linux-arm64
fi
