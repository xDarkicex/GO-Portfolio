package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// ApplicationIndex is our index action.
func ApplicationIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/index")
}

// ApplicationExamples example pages
func ApplicationExamples(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/examples")
}

//ApplicationAbout About me Pages
func ApplicationAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/about")
}
