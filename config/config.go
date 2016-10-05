package config

import "net"

// Version number
var Version = "1.0.0"

// Port is the hole we should use.
var Port = 3000

// DynamicPort scan for availible port
func DynamicPort() int {
	///////////////////////////////////
	// Returning Zero because of this
	//////////////////////////////////
	var Port int
	/////////////////////////////////
	for i := 3000; i < 8080; i++ {
		_, err := net.Listen("tcp", string(i))
		if err == nil {
			Port := i
			///////////////////////////////////
			//Want to pass this to next return
			//////////////////////////////////
			return Port
			//////////////////////////////////
		}
	}
	return Port
}

// Host is the ip we use to listen on.
var Host = "0.0.0.0"

//ENV is the enviroment for database
var ENV = "development"
