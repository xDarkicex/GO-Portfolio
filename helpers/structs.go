package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// Controller Struct
type Controller struct {
	Controller interface{}
	Globals    struct {
		Count int
	}
}

// RouterArgs These are the arguments passed in from a router.
type RouterArgs struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	Session  *sessions.Session
	User     interface{}
}

// Flash ...
type Flash struct {
	Type    string
	Message string
}

// RoutesHandler for handling padding multiple objects into routes
type RoutesHandler func(a RouterArgs)
