package helpers

import (
	"io"
	"log"
)

// Close Wrapper not yet finished
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}
