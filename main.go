package main

import (
	"fmt"
	"io"
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

/////////////////////////////////////////////////////////////
// Declared Var
/////////////////////////////////////////////////////////////

var routes *httprouter.Router
var session *mgo.Session

// var Store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	compileAssets()
	fmt.Println("Getting routes")
	routes = router.GetRoutes()
	db.Dial()
}

// Close Wrapper not yet finished
func _close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
	session.Close()
}
func main() {
	// go func() {
	// 	interruptChannel := make(chan os.Signal, 0)
	// 	signal.Notify(interruptChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	// 	<-interruptChannel

	// 	// Other cleanup tasks
	// 	// Dont for get this is not fucntion 100% correct.
	// 	// _close()
	// 	session.Close()
	// 	// Other cleanup tasks

	// 	os.Exit(0)
	// }()

	defer session.Close()
	listen := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, routes))
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
		fmt.Println(err)
	}
	files, _ := filepath.Glob("./app/assets/typescripts/*.ts")
	err = exec.Command("tsc", append([]string{"--outDir", "./public/assets/scripts/", "--watch"}, files...)...).Start()
	if err != nil {
		fmt.Println(err)
	}
}
