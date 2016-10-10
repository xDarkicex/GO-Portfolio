package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// UserIndex for indexing all users
func UserIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	users, err := models.AllUsers()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// fmt.Printf("Users %s", users)
	helpers.RenderDynamic(res, "users/index", map[string]interface{}{
		"users": users,
	})
}

// UserNew a new user
func UserNew(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	/////////////////////////////////////////////
	// This is for testing.
	/////////////////////////////////////////////
	// fmt.Println(req.FormValue("name"))
	// fmt.Println(req.FormValue("email"))
	// fmt.Println(req.FormValue("password"))

	success, _ := models.CreateUser(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	// fmt.Fprintln(res, message)
	if success {
		user, err := models.Login(req.FormValue("name"), req.FormValue("password"))
		if err != nil {
			fmt.Println(err)
		} else {
			// SessionsSignIn(user, res)
			fmt.Println("I got here", user)
		}
		http.Redirect(res, req, "/", 302)
	}
}

//UserShow Show page for users
func UserShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	user, err := models.FindUserByName(params.ByName("name"))
	helpers.Logger.Println(user)
	if err != nil {
		http.Redirect(res, req, "/404", 302)
		// fmt.Println("/////////////////////////")
		// defer fmt.Println("/////////////////////////")
		// fmt.Println("Error /404")
		// fmt.Println(err)
		// http.Redirect(res, req, "/404", 404)

	} else {
		helpers.RenderDynamic(res, "users/show", map[string]interface{}{
			"UserID": session.Values["UserID"],
			"user":   user,
		})
	}
}
