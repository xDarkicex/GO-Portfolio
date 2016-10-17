package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Users Controller!
type Users helpers.Controller

// Index function
func (c Users) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	users, err := models.AllUsers()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// fmt.Printf("Users %s", users)
	helpers.Render(a, "users/index", map[string]interface{}{
		"users": users,
	})
}

// Create a new user
func (c Users) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	success, createErr := models.CreateUser(a.Request.FormValue("email"), a.Request.FormValue("name"), a.Request.FormValue("password"))
	// fmt.Fprintln(a.Response, message)
	if success {
		user, err := models.Login(a.Request.FormValue("name"), a.Request.FormValue("password"))
		if err != nil {
			fmt.Println(err)
		} else {
			log.Print(user)
		}
		http.Redirect(a.Response, a.Request, "/", 302)
	} else {
		session, err := helpers.Store().Get(a.Request, "flash-session")
		if err != nil {
			http.Error(a.Response, err.Error(), http.StatusInternalServerError)
			return
		}
		session.AddFlash(createErr, "message")
		err = session.Save(a.Request, a.Response)
		if err != nil {
			helpers.Logger.Println(err)
		}
		fmt.Println(" I got here", createErr)
		fm := session.Flashes("message")
		fmt.Println(fm)
		if fm == nil {
			fmt.Fprint(a.Response, "No flash messages")
			return
		}
		session.Save(a.Request, a.Response)
		http.Redirect(a.Response, a.Request, "/register", http.StatusFound)
	}
}

// Show Show page for users
func (c Users) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	user, err := models.FindUserByName(a.Params.ByName("name"))
	if err != nil {
		http.Redirect(a.Response, a.Request, "/404", 302)
	}
	helpers.Render(a, "users/show", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"user":   user,
	})
}

// New ...
func (c Users) New(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	view := "users/new"
	if session.Values["UserID"] == nil {
		helpers.Render(a, view, map[string]interface{}{
			"UserID": session.Values["UserID"],
		})
	} else {
		http.Redirect(a.Response, a.Request, "/", 302)
	}
}
