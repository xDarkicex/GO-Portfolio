package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/xDarkicex/PortfolioGo/app/controllers/filter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

// Projects controllers
type Projects helpers.Controller

//URL data type for url Shortener
//Hash is the key to the data
//Original is the original url to redirect too
//New is the new shortened url
// type URL struct {
// 	Hash     string
// 	Original string
// 	New      string
// }

// Index ...
func (c Projects) Index(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	var projects []models.Project
	if len(strings.ToLower(a.Request.FormValue("search"))) > 0 {
		projects, err = models.GetProjectsByTags(strings.ToLower(a.Request.FormValue("search")))
		if err != nil {
			helpers.Logger.Printf("Error: %s", err)
		}
	} else {
		projects, err = models.AllProjects()
		if err != nil {
			helpers.Logger.Printf("Error: %s", err)
			return
		}
	}
	proTop, err := models.AllProjects()
	if err != nil {
		helpers.Logger.Printf("Error: %s", err)
	}
	if len(proTop) >= 5 {
		proTop = projects[0:5]
	}
	helpers.Render(a, "projects/index", map[string]interface{}{
		"project": projects,
		"top":     proTop,
		"title":   "Pet Projects",
	})
}

// Create ..
func (c Projects) Create(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
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
	_, err = models.ProjectCreate(a.Request.FormValue("title"), a.Request.FormValue("body"), a.Request.FormValue("summary"), tags, bson.ObjectIdHex(User.(string)), URL, fileBytes, a.Request.FormValue("CustomURL"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/post/"+URL, 302)
}

// New ...
func (c Projects) New(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	helpers.Render(a, "projects/new", map[string]interface{}{
		"project": &models.Project{
			Title:     "",
			Body:      "",
			Summary:   "",
			Tags:      []string{},
			URL:       "",
			CustomURL: "",
		},
	})
}

// Show shows selected project
func (c Projects) Show(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
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

// Edit shows selected blog
func (c Projects) Edit(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	project, err := models.FindProjectByURL(a.Params.ByName("url"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	helpers.Render(a, "projects/edit", map[string]interface{}{
		"project": project,
	})
}

// Image shows selected project
func (c Projects) Image(a helpers.RouterArgs) {
	b, err := models.GetImageByID(a.Params.ByName("imageID"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	a.Response.Write(b)
}

// Update ...
func (c Projects) Update(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	if len(a.Request.FormValue("_method")) > 0 && string(a.Request.FormValue("_method")) == "DELETE" {
		project, err := models.FindProjectByURL(a.Params.ByName("url"))
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		// Actually update
		err = models.DestroyProject(project.ID)
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		http.Redirect(a.Response, a.Request, "/posts", 302)
		return
	}
	project, err := models.FindProjectByURL(a.Params.ByName("url"))
	tags := strings.Split(a.Request.FormValue("tags"), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	hasScript, err := regexp.MatchString("(?:<script.*?>|on(?:click|load|blur|focus|mouse(?:in|out)|hover)\\s*=\\s*['\"]|href\\s*=\\s*['\"]javascript\\:)", a.Request.FormValue("body"))
	if err != nil {
		helpers.Logger.Printf("There is an error in %s", err)
		return
	}
	if hasScript {
		helpers.Logger.Printf("Body form has script tag")
		http.Redirect(a.Response, a.Request, "/project/"+a.Params.ByName("url")+"/edit", 302)
		return
	}
	newProject := make(map[string]interface{})
	for _, key := range []string{"title", "body", "url", "customURL"} {
		value := a.Request.FormValue(key)
		if len(value) > 0 {
			newProject[key] = value
		}
	}

	if tags != nil {
		newProject["tags"] = tags
	}
	// Get file
	file, _, err := a.Request.FormFile("file")
	if err == nil {
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		} else {
			newProject["Image"] = fileBytes
		}
	}
	// Actually update
	err = models.ProjectUpdate(project.ID.Hex(), newProject)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/project/"+string(a.Request.FormValue("url")), 302)
}

//NeuronShow Method that is used to render Neuron Page
func (c Projects) NeuronShow(a helpers.RouterArgs) {
	helpers.Render(a, "projects/neuron", map[string]interface{}{})
}

func (c Projects) ClassLocations(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	fmt.Println(a.Request.FormValue("lat"), a.Request.FormValue("lng"))
	file, _ := os.OpenFile("locations.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	r := strings.NewReplacer(",", " ")
	file.WriteString(a.Request.RemoteAddr + ", " + a.Request.FormValue("lat") + ", " + a.Request.FormValue("lng") + ", " + r.Replace(a.Request.FormValue("address")) + ", " + "\n")
	resp, err := http.PostForm(config.Data.DiscordEndPoint,
		url.Values{"content": {a.Request.RemoteAddr + ", " + a.Request.FormValue("lat") + ", " + a.Request.FormValue("lng") + ", " + r.Replace(a.Request.FormValue("address")) + ", " + "\n"}})
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	a.Response.Header().Set("Origin", "localhost")
	a.Response.Header().Set("Access-Control-Request-Method", "PUT")
	a.Response.Header().Set("Access-Control-Allow-Origin", "*")
	a.Response.Header().Set("Access-Control-Request-Headers", "application/json")
}

func (c Projects) Shorten(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	helpers.Render(a, "projects/urlshortener", map[string]interface{}{})
}
