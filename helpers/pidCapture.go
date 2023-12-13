package helpers

import (
	"fmt"
	"os"
	"syscall"
)

const (
	pidFileName      = "pid.txt"
	filePermission   = os.O_RDWR | os.O_CREATE
	filePermissionRW = 0666
)

// Pidder captures the program's active PID and writes it to a file.
func Pidder() error {
	SPID := syscall.Getpid()

	file, err := os.OpenFile(pidFileName, filePermission, filePermissionRW)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer Close(file) // Close is a helper function to safely close an io.Closer

	_, err = file.WriteString(fmt.Sprintf("%d", SPID))
	if err != nil {
		return fmt.Errorf("failed to write PID to file: %w", err)
	}

	return nil
}
