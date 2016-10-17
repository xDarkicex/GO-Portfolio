package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
)

// GetRoutes func to setup all routes
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	// Might switch to iris if I can figure it out...
	// func GetRoutes() *iris.Context {
	// 	router := iris.New()
	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////

	application := controllers.Application{}
	router.GET("/", application.Index)
	router.GET("/about", application.About)

	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////

	users := controllers.Users{}
	router.GET("/users", users.Index)
	router.GET("/users/:name", users.Show)
	router.GET("/register", users.New)
	router.POST("/register", users.Create)

	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////

	sessions := controllers.Sessions{}
	router.GET("/signin", sessions.New)
	router.POST("/signin", sessions.Create)
	router.GET("/signout", sessions.Destroy)

	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////

	blog := controllers.Blog{}
	blog.Globals.Count++
	router.GET("/posts", blog.Index)          // index
	router.GET("/posts/new", blog.New)        // new 		To make a new Post
	router.POST("/posts", blog.Create)        // create	To actually throw it in the database
	router.GET("/post/:url", blog.Show)       // show		Show a specific post
	router.POST("/post/:url", blog.Update)    // update Update a specific post
	router.GET("/post/:url/edit/", blog.Edit) // So Form for updating a specific post I maybe should mke a new method to make a more tailored form
	router.GET("/post/:url/images/:imageID", blog.Image)

	// example := controllers.Example{}

	///////////////////////////////////////////////////////////
	// Static routes
	///////////////////////////////////////////////////////////

	router.ServeFiles("/static/*filepath", http.Dir("public"))
	return router
}
