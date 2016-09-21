package users

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models/users"
)

// Create a new user
func Create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// ctx := appengine.NewContext(req)
	users.Create(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	fmt.Fprintln(res, "Created user?")
	// createSession(res, req, user)
	// redirect
	// http.Redirect(res, req, "/", 302)
}
