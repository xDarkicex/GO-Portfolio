package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// ApplicationIndex is our index action.
func ApplicationIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	view := "application/index"
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
	})
}

// ApplicationExamples example pages
func ApplicationExamples(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	// user, err := models.FindUserByName(params.ByName("name"))

	if err != nil {
		fmt.Println("There was an error")
	}
	view := "application/examples"
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		// "user":   user,
	})
}

//ApplicationAbout About me Pages
func ApplicationAbout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	user, err := models.FindUserByName("xDarkicex")
	if err != nil {
		fmt.Println("There was an error")
	}
	view := "application/about"
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"user":   user,
	})
}

// Error404 is our index action.
func Error404(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/404")
}

// RegisterNew Users
func RegisterNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	view := "users/new"
	if session.Values["UserID"] == nil {
		helpers.RenderDynamic(req, res, view, map[string]interface{}{
			"UserID": session.Values["UserID"],
		})
	} else {
		http.Redirect(res, req, "/", 302)
	}
}
