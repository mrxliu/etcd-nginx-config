package main

import (
	"io"
	"os"
	"strconv"
)

// Returns the process ID of NGiNX read from the PID file
func nginxPID(config *Config) (int, error) {
	pid := -1
	pidFile, err := os.Open(config.PIDfile)
	if err == nil {
		defer func() { // don't forget to close the file
			if err := pidFile.Close(); err != nil {
				panic(err)
			}
		}()
		buf := make([]byte, 1024) // read up to 1K from the PID file
		bytes, err := pidFile.Read(buf)
		if err == nil || err == io.EOF {
			pid, err = strconv.Atoi(string(buf[0 : bytes-1])) // convert to int
			if err != nil {
				return -1, err
			}
		}
	}
	return pid, err
}

// Returns a pointer to the NGiNX process specified by the PID file, if running.
func nginxProcess(config *Config) (*os.Process, error) {
	pid, err := nginxPID(config)
	if err != nil {
		return nil, err
	}
	return os.FindProcess(pid)
}
