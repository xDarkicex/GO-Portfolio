package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Sessions Controller!
type Sessions helpers.Controller

// New ...
func (c Sessions) New(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	helpers.Render(a, "sessions/new", map[string]interface{}{})
}

// Create ..
func (c Sessions) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	user, err := models.Login(a.Request.FormValue("name"), a.Request.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		helpers.AddFlash(a, helpers.Flash{Type: "danger", Message: err.Error()})
		err = session.Save(a.Request, a.Response)
		http.Redirect(a.Response, a.Request, "/signin", 302)
	} else {
		helpers.Logger.Println("Logging in as " + user.Name)
		session.Values["UserID"] = user.ID.Hex()
		helpers.AddFlash(a, helpers.Flash{Type: "success", Message: "Sucessully logged in!"})
		err := session.Save(a.Request, a.Response)
		if err != nil {
			helpers.Logger.Println(err)
		}
		http.Redirect(a.Response, a.Request, "/users/"+user.Name, 302)
	}
}

// Destroy ...
func (c Sessions) Destroy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	session.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}

	session.Save(a.Request, a.Response)
	http.Redirect(a.Response, a.Request, "/", 302)
}
