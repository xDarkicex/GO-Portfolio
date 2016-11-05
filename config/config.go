package config

import "os"

// Version number
var Version = "1.8.1"

// Port is the hole we should use.
var Port = "3000"

// Host is the ip we use to listen on.
var Host = "0.0.0.0"

//ENV is the enviroment for server
var ENV = os.Getenv("ENV")

// EMAIL address ..
var EMAIL = os.Getenv("EMAIL")

// SMTPHOST smtp ...
var SMTPHOST = os.Getenv("SMTPHOST")

// SMTPPORT ..
var SMTPPORT = os.Getenv("SMTPPORT")

// SMTPPASSWORD ...
var SMTPPASSWORD = os.Getenv("SMTPPASSWORD")

// var SMTPPASSWORD = "Vh402152Go!"

//ErrorFile error files
var ErrorFile = "error.log"

//Verbose is for turning error logs of and on.
var Verbose = 1
