package helpers

import "io"

// layout, err := jade.ParseFile("./app/views/layouts/application.pug")
// if err != nil {
// 	Logger.Printf("\nParseFile error: %v", err)
// }

// layout := handleErr("ParseFile error: %v\n", jade.ParseFile("./app/views/layouts/application.pug"))
//Close is a closer
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		Logger.Println(err)
	}
}
