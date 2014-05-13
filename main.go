package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
	"time"
)

const VERSION = "0.1.3" // Package version

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
	signalNginx(&config)

	fmt.Println("Watching etcd for changes on prefix:", config.Prefix)
	etcdQueue := make(chan *etcd.Response)
	go client.Watch(config.Prefix, 0, true, etcdQueue, nil)
	for {
		response := <-etcdQueue
		if response != nil {
			fmt.Println("Got changes from node", response.Node.Key)
			err := writeNginxFiles(client, &config)
			if err != nil {
				fmt.Println("WARNING: error writing nginx config files:", err)
			} else {
				signalNginx(&config)
			}
		} else {
			fmt.Println("WARNING: etcd missing? Sleeping", ETCD_DELAY, "seconds")
			time.Sleep(5 * time.Second)           // wait for etcd to recover
			etcdQueue = make(chan *etcd.Response) // try to reconnect
			go client.Watch(config.Prefix, 0, true, etcdQueue, nil)
		}
	}
}
