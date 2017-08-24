package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"

	"github.com/xDarkicex/PortfolioGo/app/controllers"
	"github.com/xDarkicex/PortfolioGo/app/neuron"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// type routesHandler func(a helpers.Params, user interface{})

// func route(controller helpers.RoutesHandler, authRequired bool) httprouter.Handle {
// 	return gzip.Middleware(auth.Auth(controller, authRequired))
// }

func GetRoutes() *helpers.Server {
	server := helpers.New()
	///////////////////////////////////////////////////////////
	// Static routes
	///////////////////////////////////////////////////////////
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./public/")))
	server.AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^/static/",
		HasRegex: true,
		Handler: func(p *helpers.Params) {
			fileServer.ServeHTTP(p.Response, p.Request)
		},
	})

	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////
	application := controllers.Application{}
	server.AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/",
		Handler: application.Index,
	}).AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/about",
		Handler: application.About,
	}).AddRoute(&helpers.Route{
		Method:  "POST",
		Path:    "/contact",
		Handler: application.Contact,
	})

	///////////////////////////////////////////////////////////
	// users routes
	///////////////////////////////////////////////////////////
	users := controllers.Users{}
	server.AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/users",
		Handler: users.Index,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/users\\/\\w+\\/?$",
		HasRegex: true,
		Handler:  users.Show,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/users\\/\\w+\\/edit\\/?$",
		HasRegex: true,
		Handler:  users.Edit,
	}).AddRoute(&helpers.Route{
		Method:   "POST",
		Path:     "^\\/users\\/\\w+\\/?$",
		HasRegex: true,
		Handler:  users.Update,
	}).AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/register",
		Handler: users.New,
	}).AddRoute(&helpers.Route{
		Method:  "POST",
		Path:    "/register",
		Handler: users.Create,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/users\\/\\w+\\/images\\/\\w+\\/?$",
		HasRegex: true,
		Handler:  users.Image,
	})

	///////////////////////////////////////////////////////////
	// Session Management
	///////////////////////////////////////////////////////////
	sessions := controllers.Sessions{}
	server.AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/signin",
		Handler: sessions.New,
	}).AddRoute(&helpers.Route{
		Method:  "POST",
		Path:    "/signin",
		Handler: sessions.Create,
	}).AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/signout",
		Handler: sessions.Destroy,
	})

	///////////////////////////////////////////////////////////
	// Blog routes
	///////////////////////////////////////////////////////////
	blog := controllers.Blog{}
	server.AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/posts/new",
		Handler: blog.New,
	}).AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/posts",
		Handler: blog.Index,
	}).AddRoute(&helpers.Route{
		Method:  "POST",
		Path:    "/posts",
		Handler: blog.Create,
	}).AddRoute(&helpers.Route{
		Method:  "POST",
		Path:    "/posts/search",
		Handler: blog.Search,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/post\\/\\w+?[_-][\\w-]+\\/?$",
		HasRegex: true,
		Handler:  blog.Show,
	}).AddRoute(&helpers.Route{
		Method:   "POST",
		Path:     "^\\/post\\/\\w+?[_-][\\w-]+\\/?$",
		HasRegex: true,
		Handler:  blog.Update,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/post\\/\\w+?[_-][\\w-]+\\/edit\\/?$",
		HasRegex: true,
		Handler:  blog.Edit,
	}).AddRoute(&helpers.Route{
		Method:   "GET",
		Path:     "^\\/post\\/\\w+?[_-][\\w-]+\\/images\\/\\w+\\/?$",
		HasRegex: true,
		Handler:  blog.Image,
	}).AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/api/posts/search",
		Handler: blog.APIIndex,
	})

	///////////////////////////////////////////////////////////
	// Websocket Connection
	///////////////////////////////////////////////////////////
	w := controllers.Websocket{}
	server.AddRoute(&helpers.Route{
		Method:  "GET",
		Path:    "/api/websocket",
		Handler: w.DialSocket,
	})
	return server
}

// GetRoutes func to setup all routes
// func GetRoutes() *httprouter.Router {
// 	router := httprouter.New()
// 	///////////////////////////////////////////////////////////
// 	// Main application routes
// 	///////////////////////////////////////////////////////////

// 	application := controllers.Application{}
// 	router.GET("/", route(application.Index, false))
// 	router.GET("/about", route(application.About, false))
// 	router.POST("/contact", route(application.Contact, false))

// 	///////////////////////////////////////////////////////////
// 	// users routes
// 	///////////////////////////////////////////////////////////

// 	users := controllers.Users{}
// 	router.GET("/users", route(users.Index, false))
// 	router.GET("/users/:name", route(users.Show, false))
// 	router.GET("/register", route(users.New, false))
// 	router.POST("/users/:name", route(users.Update, false))
// 	router.GET("/users/:name/edit/", route(users.Edit, true))
// 	router.POST("/register", route(users.Create, false))
// 	router.GET("/users/:name/images/:imageID", route(users.Image, false))

// 	///////////////////////////////////////////////////////////
// 	// Session Management
// 	///////////////////////////////////////////////////////////

// 	sessions := controllers.Sessions{}
// 	router.GET("/signin", route(sessions.New, false))
// 	router.POST("/signin", route(sessions.Create, false))
// 	router.GET("/signout", route(sessions.Destroy, false))

// 	///////////////////////////////////////////////////////////
// 	// Blog routes
// 	///////////////////////////////////////////////////////////

// 	blog := controllers.Blog{}
// 	router.GET("/posts", route(blog.Index, false))          // index
// 	router.POST("/posts/search", route(blog.Search, false)) // Route for searching
// 	router.GET("/posts/new", route(blog.New, true))         // new 		To make a new Post
// 	router.POST("/posts", route(blog.Create, true))         // create	To actually throw it in the database
// 	router.GET("/post/:url", route(blog.Show, false))       // show		Show a specific post
// 	router.POST("/post/:url", route(blog.Update, true))     // update Update a specific post
// 	router.GET("/post/:url/edit/", route(blog.Edit, true))  // So Form for updating a specific post I maybe should mke a new method to make a more tailored form
// 	router.GET("/post/:url/images/:imageID", route(blog.Image, false))
// 	router.GET("/api/posts/search", route(blog.APIIndex, false))

// 	///////////////////////////////////////////////////////////
// 	// Examples routes
// 	///////////////////////////////////////////////////////////

// 	projects := controllers.Projects{}
// 	router.GET("/projects", route(projects.Index, false))  // index pages
// 	router.GET("/projects/new", route(projects.New, true)) // render page for creation of new project
// 	router.POST("/projects", route(projects.Create, true)) // post to db
// 	router.GET("/project/:url", route(projects.Show, false))
// 	router.POST("/project/:url", route(projects.Update, true))
// 	router.GET("/project/:url/edit/", route(projects.Edit, true))
// 	router.GET("/project/:url/images/:imageID", route(projects.Image, false))

// 	// custom routes
// 	router.GET("/projects/examples/neuron-demo", route(projects.NeuronShow, false))
// 	router.GET("/api/websocket", DialSocket)

///////////////////////////////////////////////////////////
// Static routes
// Caching Static files
///////////////////////////////////////////////////////////
// 	fileServer := http.FileServer(http.Dir("public"))
// 	router.GET("/static/*filepath", gzip.Middleware(func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
// 		res.Header().Set("Vary", "Accept-Encoding")
// 		res.Header().Set("Cache-Control", "public, max-age=7776000")
// 		req.URL.Path = pm.ByName("filepath")
// 		fileServer.ServeHTTP(res, req)
// 	}))
// 	return router
// }

func DialSocket(w http.ResponseWriter, r *http.Request) {
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
