package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func init() {

}

// Data struct for config.Data
var Data config

// config is for config.Data
type config struct {
	Version string
	Port    int
	Host    string
	Secret  string
	Verbose bool
	Env     string
	Email   string
	SMTP    smtp
}

// SMTP for smtp settings
type smtp struct {
	Host     string
	Port     int
	Password string
}

//Load loads config file
func Load() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err, config)
	}
	// fmt.Println(config)
	err = json.Unmarshal(config, &Data)
	if err != nil {
		fmt.Println(err)
	}
	if Data.Env != "production" {
		fmt.Printf("\n")
		fmt.Println(Data)
		fmt.Printf("\n")
	}

}
