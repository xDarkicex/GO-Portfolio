package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
)

// UserCreate a new user
func UserCreate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("Username")
	fmt.Println(req.FormValue("name"))
	fmt.Println("Email")
	fmt.Println(req.FormValue("email"))
	fmt.Println("Password")
	fmt.Println(req.FormValue("password"))
	success, _ := models.CreateUser(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	// fmt.Fprintln(res, message)
	if success {
		http.Redirect(res, req, "/", 302)
	}
	// redirect
	// http.Redirect(res, req, "/", 302)
}

func UserAuth(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user, err := models.GetUser(req.FormValue("name"), req.FormValue("password"))
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
}
