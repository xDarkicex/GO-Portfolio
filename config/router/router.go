package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
)

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// GetRoutes. Here's a java example.
// public Router getRoutes() { Router router = new httprouter.Router(); return router; }
func GetRoutes() *httprouter.Router {
	router := httprouter.New()

	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////
	router.GET("/", controllers.ApplicationIndex)
	router.GET("/examples", controllers.ApplicationExamples)
	router.GET("/about", controllers.ApplicationAbout)
	router.GET("/404", controllers.Error404)
	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////

	router.GET("/users", controllers.UserIndex)
	router.GET("/users/:name", controllers.UserShow)
	router.POST("/register", controllers.UserNew)
	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////
	router.GET("/signin", controllers.SessionNew)
	router.POST("/signin", controllers.SessionCreate)
	router.GET("/signout", controllers.SessionDestroy)
	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////
	router.GET("/blog", controllers.BlogIndex)
	router.GET("/blog/new", controllers.BlogNew)
	// router.POST("/blog/new/create", controllers.BlogCreate)
	router.GET("/blog/edit/:title", controllers.BlogEdit)
	///////////////////////////////////////////////////////////
	// Static routes
	///////////////////////////////////////////////////////////
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	return router
}
