package main

import (
	"flag"
)

// Represents this app's possible configuration values
type Config struct {
	Hosts   string
	Prefix  string
	Outdir  string
	PIDfile string
}

// Generates and returns a new Config based on the command-line
func newConfig() Config {
	var (
		hosts = flag.String("etcd-hosts",
			"http://127.0.0.1:4001",
			"etcd server URLs")
		prefix = flag.String("etcd-prefix",
			"/apps",
			"top-level etcd key for apps")
		outdir = flag.String("nginx-dir",
			"/etc/nginx/conf.d",
			"output dir for nginx virtual host files")
		pidfile = flag.String("nginx-pid",
			"/var/run/nginx.pid",
			"location of the nginx PID file")
	)
	flag.Parse()
	return Config{
		Hosts:   *hosts,
		Prefix:  *prefix,
		Outdir:  *outdir,
		PIDfile: *pidfile,
	}
}
