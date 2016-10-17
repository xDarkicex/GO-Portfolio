package helpers

import (
	"net/http"

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
}
