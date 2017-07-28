package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	mgo "gopkg.in/mgo.v2"

	"sync"

	"github.com/weidewang/go-strftime"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/config/router"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

/////////////////////////////////////////////////////////////
// Declared Var
/////////////////////////////////////////////////////////////

var (
	routes    *helpers.Server
	waitGroup sync.WaitGroup
	session   *mgo.Session
)

func init() {
	config.Load()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if config.Data.Env != "production" {
		helpers.CompileAssets()
		fmt.Println("Getting routes")
		t := time.Now()
		s := strftime.Strftime(&t, "%H:%M:%S")
		fmt.Printf("[%s] \n", s)
		fmt.Printf("Number Of Cores on server: %d\n", runtime.GOMAXPROCS(runtime.NumCPU()))
	}

	//LOG pid ID
	helpers.Pidder()
	// Create routes

	// Dial mongo Datastore
	err := db.Dial()
	if err != nil {
		log.Fatal(err)
	}
	session = db.Session()
}

func main() {
	go func() {
		for true {
			time.Sleep(5 * time.Second)
			helpers.FlushLog()
		}
	}()
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

		if config.Data.Env != "production" {
			// Removes Temp, compiled JS files
			os.RemoveAll("./public/assets/scripts/")
			// // Remove Temp, comipled stylesheets
			os.RemoveAll("./public/assets/stylesheets/")
		}
		helpers.Logger.Printf("Server Shutdown.")
		// Explicitly call for system exit this is more graceful
		os.Exit(0)
	}()
	listen := fmt.Sprintf("%s:%d", config.Data.Host, config.Data.Port)
	log.Printf("Listening on %s\n", listen)
	router.GetRoutes().Serve(":3000")

}
