package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"io/ioutil"
	"os"
	"regexp"
)

// Makes this app's conf files less likely to collide with others'
const NGINX_SUFFIX = "-etcd-vhost.conf"

// Returns the name of an Nginx configuration file for the given web app
func confFileName(app *Webapp) string {
	return fmt.Sprintf("%s%s", app.ID, NGINX_SUFFIX)
}

// Queries etcd, then writes out an NGiNX configuration file for each app.
// Finally, calls removeOldNginxFiles to remove outdated conf files.
func writeNginxFiles(client *etcd.Client, config *Config) error {
	apps, err := readAppsFromEtcd(client, config.Prefix)
	if err == nil {
		for _, app := range apps {
			outPath := fmt.Sprintf("%s/%s", config.Outdir, confFileName(&app))
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
		err = removeOldNginxFiles(config, apps)
	}
	return err
}

// Deletes configuration files for apps that are no longer in etcd
func removeOldNginxFiles(config *Config, apps []Webapp) error {
	webappFile, _ := regexp.Compile(fmt.Sprintf("%s$", NGINX_SUFFIX))
	validFiles := map[string]bool{} // create a map of valid filenames
	for _, app := range apps {
		validFiles[confFileName(&app)] = true
	}
	files, err := ioutil.ReadDir(config.Outdir) // list all existing conf files
	if err == nil {
		for _, f := range files {
			if webappFile.Match([]byte(f.Name())) {
				if _, valid := validFiles[f.Name()]; !valid {
					outpath := fmt.Sprintf("%s/%s", config.Outdir, f.Name())
					fmt.Println("WARNING: deleting file", outpath)
					err = os.Remove(outpath)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}
