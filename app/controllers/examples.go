package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Examples controllers
type Examples helpers.Controller

// Index ...
func (c Examples) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	examples, err := models.AllExamples()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(examples) >= 5 {
		examples = examples[0:5]
	}
	view := "examples/index"
	helpers.Render(a, view, map[string]interface{}{
		"UserID":  session.Values["UserID"],
		"example": examples,
	})
}
