package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
)

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// GetRoutes does some shit. Here's a java example.
// public Router getRoutes() { Router router = new httprouter.Router(); return router; }
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	// Pages and shit
	router.GET("/", controllers.ApplicationIndex)
	router.GET("/example", controllers.ApplicationExamples)
	router.GET("/about", controllers.ApplicationAbout)

	// Users
	router.GET("/users", controllers.UserIndex)
	router.GET("/users/:name", controllers.UserShow)
	router.POST("/register", controllers.UserNew)
	// Session Management
	router.GET("/signin", controllers.SessionNew)
	router.POST("/signin", controllers.SessionCreate)
	router.GET("/signout", controllers.SessionDestroy)

	// Blog routes
	router.GET("/blog", controllers.BlogIndex)

	router.ServeFiles("/static/*filepath", http.Dir("public"))
	return router
}
