package helpers

import (
	"fmt"
	"os"
	"syscall"
)

// Pidder function captures the programs active pi for later use inside AWS, very usefull to know programs pid ID
func Pidder() {
	SPID := syscall.Getpid()
	file, _ := os.OpenFile("pid.txt", os.O_RDWR|os.O_CREATE, 0666)
	file.WriteString(fmt.Sprintf("%d", SPID))
}
