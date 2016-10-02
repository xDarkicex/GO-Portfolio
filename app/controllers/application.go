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

// ApplicationExamples is the project examples I guess
func ApplicationExamples(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/examples")
}

//ApplicationAbout I guess this is the about me
func ApplicationAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/about")
}
