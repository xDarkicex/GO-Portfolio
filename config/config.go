package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func init() {
	Load()
}

type Config struct {
	Version   string
	Port      int
	Host      string
	Errorfile string
	Verbose   bool
	Env       string
	Email     string
	Stmp      struct {
		Host     string
		Port     int
		Password string
	}
}

var configed Config

//Load loads config file
func Load() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err, config)
	}
	// fmt.Println(config)
	err = json.Unmarshal(config, &configed)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(configed)
}

// Version number
var Version = configed.Version

// // Port is the hole we should use.
var Port = configed.Port

// // Host is the ip we use to listen on.
var Host = configed.Host

// //ENV is the enviroment for server
var ENV = configed.Env

// // EMAIL address ..
var EMAIL = configed.Email

// // SMTPHOST smtp ...
var SMTPHOST = configed.Stmp.Host

// // SMTPPORT ..
var SMTPPORT = configed.Stmp.Port

// // SMTPPASSWORD ...
var SMTPPASSWORD = configed.Stmp.Password

// //ErrorFile error files
var ErrorFile = configed.Errorfile

//Verbose is for turning error logs of and on.
var Verbose = configed.Verbose
