package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// GetRoutes does some shit. Here's a java example.
// public Router getRoutes() { Router router = new httprouter.Router(); return router; }
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", controllers.ApplicationIndex)
	router.GET("/example", controllers.ExampleIndex)
	router.GET("/about_me", controllers.AboutIndex)
	router.POST("/register", controllers.UserCreate)
	router.GET("/assets/stylesheets/*sheet", helpers.HandleScssRequest)
	router.GET("/assets/javascripts/*sheet", helpers.HandleKobraRequest)
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	// router.GET("/hello/:name", Hello)
	return router
}
