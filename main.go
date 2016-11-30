package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"sync"

	"github.com/julienschmidt/httprouter"
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
	routes    *httprouter.Router
	waitGroup sync.WaitGroup
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	t := time.Now()
	s := strftime.Strftime(&t, "%H:%M:%S")
	fmt.Printf("[%s] \n", s)
	fmt.Printf("Number Of Cores on server: %d\n", runtime.GOMAXPROCS(runtime.NumCPU()))

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
		// // Remove Temp, comipled stylesheets
		os.RemoveAll("./public/assets/stylesheets/")
		helpers.ShutDown.Printf("Server Shutdown by User.")
		// Explicitly call for system exit this is more graceful
		os.Exit(0)
	}()
	listen := fmt.Sprintf("%s:%s", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	go helpers.Logger.Fatal(http.ListenAndServe(listen, routes))

}

/////////////////////////////////////////////////////////////
//Compile Asset Pipeline
/////////////////////////////////////////////////////////////

func compileAssets() {
	compiled := make(chan bool)

	go func() {

		err := exec.Command(
			"sass",
			"--watch",
			"./app/assets/stylesheets/:./public/assets/stylesheets/", "--style", "compressed").Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		compiled <- true
		close(compiled)
	}()

	go func() {

		applicationFiles, _ := filepath.Glob("./app/assets/typescripts/application/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/application/", "--watch"}, applicationFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}
		compiled <- true
		close(compiled)
	}()
	go func() {

		blogFiles, _ := filepath.Glob("./app/assets/typescripts/blog/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/blog/", "--watch"}, blogFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		compiled <- true
		close(compiled)
	}()
	go func() {

		exampleFiles, _ := filepath.Glob("./app/assets/typescripts/examples/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/examples/", "--watch"}, exampleFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		compiled <- true
		close(compiled)
	}()
	go func() {

		userFiles, _ := filepath.Glob("./app/assets/typescripts/users/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/users/", "--watch"}, userFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		compiled <- true
		close(compiled)
	}()
	go func() {

		files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		compiled <- true
		close(compiled)
	}()

}
