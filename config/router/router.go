package router

import (
	"encoding/json"

	"net/http"
	"strconv"

	"golang.org/x/net/websocket"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/controllers"
	"github.com/xDarkicex/PortfolioGo/app/neuron"
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
	router.HEAD("/", route(application.Index, false))
	router.GET("/about", route(application.About, false))
	router.HEAD("/about", route(application.About, false))
	router.POST("/contact", route(application.Contact, false))

	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////

	users := controllers.Users{}
	router.GET("/users", route(users.Index, false))
	router.HEAD("/users", route(users.Index, false))
	router.GET("/users/:name", route(users.Show, false))
	router.HEAD("/users/:name", route(users.Show, false))
	router.GET("/register", route(users.New, false))
	router.POST("/users/:name", route(users.Update, false))
	router.GET("/users/:name/edit/", route(users.Edit, true))
	router.HEAD("/users/:name/edit/", route(users.Edit, true))
	router.POST("/register", route(users.Create, false))
	router.GET("/users/:name/images/:imageID", route(users.Image, false))
	router.HEAD("/users/:name/images/:imageID", route(users.Image, false))

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
	router.GET("/api/posts/search", route(blog.APIIndex, false))

	///////////////////////////////////////////////////////////
	// Examples routes
	///////////////////////////////////////////////////////////

	projects := controllers.Projects{}
	router.GET("/projects", route(projects.Index, false))  // index pages
	router.GET("/projects/new", route(projects.New, true)) // render page for creation of new project
	router.POST("/projects", route(projects.Create, true)) // post to db
	router.GET("/project/:url", route(projects.Show, false))
	router.POST("/project/:url", route(projects.Update, true))
	router.GET("/project/:url/edit/", route(projects.Edit, true))
	router.GET("/project/:url/images/:imageID", route(projects.Image, false))

	blacklist := controllers.Blacklist{}
	router.GET("/blacklist", route(blacklist.Index, true))
	router.POST("/blacklist", route(blacklist.Index, true))
	router.GET("/blacklist/:ip", route(blacklist.Remove, true))

	// custom routes
	router.GET("/projects/examples/neuron-demo", route(projects.NeuronShow, false))
	router.GET("/projects/examples/url-shortener", route(projects.Shorten, false))
	router.POST("/api/classLocations", route(projects.ClassLocations, false))
	router.GET("/api/websocket", DialSocket)

	///////////////////////////////////////////////////////////
	// Static routes
	// Caching Static files
	///////////////////////////////////////////////////////////

	router.GET("/static/*filepath", func(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		r.URL.Path = pm.ByName("filepath")
		http.FileServer(http.Dir("public")).ServeHTTP(w, r)

	})

	return router
}

func DialSocket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	websocket.Handler(Socket).ServeHTTP(w, r)
}

func Socket(ws *websocket.Conn) {
	var msg string
	for {
		if websocket.Message.Receive(ws, &msg) != nil {
			break
		}
		var data = make(map[string]interface{})
		json.Unmarshal([]byte(msg), &data)
		switch data["api"] {
		case "neuron":
			point := data["data"].(map[string]interface{})
			output := neuron.Ne.Process([]float64{point["x"].(float64), point["y"].(float64)})
			websocket.Message.Send(ws, `{
							"output": "`+strconv.FormatFloat(output, 'f', -1, 64)+`",
							"M": `+strconv.FormatFloat(neuron.M, 'f', -1, 64)+`, 
							"B": `+strconv.FormatFloat(neuron.B, 'f', -1, 64)+`}`)
		default:
			break
		}
	}
}
