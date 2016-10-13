package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// SessionNew GET
func SessionNew(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	helpers.Render(res, "sessions/new")
}

// SessionCreate POST validate and login
func SessionCreate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	user, err := models.Login(req.FormValue("name"), req.FormValue("password"))
	if err != nil {
		http.Redirect(res, req, "/signin", 302)
	} else {
		session.Values["UserID"] = user.ID.Hex()
		err := session.Save(req, res)
		if err != nil {
			helpers.Logger.Println(err)
		}
		http.Redirect(res, req, "/users/"+user.Name, 302)
	}
}

// SessionDestroy GET destroy our session
func SessionDestroy(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	session.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}

	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
}

// SessionsSignIn creates new cookie on signin
// func SessionsSignIn(user models.User, w http.ResponseWriter) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "id",
// 		Value:    hex.EncodeToString([]byte(user.ID)),
// 		MaxAge:   0,
// 		Secure:   false,
// 		HttpOnly: true,
// 	})
// }

// SessionsSignOut Sign out of whatever session existed
// func SessionsSignOut(w http.ResponseWriter) {
// 	// http.SetCookie(w, &http.Cookie{
// 	// 	Name:     "id",
// 	// 	Value:    "",
// 	// 	MaxAge:   -1,
// 	// 	Secure:   false,
// 	// 	HttpOnly: true,
// 	// })

// }
