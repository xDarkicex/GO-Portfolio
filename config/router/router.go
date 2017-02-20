package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
	"github.com/xDarkicex/PortfolioGo/config/auth"
	"github.com/xDarkicex/PortfolioGo/config/gzip"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// type routesHandler func(a RouterArgs, user interface{})
func route(controller helpers.RoutesHandler, authRequired bool) httprouter.Handle {
	return gzip.Middleware(auth.Auth(controller, authRequired))
}

// GetRoutes func to setup all routes
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////

	application := controllers.Application{}
	router.GET("/", route(application.Index, false))
	router.GET("/about", route(application.About, false))
	router.POST("/contact", route(application.Contact, false))

	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////

	users := controllers.Users{}
	router.GET("/users", route(users.Index, false))
	router.GET("/users/:name", route(users.Show, false))
	router.GET("/register", route(users.New, false))
	router.POST("/users/:name", route(users.Update, false))
	router.GET("/users/:name/edit/", route(users.Edit, true))
	router.POST("/register", route(users.Create, false))
	router.GET("/users/:name/images/:imageID", route(users.Image, false))

	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////

	sessions := controllers.Sessions{}
	router.GET("/signin", route(sessions.New, false))
	router.POST("/signin", route(sessions.Create, false))
	router.GET("/signout", route(sessions.Destroy, false))

	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////

	blog := controllers.Blog{}
	router.GET("/posts", route(blog.Index, false))          // index
	router.POST("/posts/search", route(blog.Search, false)) // Route for searching
	router.GET("/posts/new", route(blog.New, true))         // new 		To make a new Post
	router.POST("/posts", route(blog.Create, true))         // create	To actually throw it in the database
	router.GET("/post/:url", route(blog.Show, false))       // show		Show a specific post
	router.POST("/post/:url", route(blog.Update, true))     // update Update a specific post
	router.GET("/post/:url/edit/", route(blog.Edit, true))  // So Form for updating a specific post I maybe should mke a new method to make a more tailored form
	router.GET("/post/:url/images/:imageID", route(blog.Image, false))

	///////////////////////////////////////////////////////////
	// Examples routes
	///////////////////////////////////////////////////////////

	projects := controllers.Projects{}
	router.GET("/projects", route(projects.Index, false))

	///////////////////////////////////////////////////////////
	// Static routes
	// Caching Static files
	///////////////////////////////////////////////////////////
	fileServer := http.FileServer(http.Dir("public"))
	router.GET("/static/*filepath", gzip.Middleware(func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Cache-Control", "public, max-age=7776000")
		req.URL.Path = pm.ByName("filepath")
		fileServer.ServeHTTP(res, req)
	}))
	return router
}
