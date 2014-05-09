package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
)

func writeNginxFiles(client *etcd.Client, prefix string) error {
	apps, err := readAppsFromEtcd(client, prefix)
	if err == nil {
		for _, app := range apps {
			outPath := fmt.Sprintf("conf.d/%s-vhost.conf", app.ID)
			fmt.Printf("Writing file: %s...\n", outPath)
			var outFile *os.File
			outFile, err = os.Create(outPath)
			if err == nil { // close file on exit and check for returned error
				defer func() {
					if err := outFile.Close(); err != nil {
						panic(err)
					}
				}()
				err = app.Render(outFile)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}
