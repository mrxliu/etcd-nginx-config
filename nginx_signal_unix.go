// +build !plan9
// +build !windows

package main

import (
	"fmt"
	"syscall"
)

// sends SIGHUP to a process
func signalNginx(config *Config) {
	nginx, err := nginxProcess(config)
	if err == nil {
		fmt.Println("Sending SIGHUP to nginx process:", nginx.Pid)
		err = nginx.Signal(syscall.SIGHUP)
		if err != nil {
			fmt.Println("WARNING: Can't signal nginx:", err)
		}
	} else {
		fmt.Println("WARNING: Can't find nginx:", err)
	}
}
