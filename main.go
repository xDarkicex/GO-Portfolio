package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/xDarkicex/PortfolioGo/config"
)

var routes *httprouter.Router

func init() {
	fmt.Println("Getting routes")
	routes = config.GetRoutes()
}
func main() {
	listen := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, routes))
}
