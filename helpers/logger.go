package helpers

import (
	"log"
	"os"
)

// Logger is a helpper method to print out a more useful error message
var Logger = log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile)
