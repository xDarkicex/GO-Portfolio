package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// SignupIndex is our index action.
func SignupIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "signup/index")
}
