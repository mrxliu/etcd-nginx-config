#!/usr/bin/env bash
# seeds etcd from runing Docker containers
# docker run -P -d --name=etcd coreos/etcd
set -e

echo "Checking for container named 'etcd'..."
docker inspect etcd >/dev/null 2>&1 || (echo "Starting etcd..." && \
  docker run -d --name=etcd -P -p 0.0.0.0:4001:4001 coreos/etcd)

#export ETCD_HOST='127.0.0.1'
export ETCD_PREFIX='/apps'

etcdctl ls $ETCD_PREFIX >/dev/null 2>&1 ||
  (echo "Creating key '${ETCD_PREFIX}'..." && etcdctl mkdir $ETCD_PREFIX)

echo "Creating HTTP endpoints for all local docker containers..."
IDS=$(docker ps -q)
for ID in $IDS ; do
  NAME=$(docker inspect --format '{{ .Name }}' $ID | sed 's/\///')
  if [[ "$NAME" != "router" && "$NAME" != "etcd" ]] ; then
    echo "Checking container $NAME for exposed ports..."
    VHOST="$NAME.*"
    URL="http://${NAME}.127.0.0.1.xip.io/"
    FORMAT='{{range $p, $c := .Config.ExposedPorts}} {{$p}} {{end}}'
    PORTS=$(docker inspect --format="$FORMAT" $NAME)
    IP=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' $NAME)
    for PORT in $PORTS ; do
      PORTNUM=$(echo $PORT | sed 's/[^0-9]//g')
      echo "Routing http://$VHOST -> ${IP}:${PORTNUM}"
      etcdctl set $ETCD_PREFIX/$NAME/endpoint "${IP}:${PORTNUM}" >/dev/null
    done
    etcdctl set $ETCD_PREFIX/$NAME/vhost $VHOST >/dev/null
    echo "$URL now proxies to ${NAME}, port $PORTNUM"
  fi
done

echo "Done."
