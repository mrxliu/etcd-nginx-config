package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
	"syscall"
)

const VERSION = `0.1.1`

func main() {
	config := newConfig()

	// write initial nginx config to filesystem
	fmt.Println("Connecting to etcd on hosts:", config.Hosts)
	client := etcd.NewClient([]string{config.Hosts})
	err := writeNginxFiles(client, &config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Watching etcd for changes on prefix:", config.Prefix)
	etcdQueue := make(chan *etcd.Response)
	go client.Watch(config.Prefix, 0, true, etcdQueue, nil)
	for {
		<-etcdQueue
		err := writeNginxFiles(client, &config)
		if err != nil {
			fmt.Println(err.Error())
		}
		nginx, err := nginxProcess(&config)
		if err == nil {
			fmt.Println("Sending SIGHUP to NGiNX process:", nginx.Pid)
			err = nginx.Signal(syscall.SIGHUP)
			if err != nil {
				fmt.Println("WARNING: Can't signal NGiNX:", err)
			}
		} else {
			fmt.Println("WARNING: Can't find NGiNX:", err)
		}
	}
}
