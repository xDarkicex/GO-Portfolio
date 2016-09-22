package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// AboutIndex is our index action.
func AboutIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "about/index")
}
