package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"

	"github.com/xDarkicex/PortfolioGo/config"
)

var routes *httprouter.Router
var session *mgo.Session

//Session is a session to DB

func init() {
	fmt.Println("Getting routes")
	routes = config.GetRoutes()
	session := config.Dial()

	session.SetMode(mgo.Monotonic, true)
}

func main() {

	defer session.Close()
	listen := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, routes))
}
