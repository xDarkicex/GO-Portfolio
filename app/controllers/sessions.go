package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Sessions Controller!
type Sessions helpers.Controller

// New ...
func (c Sessions) New(a *helpers.Params) {
	helpers.Render(a, "sessions/new", map[string]interface{}{})
}

// Create ..
func (c Sessions) Create(a *helpers.Params) {
	session := a.Session
	user, err := models.Login(a.Request.FormValue("name"), a.Request.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		helpers.AddFlash(a, helpers.Flash{Type: "danger", Message: err.Error()})
		err = session.Save(a.Request, a.Response)
		http.Redirect(a.Response, a.Request, "/signin", 302)
	} else {
		// helpers.Logger.Println("Logging in as " + user.Name)
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
func (c Sessions) Destroy(a *helpers.Params) {
	session := a.Session
	session.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}

	session.Save(a.Request, a.Response)
	http.Redirect(a.Response, a.Request, "/", 302)
}
