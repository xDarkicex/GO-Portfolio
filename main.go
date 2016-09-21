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

//Session is a session to DB
var Session *mgo.Session

func init() {
	fmt.Println("Getting routes")
	routes = config.GetRoutes()
}

func main() {

	Session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer Session.Close()
	Session.SetMode(mgo.Monotonic, true)
	listen := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, routes))
}
