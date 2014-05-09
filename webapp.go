package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"io"
	"strings"
	"text/template"
)

// Represents some hostname patterns that route to a given Array of Listeners
type Webapp struct {
	ID        string
	Endpoints []string
	VHosts    []string
}

// Returns a space-delimited string containing the VHosts
func (app *Webapp) HostList() string {
	return strings.Join(app.VHosts, " ")
}

// Renders an NGinx VHost template for the given app to the Writer output
func (app *Webapp) Render(output io.Writer) error {
	tmpl, err := template.New(app.ID).Parse(DEFAULT_VHOST_TEMPLATE)
	if err == nil {
		err = tmpl.Execute(output, app)
	}
	return err
}

// Returns all Webapps from the given Client beneath prefix
func readAppsFromEtcd(client *etcd.Client, prefix string) ([]Webapp, error) {
	response, err := client.Get(prefix, true, true)
	if err != nil {
		return []Webapp{}, err
	}
	apps := []Webapp{}
	for _, appNode := range response.Node.Nodes {
		appPrefix := fmt.Sprintf("%s/", prefix)
		appName := strings.Replace(appNode.Key, appPrefix, "", 1)
		app := Webapp{
			ID:        appName,
			Endpoints: getEtcdValues(&appNode, "endpoint"),
			VHosts:    getEtcdValues(&appNode, "vhost"),
		}
		apps = append(apps, app)
	}
	return apps, err
}
