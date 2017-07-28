package helpers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"log"

	"github.com/gorilla/sessions"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/redirect"
)

type Params struct {
	Response http.ResponseWriter
	Request  *http.Request
	Session  *sessions.Session
	User     interface{}
}

// Route type for adding to the server
type Route struct {
	Regexp   *regexp.Regexp
	HasRegex bool
	Path     string
	Method   string
	Handler  func(*Params)
}

// Server struct for fancy methods
type Server struct {
	Mux    *http.ServeMux
	Routes []*Route
}

func (s *Server) AddRoute(route *Route) *Server {
	if route.HasRegex {
		var err error
		route.Regexp, err = regexp.Compile(route.Path)
		if err != nil {
			fmt.Println(errors.WithStack(err))
		}
	}
	s.Routes = append(s.Routes, route)
	return s
}

func (s *Server) Serve(port string) {
	lime := ansi.ColorFunc("green+h:black")
	red := ansi.ColorFunc("red+h:black")

	s.Mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		session, err := Store().Get(request, "user-session")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		for _, route := range s.Routes {
			if (route.HasRegex && route.Regexp.MatchString(request.URL.Path) || request.URL.Path == route.Path) && route.Method == request.Method {
				route.Handler(&Params{
					Response: response,
					Request:  request,
					Session:  session,
				})
				log.Println("["+request.Method+"]", lime("200"), ":", request.Host, request.URL.String())
				return
			}
		}
		log.Println("["+request.Method+"]", red("404"), ":", request.Host, request.URL.String())
		http.NotFound(response, request)
	})
	listen := fmt.Sprintf("%s:%d", config.Data.Host, config.Data.Port)
	if config.Data.SSL {
		go func() {
			Logger.Fatal(http.ListenAndServe(config.Data.Host+":80", http.HandlerFunc(redirect.HTTPS)))
		}()
		Logger.Fatal(http.ListenAndServeTLS(listen, config.Data.Cert, config.Data.Key, s.Mux))
	} else {
		Logger.Fatal(http.ListenAndServe(listen, s.Mux))
	}
}

// New will load all server and routes
func New() *Server {
	return &Server{
		Mux: http.NewServeMux(),
	}
}
