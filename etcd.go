package main

import (
	"github.com/coreos/go-etcd/etcd"
	"strings"
)

// Retrieves all unique values for a given key from a node and its children.
func getEtcdValues(node *etcd.Node, key string) []string {
	valuesMap := map[string]bool{} // (track vaules in a map, for uniqueness)
	if node.Value != "" {
		if lastKey(node.Key) == key {
			valuesMap[node.Value] = true
		}
	} else {
		for _, subNode := range node.Nodes {
			subValues := getEtcdValues(subNode, key)
			for _, subValue := range subValues {
				valuesMap[subValue] = true
			}
		}
	}
	values := []string{} // compute just the keys of the valuesMap
	for value, _ := range valuesMap {
		values = append(values, value)
	}
	return values
}

// Returns the last component of an etcd key
func lastKey(key string) string {
	components := strings.Split(key, "/")
	return strings.Join(components[len(components)-1:len(components)], "/")
}
