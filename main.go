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

	mgo "gopkg.in/mgo.v2"

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
	session   *mgo.Session
)

func init() {
	config.Load()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if config.Data.Env != "production" {
		compileAssets()
		fmt.Println("Getting routes")
		t := time.Now()
		s := strftime.Strftime(&t, "%H:%M:%S")
		fmt.Printf("[%s] \n", s)
		fmt.Printf("Number Of Cores on server: %d\n", runtime.GOMAXPROCS(runtime.NumCPU()))
	}
	// mem, err := memcache.New("127.0.0.1:11211")
	// if err != nil {
	// 	helpers.Logger.Fatalln(err)
	// }
	// mem.Add(&memcache.Item{Key: "foo", Value: []byte("testmemcache setup")})
	// mc1, err := mem.Get("foo")
	// if err != nil {
	// 	helpers.Logger.Println(err)
	// }
	// fmt.Println(string(mc1.Value))
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

		if config.Data.Env != "production" {
			// Removes Temp, compiled JS files
			os.RemoveAll("./public/assets/scripts/")
			// // Remove Temp, comipled stylesheets
			os.RemoveAll("./public/assets/stylesheets/")
		}

		helpers.ShutDown.Printf("Server Shutdown by User.")
		// Explicitly call for system exit this is more graceful
		os.Exit(0)
	}()
	listen := fmt.Sprintf("%s:%d", config.Data.Host, config.Data.Port)
	fmt.Printf("Listening on %s\n", listen)
	go helpers.Logger.Fatal(http.ListenAndServe(listen, routes))

}

/////////////////////////////////////////////////////////////
//Compile Asset Pipeline
/////////////////////////////////////////////////////////////

func compileAssets() {
	compiled := make(chan bool)

	go func() {
		fmt.Println("Sass Assets")
		err := exec.Command(
			"sass",
			"--watch",
			"./app/assets/stylesheets/:./public/assets/stylesheets/", "--style", "compressed").Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		fmt.Println("Sass Assets Compiled")
		compiled <- true
		close(compiled)
	}()

	go func() {
		fmt.Println("Typscripts Assets")
		applicationFiles, _ := filepath.Glob("./app/assets/typescripts/application/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/application/", "--watch"}, applicationFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}
		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		blogFiles, _ := filepath.Glob("./app/assets/typescripts/blog/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/blog/", "--watch"}, blogFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		exampleFiles, _ := filepath.Glob("./app/assets/typescripts/examples/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/examples/", "--watch"}, exampleFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		userFiles, _ := filepath.Glob("./app/assets/typescripts/users/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/users/", "--watch"}, userFiles...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()
	go func() {
		fmt.Println("Typscripts Assets")
		files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
		err := exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
		if err != nil {
			helpers.Logger.Println(err)
			return
		}

		fmt.Println("Typescript Assets Compiled")
		compiled <- true
		close(compiled)
	}()

}
