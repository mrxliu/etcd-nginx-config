#!/usr/bin/env bash
# Runs NGiNX, then etcd-nginx-config. (Passes all arguments to the latter)
set -e

echo "Starting NGiNX..."
service nginx start
tail -f /var/log/nginx/* &

etcd-nginx-config $*

echo "Stopping NGiNX..."
service nginx stop

# stop tailing the logs...
kill %1
