// +build plan9 windows

package main

import (
	"errors"
	"os"
)

// Returns an error indicating no SIGHUP is available on this platform
func sighup(process *os.Process) error {
	return errors.New("SIGHUP not supported on this platform")
}
