package auth

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Auth wraps Route in authenticates for userID
func Auth(next helpers.RoutesHandler, requireAuth bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
		session, err := helpers.Store().Get(a.Request, "user-session")
		if err != nil {
			http.Error(a.Response, err.Error(), http.StatusInternalServerError)
			return
		}
		a.Session = session

		userID, ok := session.Values["UserID"].(string)
		if ok {
			user, err := models.FindUserByID(bson.ObjectIdHex(userID))
			if err != nil {
				helpers.Logger.Println(err)
			} else {
				a.User = user
				next(a)
			}
		} else if requireAuth {
			http.Redirect(a.Response, a.Request, "/", http.StatusFound)
			return
		} else {
			next(a)
		}
	}
}
