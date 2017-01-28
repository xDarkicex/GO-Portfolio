package controllers

import (
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Projects controllers
type Projects helpers.Controller

// Index ...
func (c Projects) Index(a helpers.RouterArgs) {
	projects, err := models.AllProjects()
	if err != nil {
		helpers.Logger.Printf("Error: %s", err)
		return
	}
	if len(projects) >= 5 {
		projects = projects[0:5]
	}
	helpers.Render(a, "projects/index", map[string]interface{}{
		"project": projects,
		"title":   "Pet Projects",
	})
}
