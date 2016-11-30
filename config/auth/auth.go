package auth

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Auth does some stuff.
func Auth(fn helpers.RoutesHandler, requireAuth bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
		session, err := helpers.Store().Get(a.Request, "user-session")
		if err != nil {
			http.Error(a.Response, err.Error(), http.StatusInternalServerError)
			return
		}
		a.Session = session
		userID := session.Values["UserID"]
		if userID != nil {
			// Not logged in
			// helpers.Logger.Println(userID)
			user, err := models.FindUserByID(bson.ObjectIdHex(userID.(string)))
			if err != nil {
				helpers.Logger.Println(err)
			} else {
				a.User = user
				fn(a) // Rewrite your controllers to be fancy.
			}
		} else if requireAuth {
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		} else {
			// Auth not required, userID not found.
			fn(a)
		}
	}
}

// // Middleware force - bool, whether or not to force Gzip regardless of the sent headers.
// func Middleware(fn httprouter.Handle) httprouter.Handle {
// 	return func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
// 		res.Header().Set("Server", "Golang")
// 		if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip")  {
// 			fn(res, req, pm)
// 			return
// 		}
// 		res.Header().Set("Vary", "Accept-Encoding")
// 		res.Header().Set("Content-Encoding", "gzip")
// 		gz := gzip.NewWriter(res)
// 		defer gz.Close()
// 		gzr := GzipResponseWriter{Writer: gz, ResponseWriter: res}
// 		fn(gzr, req, pm)
// 	}
// }
