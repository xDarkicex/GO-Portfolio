package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// SignupIndex is our index action.
func SignupIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "signup/index")

}

// Create a new user
func Create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// ctx := appengine.NewContext(req)#
	fmt.Println("Username")
	fmt.Println(req.FormValue("name"))
	fmt.Println("Email")
	fmt.Println(req.FormValue("email"))
	fmt.Println("Password")
	fmt.Println(req.FormValue("password"))
	worked := models.CreateUser(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	if worked {
		fmt.Fprintln(res, "User Created.")
	} else {
		fmt.Fprintln(res, "Error creating user.")
	} // createSession(res, req, user)
	// redirect
	// http.Redirect(res, req, "/", 302)
}
