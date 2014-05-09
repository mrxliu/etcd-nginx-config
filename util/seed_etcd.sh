#!/usr/bin/env bash
# seeds etcd
# docker run -d --name=etcd -p 0.0.0.0:4001:4001 coreos/etcd
set -e

export ETCD_HOST='127.0.0.1'
export ETCD_PORT='4001'
export ETCD_PREFIX='/apps'

# A simple app with only one endpoint and one vhost
etcdctl set $ETCD_PREFIX/app1/endpoint  '127.0.0.1:8080'
etcdctl set $ETCD_PREFIX/app1/vhost     "$APP_NAME\.*"

# an app with multiple endpoints and vhosts
etcdctl set $ETCD_PREFIX/app2/host1/vhost       "$APP_NAME\.mydomain\.com"
etcdctl set $ETCD_PREFIX/app2/host2/vhost       "$APP_NAME\.*"
etcdctl set $ETCD_PREFIX/app2/e1/endpoint       '127.0.0.1:8081'
etcdctl set $ETCD_PREFIX/app2/e2/endpoint       '127.0.0.1:8082'
etcdctl set $ETCD_PREFIX/app2/e3/rtsp/endpoint  '127.0.0.1:8083'

echo "Done."
