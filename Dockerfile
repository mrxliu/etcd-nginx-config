FROM ubuntu
MAINTAINER Benton Roberts <broberts@mdsol.com>

ADD . /app
WORKDIR /app

RUN apt-get -y install nginx curl
RUN rm -f /etc/nginx/sites-enabled/default
RUN curl http://download.bentonroberts.com/etcd-nginx-config/0.1.3/etcd-nginx-config_0.1.3_linux_amd64.tar.gz | tar -xvzf -
RUN cp -p etcd-nginx-config_0.1.3_linux_amd64/etcd-nginx-config /usr/local/sbin/

EXPOSE 80
EXPOSE 443
CMD /app/util/run_with_nginx.sh -etcd-hosts=http://${ETCD_PORT_4001_TCP_ADDR}:4001
