Etcd-Nginx Web Router
================
Writes NGiNX reverse-proxy config files for multiple backends, based on data in [etcd][1].

  *ALPHA VERSION - needs more tests*


----------------
Overview
----------------
This program enables dynamic, centralized HTTP routing based on minimal config. It connects to the etcd configuration service and looks for information about backend web apps. It builds a list of such apps, their HTTP endpoints and virtual host names, and then writes one Nginx reverse-proxy configuration file for each app.

The software then watches etcd for changes, updates the files as needed, and signals nginx (via its PID file) to reload its configuration.


----------------
Where is it? (Installation)
----------------
Download the latest binary for your OS from the [downloads page](downloads.md).

Uncompress it wherever you like. Somewhere in your $PATH might be handy.


----------------
How is it [done]? (Usage)
----------------
1) Install nginx and configure it to load all `*.conf` files from a configuration directory (this is the default setup for most packages). The `etcd-nginx-config` binary will write config files to this location, which can be passed to `etcd-nginx-config` as `-nginx-dir`. It defaults to `/etc/nginx/conf.d`.

2) Start nginx. It will create a PID file, which this software reads in order to signal nginx to reload its config. The path to this PID file can be passed as `-nginx-pid`, but defaults to `/var/run/nginx.pid`.

3) Seed etcd with info about your backend web apps. Each backend app is just a list of HTTP endpoints and a list of virtual hostname patterns, both assocated with an ID. All data about the web apps is stored beneath a top-level key, the `-etcd-prefix`, which defaults to `/apps`. So here's how you might route all requests for `myapp1.mydomain.com` to a single HTTP listener:

    etcdctl set /apps/tomcatapp/vhost    "myapp1.mydomain.com"
    etcdctl set /apps/tomcatapp/endpoint "127.0.0.1:8080"

In this example, `tomcatapp` is treated as the ID for a backend web app, because it is directly beneath the `-etcd-prefix` key. **Each etcd node immediately beneath the `-etcd-prefix` defines a backend web app; the name of the node becomes the app ID.** In the above example, a configuration file named `tomcatapp-etcd-vhost.conf` would be generated in the `-nginx-dir`.

Besides the node names directly beneath the `-etcd-prefix`, the only other values that are treated as significant are any values associated with the keys `vhost` or `endpoint`. These get associated with the corresponding app. None of the intervening node names matter, so you can bury the `vhost` and `endpoint` values as deep as you like in the etcd hierarchy. Here's an example with multiple backends and multiple virtual hostnames:

    etcdctl set /apps/railsapp/host1/vhost        "myapp2.mydomain.com"
    etcdctl set /apps/railsapp/host2/vhost        "myapp2.*"
    etcdctl set /apps/railsapp/listener1/endpoint "127.0.0.1:3000"
    etcdctl set /apps/railsapp/listener2/endpoint "127.0.0.1:3001"

4) Run `etcd-nginx-config` with the correct config values:

    etcd-nginx-config --help
    Usage of etcd-nginx-config:
      -etcd-hosts="http://127.0.0.1:4001": etcd server URLs
      -etcd-prefix="/apps": top-level etcd key for apps
      -nginx-dir="/etc/nginx/conf.d": output dir for nginx virtual host files
      -nginx-pid="/var/run/nginx.pid": location of the nginx PID file
      -version=false: print version and exit


  _Note that if nginx is running as root, then `etcd-nginx-config` will also need to run as root, or it cannot signal nginx to reload the configration files._

Once running, `etcd-nginx-config` will watch etcd for any changes beneath the `-etd-prefix`. When changes occur, all config files are re-rendered, and a `SIGHUP` is sent to nginx. This causes the web server to reload its configuration without dropping any connections.


----------------
Known Limitations / Bugs
----------------
Interprocess signalling with `SIGHUP` is not supported on Windows, so the auto-reload feature won't work there.


----------------
Who is it? (Contribution)
----------------
This software was created by Benton Roberts _(benton@bentonroberts.com)_

The project is still in its early stages. Pull requests are welcome! Here's how to get started:

1) Install [Go][2].

2) Fetch the project code:

    go get github.com/benton/etcd-nginx-config

3) Run the tests:

    cd $GOPATH/src/github.com/benton/etcd-nginx-config
    go test


--------
[1]: https://github.com/coreos/etcd
[2]: http://golang.org/doc/install

[linux-latest]: https://github.com/benton/etcd-nginx-config/releases/download/0.1.1/etcd-nginx-config-linux-x86_64.tgz
[osx-latest]: https://github.com/benton/etcd-nginx-config/releases/download/0.1.1/etcd-nginx-config-osx.tgz
[windows-latest]: https://github.com/benton/etcd-nginx-config/releases/download/0.1.1/etcd-nginx-config-win64.zip
