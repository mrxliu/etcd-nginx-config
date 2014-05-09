Etcd-Nginx Web Router
================
Writes NGinx reverse-proxy config files for multiple backends, based on data in [etcd][1].

  *ALPHA VERSION - needs more tests*


----------------
What is it?
----------------
This program connects to the etcd configuration service and reads information about web apps. It builds a list of web apps, including each app's HTTP endpoints and virtual host names, then exports one Nginx reverse-proxy configuration file for each app.

The software then watches etcd for changes, updates the files as needed, and signals NGinx (via PID file) to reload its configuration.


----------------
Why is it?
----------------
* Provides dynamic, virtual-host-based HTTP routing based on minimal config


----------------
Where is it? (Installation)
----------------
Download the latest binary for your OS:

Put it wherever you like. Somewhere in your $PATH might ne handy.


----------------
How is it [done]? (Usage)
----------------


----------------
Known Limitations / Bugs
----------------


----------------
Who is it? (Contribution)
----------------
This software was created by Benton Roberts _(benton@bentonroberts.com)_

The project is still in its early stages.

1) Install [Go][2].

2) Fetch the project code:

    go get https://github.com/benton/etcd-nginx-config

3) Run the tests:

    cd $GOPATH/src/github.com/benton/etcd-nginx-config
    go build

4) Run the tests:

    go test


--------
[1]: https://github.com/coreos/etcd
[2]: http://golang.org/doc/install