package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
)

// UserCreate a new user
func UserCreate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// ctx := appengine.NewContext(req)#
	fmt.Println("Email")
	fmt.Println(req.FormValue("email"))
	worked := models.CreateUser(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	if worked {
		fmt.Fprintln(res, "User Created.")
	} else {
		fmt.Fprintln(res, "Error creating user.")
	} // createSession(res, req, user)
	// redirect
	// http.Redirect(res, req, "/", 302)
}
