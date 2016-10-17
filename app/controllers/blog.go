package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Blog Controller
type Blog helpers.Controller

//Index New index function
func (c Blog) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c.Globals.Count++
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	var blogs []models.Blog
	if len(a.Request.FormValue("search")) > 0 {
		blogs, err = models.GetBlogsByTags(a.Request.FormValue("search"))
	} else {
		blogs, err = models.AllBlogs()
	}
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	users, _ := models.AllUsers()
	// blog, err := models.FindBlogByURL(params.ByName("url"))
	// author, err := models.FindUserByID(blog.UserID)
	view := "blog/index"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
		"users":  users,
		"count":  c.Globals.Count,
		// "author": author,
	})
}

//New ....
func (c Blog) New(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	helpers.Render(a, "blog/new", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog": &models.Blog{
			Title: "",
			Body:  "",
			Tags:  []string{},
			URL:   "",
		},
	})
}

// Create ...
func (c Blog) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	User := session.Values["UserID"]

	// File processing ...
	// Note to self, this needs to be made optional...
	file, _, _ := a.Request.FormFile("file")
	fileBytes, _ := ioutil.ReadAll(file)
	tags := strings.Split(a.Request.FormValue("tags"), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	_, err = models.BlogCreate(a.Request.FormValue("title"), a.Request.FormValue("body"), tags, bson.ObjectIdHex(User.(string)), a.Request.FormValue("url"), fileBytes)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/post/"+string(a.Request.FormValue("url")), 302)
}

// Update ...
func (c Blog) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	_, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	if len(a.Request.FormValue("_method")) > 0 && string(a.Request.FormValue("_method")) == "DELETE" {
		blog, err := models.FindBlogByURL(a.Params.ByName("url"))
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		// Actually update
		err = models.BlogDestroy(blog.ID)
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		http.Redirect(a.Response, a.Request, "/posts", 302)
		return
	}
	blog, err := models.FindBlogByURL(a.Params.ByName("url"))
	tags := strings.Split(a.Request.FormValue("tags"), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	newPost := map[string]interface{}{}
	for _, key := range []string{"title", "body", "url"} {
		value := a.Request.FormValue(key)
		if len(value) > 0 {
			newPost[key] = value
		}
	}
	if tags != nil {
		newPost["tags"] = tags
	}
	// Get file
	file, _, err := a.Request.FormFile("file")
	if err == nil {
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		} else {
			newPost["blogImage"] = fileBytes
		}
	}
	// Actually update
	err = models.BlogUpdate(blog.ID.Hex(), newPost)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/post/"+string(a.Request.FormValue("url")), 302)
}

// Show shows selected blog
func (c Blog) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	blog, err := models.FindBlogByURL(a.Params.ByName("url"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return

	}
	view := "blog/show"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"post":   blog,
	})
}

// Edit shows selected blog
func (c Blog) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	blog, err := models.FindBlogByURL(a.Params.ByName("url"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return

	}
	helpers.Render(a, "blog/edit", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blog,
	})
}

// Image shows selected blog
func (c Blog) Image(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	b, err := models.GetImageByID(a.Params.ByName("imageID"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	a.Response.Write(b)
}
