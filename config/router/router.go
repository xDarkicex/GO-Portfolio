package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
	"github.com/xDarkicex/PortfolioGo/config/gzip"
)

// GetRoutes func to setup all routes
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////

	application := controllers.Application{}
	router.GET("/", gzip.Middleware(application.Index, false))
	router.GET("/about", gzip.Middleware(application.About, false))
	router.POST("/contact", gzip.Middleware(application.Contact, false))

	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////

	users := controllers.Users{}
	router.GET("/users", gzip.Middleware(users.Index, false))
	router.GET("/users/:name", gzip.Middleware(users.Show, false))
	router.GET("/register", gzip.Middleware(users.New, false))
	router.POST("/users/:name", gzip.Middleware(users.Update, false))
	router.GET("/users/:name/edit/", gzip.Middleware(users.Edit, false))
	router.POST("/register", gzip.Middleware(users.Create, false))
	router.GET("/users/:name/images/:imageID", gzip.Middleware(users.Image, false))

	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////

	sessions := controllers.Sessions{}
	router.GET("/signin", gzip.Middleware(sessions.New, false))
	router.POST("/signin", gzip.Middleware(sessions.Create, false))
	router.GET("/signout", gzip.Middleware(sessions.Destroy, false))

	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////

	blog := controllers.Blog{}
	router.GET("/posts", gzip.Middleware(blog.Index, false))          // index
	router.GET("/posts/new", gzip.Middleware(blog.New, false))        // new 		To make a new Post
	router.POST("/posts", gzip.Middleware(blog.Create, false))        // create	To actually throw it in the database
	router.GET("/post/:url", gzip.Middleware(blog.Show, false))       // show		Show a specific post
	router.POST("/post/:url", gzip.Middleware(blog.Update, false))    // update Update a specific post
	router.GET("/post/:url/edit/", gzip.Middleware(blog.Edit, false)) // So Form for updating a specific post I maybe should mke a new method to make a more tailored form
	router.GET("/post/:url/images/:imageID", gzip.Middleware(blog.Image, false))

	///////////////////////////////////////////////////////////
	// Examples routes
	///////////////////////////////////////////////////////////

	examples := controllers.Examples{}
	router.GET("/examples", gzip.Middleware(examples.Index, false))

	///////////////////////////////////////////////////////////
	// Static routes
	// Caching Static files
	///////////////////////////////////////////////////////////
	fileServer := http.FileServer(http.Dir("public"))
	router.GET("/static/*filepath", func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Cache-Control", "public, max-age=7776000")
		req.URL.Path = pm.ByName("filepath")
		fileServer.ServeHTTP(res, req)
	})
	return router
}
