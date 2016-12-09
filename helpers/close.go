package helpers

import "io"

//Close is a closer
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		Logger.Println(err)
	}
}
