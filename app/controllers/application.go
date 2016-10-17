package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Application Controller.
type Application helpers.Controller

//Index New index function
func (c Application) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(blogs) >= 5 {
		blogs = blogs[0:5]
	}
	view := "application/index"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
	})
}

//About About me Pages
func (c Application) About(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	user, err := models.FindUserByName("xDarkicex")
	if err != nil {
		fmt.Println("There was an error")
	}
	view := "application/about"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"user":   user,
	})
}

// RegisterNew Users
// func RegisterNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	session, err := helpers.Store().Get(req, "user-session")
// 	if err != nil {
// 		helpers.Logger.Println(err)
// 		http.Redirect(res, req, "/", 302)
// 		return
// 	}
// 	view := "users/new"
// 	if session.Values["UserID"] == nil {
// 		helpers.RenderDynamic(req, res, view, map[string]interface{}{
// 			"UserID": session.Values["UserID"],
// 		})
// 	} else {
// 		http.Redirect(res, req, "/", 302)
// 	}
// }
