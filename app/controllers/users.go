package controllers

import (
	"fmt"
	"log"
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
	helpers.RenderDynamic(req, res, "users/index", map[string]interface{}{
		"users": users,
	})
}

// UserCreate a new user
func UserCreate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	success, createErr := models.CreateUser(req.FormValue("email"), req.FormValue("name"), req.FormValue("password"))
	// fmt.Fprintln(res, message)
	if success {
		user, err := models.Login(req.FormValue("name"), req.FormValue("password"))
		if err != nil {
			fmt.Println(err)
		} else {
			log.Print(user)
		}
		http.Redirect(res, req, "/", 302)
	} else {
		session, err := helpers.Store().Get(req, "flash-session")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		session.AddFlash(createErr, "message")
		err = session.Save(req, res)
		if err != nil {
			helpers.Logger.Println(err)
		}
		fmt.Println(" I got here", createErr)
		fm := session.Flashes("message")
		fmt.Println(fm)
		if fm == nil {
			fmt.Fprint(res, "No flash messages")
			return
		}
		session.Save(req, res)
		http.Redirect(res, req, "/register", http.StatusFound)
	}
}

//UserShow Show page for users
func UserShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	user, err := models.FindUserByName(params.ByName("name"))
	if err != nil {
		http.Redirect(res, req, "/404", 302)
	}
	helpers.RenderDynamic(req, res, "users/show", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"user":   user,
	})
}
