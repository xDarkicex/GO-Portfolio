package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
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

// Create ..
func (c Projects) Create(a helpers.RouterArgs) {
	session := a.Session
	User := session.Values["UserID"]

	// File processing ...
	file, _, _ := a.Request.FormFile("file")
	fileBytes, _ := ioutil.ReadAll(file)
	tags := strings.Split(strings.ToLower(a.Request.FormValue("tags")), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	// URL Processing
	rawURL := a.Request.FormValue("title")
	URL := strings.Replace(rawURL, " ", "-", -1)
	_, err := models.ProjectCreate(a.Request.FormValue("title"), a.Request.FormValue("body"), a.Request.FormValue("summary"), tags, bson.ObjectIdHex(User.(string)), URL, fileBytes)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/post/"+URL, 302)
}

// New ...
func (c Projects) New(a helpers.RouterArgs) {
	helpers.Render(a, "projects/new", map[string]interface{}{
		"project": &models.Project{
			Title:   "",
			Body:    "",
			Summary: "",
			Tags:    []string{},
			URL:     "",
		},
	})
}

// Show shows selected blog
func (c Projects) Show(a helpers.RouterArgs) {
	project, err := models.FindProjectByURL(a.Params.ByName("url"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	projects, err := models.AllProjects()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(projects) >= 5 {
		projects = projects[0:5]
	}
	helpers.Render(a, "projects/show", map[string]interface{}{
		"project":  project,
		"projects": projects,
	})
}
