package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/config/router"
	"github.com/xDarkicex/PortfolioGo/db"
	mgo "gopkg.in/mgo.v2"
)

var routes *httprouter.Router

var session *mgo.Session

func init() {
	compileAssets()
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

func compileAssets() {
	err := exec.Command(
		"sass",
		"--watch",
		"./app/assets/stylesheets/:./public/assets/stylesheets/").Start()
	if err != nil {
		fmt.Println(err)
	}
	files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
	if err != nil {
		fmt.Println(err)
	}
}
