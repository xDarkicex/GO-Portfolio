package main

import (
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/config/router"
	"github.com/xDarkicex/PortfolioGo/db"
)

var routes *httprouter.Router

var session *mgo.Session

func init() {
	fmt.Println("Getting routes")
	routes = router.GetRoutes()
	db.Dial()
}

func main() {

	defer session.Close()
	listen := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, routes))
}
