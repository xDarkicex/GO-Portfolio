package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func init() {

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

var Configed Config

//Load loads config file
func Load() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err, config)
	}
	// fmt.Println(config)
	err = json.Unmarshal(config, &Configed)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Configed)
	Version = Configed.Version
	Port = Configed.Port
	Host = Configed.Host
	ENV = Configed.Env
	EMAIL = Configed.Email
	SMTPHOST = Configed.Stmp.Host
	SMTPPORT = Configed.Stmp.Port
	SMTPPASSWORD = Configed.Stmp.Password
	ErrorFile = Configed.Errorfile
	Verbose = Configed.Verbose
}

// Version number
var Version string

// Port is the hole we should use.
var Port int

// Host is the ip we use to listen on.
var Host string

// ENV is the enviroment for server
var ENV string

// EMAIL address ..
var EMAIL string

// SMTPHOST smtp ...
var SMTPHOST string

// SMTPPORT ...
var SMTPPORT int

// SMTPPASSWORD ...
var SMTPPASSWORD string

// ErrorFile error files
var ErrorFile string

//Verbose is for turning error logs of and on.
var Verbose bool
