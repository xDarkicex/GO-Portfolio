package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/config/router"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	mgo "gopkg.in/mgo.v2"
)

/////////////////////////////////////////////////////////////
// Declared Var
/////////////////////////////////////////////////////////////

var (
	routes  *httprouter.Router
	session *mgo.Session
)

func init() {
	compileAssets()
	fmt.Println("Getting routes")
	routes = router.GetRoutes()
	err := db.Dial()
	if err != nil {
		log.Fatal(err)
	}
	session = db.Session()
}

func main() {
	// create self calling go routine
	go func() {
		interruptChannel := make(chan os.Signal, 0)
		// look for system interruptions
		signal.Notify(interruptChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		// lock lower code untel interruptChannel receives signal
		<-interruptChannel

		// Other cleanup tasks
		fmt.Println("Closing connection")
		fmt.Println("Saving session")
		// Accually Close DB session, maintain DATA integrity
		session.Close()
		// Removes Temp, compiled JS files
		os.RemoveAll("./public/assets/scripts/")
		// Remove Temp, comipled stylesheets
		os.RemoveAll("./public/assets/stylesheets/")
		helpers.Logger.Println("Server Shutdown!!")

		// Explicitly call for system exit this is more graceful
		os.Exit(0)
	}()

	// ServerTitle title for server
	var ServerTitle = "Golang Server version 1.7.3"
	fmt.Print("\033]0;" + ServerTitle + "\007")
	listen := fmt.Sprintf("%s:%s", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	helpers.Logger.Fatal(http.ListenAndServe(listen, routes))
}

/////////////////////////////////////////////////////////////
//Compile Asset Pipeline
/////////////////////////////////////////////////////////////

func compileAssets() {
	err := exec.Command(
		"sass",
		"--watch",
		"./app/assets/stylesheets/:./public/assets/stylesheets/").Start()
	if err != nil {
		helpers.Logger.Println(err)
		return
	}

	applicationFiles, _ := filepath.Glob("./app/assets/typescripts/application/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/application/", "--watch"}, applicationFiles...)...).Start()
	if err != nil {
		helpers.Logger.Println(err)
		return
	}
	blogFiles, _ := filepath.Glob("./app/assets/typescripts/blog/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/blog/", "--watch"}, blogFiles...)...).Start()
	if err != nil {
		helpers.Logger.Println(err)
		return
	}
	userFiles, _ := filepath.Glob("./app/assets/typescripts/users/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/users/", "--watch"}, userFiles...)...).Start()
	if err != nil {
		helpers.Logger.Println(err)
		return
	}
	files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
	if err != nil {
		helpers.Logger.Println(err)
		return
	}
}
