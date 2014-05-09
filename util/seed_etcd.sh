#!/usr/bin/env bash
# seeds etcd from runing Docker containers
# docker run -d --name=etcd -p 0.0.0.0:4001:4001 coreos/etcd
set -e

export ETCD_HOST='127.0.0.1'
export ETCD_PORT='4001'
export ETCD_PREFIX='/apps'

APPS='mauth eureka maudit archon'
for APP in $APPS ; do
  echo "Registering ${APP}..."
  etcdctl set $ETCD_PREFIX/$APP/vhost "$APP.*"
  etcdctl set $ETCD_PREFIX/$APP/endpoint $(docker port $APP 3000)
done

echo "Done."
