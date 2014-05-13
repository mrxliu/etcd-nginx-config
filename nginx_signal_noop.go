// +build plan9 windows

package main

import (
	"fmt"
)

// Returns an error indicating no SIGHUP is available on this platform
func signalNginx(config *Config) {
	fmt.Println("WARNING: SIGHUP not supported on this platform")
}
