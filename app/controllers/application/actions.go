package application

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Index is our index action.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helpers.Render(w, "application/index")
}
