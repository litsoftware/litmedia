package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

func CheckAndCreatePidFile(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		pidString, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Cannot read PID file: %v", err)
			return false
		}

		pid, err := strconv.Atoi(string(pidString))
		if err != nil {
			fmt.Printf("Cannot interpret contents of PID file: %v", err)
			return false
		}

		if pid == os.Getpid() {
			fmt.Println("Found existing PID file matching current pid")
			return true
		}

		// Try sending a signal to the process to see if it is still running
		process, err := os.FindProcess(pid)
		if err == nil {
			err = process.Signal(syscall.Signal(0))
			if (err == nil) || (err == syscall.EPERM) {
				fmt.Printf("Existing process running on PID %d. Exiting (my pid = %d)", pid, os.Getpid())
				return false
			}
		}
	}

	pidfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Cannot write PID file: %v", err)
		return false
	}
	_, _ = fmt.Fprintf(pidfile, "%v", os.Getpid())
	_ = pidfile.Close()
	return true
}

func RemovePidFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		fmt.Printf("Failed to remove PID file: %v\n", err)
	}

	return err
}
