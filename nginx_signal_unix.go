// +build !plan9
// +build !windows

package main

import (
	"os"
	"syscall"
)

// sends SIGHUP to a process
func sighup(process *os.Process) error {
	return process.Signal(syscall.SIGHUP)
}
