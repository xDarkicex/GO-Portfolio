package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// LoginIndex is our index action.
func LoginIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "login/index")
}
