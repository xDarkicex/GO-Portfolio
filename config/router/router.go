package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
)

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// GetRoutes . Here's a java example.
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
	router.GET("/register", controllers.RegisterNew)
	router.POST("/register", controllers.UserCreate)

	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////

	router.GET("/signin", controllers.SessionNew)
	router.POST("/signin", controllers.SessionCreate)
	router.GET("/signout", controllers.SessionDestroy)

	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////

	router.GET("/posts", controllers.BlogIndex)         // index
	router.GET("/posts/new", controllers.BlogPostNew)   // new 		To make a new Post
	router.POST("/posts", controllers.BlogNew)          // create	To actually throw it in the database
	router.GET("/post/:url", controllers.BlogShow)      // show		Show a specific post
	router.GET("/post/:url/edit", controllers.BlogEdit) // edit		Edit a specific post
	// router.DELETE("/posts/:title", ) 				   // destroy Destroy a specific post
	// router.PATCH("/posts/:title", )	                   // update Update a specific post
	///////////////////////////////////////////////////////////
	// Static routes
	///////////////////////////////////////////////////////////

	router.GET("/post/:url/images/:imageID", controllers.BlogImage)
	router.GET("/posts/search/:searchTerm", controllers.BlogSearch)
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	return router
}
