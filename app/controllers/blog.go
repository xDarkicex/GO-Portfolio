package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// BlogIndex for indexing all blog posts
func BlogIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	view := "blog/index"
	helpers.RenderDynamic(res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
		"view":   view,
	})
}

// BlogPostNew render blog post page
func BlogPostNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	fmt.Println(session.Values["IsAdmin"])
	if session.Values["IsAdmin"] == false || session.Values["IsAdmin"] == nil {
		// Need flash message here!!!
		http.Redirect(res, req, "/", 302)
		return
	}
	view := "blog/new"
	helpers.RenderDynamic(res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"view":   view,
	})
}

// BlogNew for new post
func BlogNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)

		return
	}
	User := session.Values["UserID"]
	// File processing ...
	// Note to self, this needs to be made optional...
	file, _, _ := req.FormFile("file")
	fileBytes, _ := ioutil.ReadAll(file)
	tags := strings.Split(req.FormValue("tags"), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	_, err = models.BlogCreate(req.FormValue("title"), req.FormValue("body"), tags, User.(string), req.FormValue("url"), fileBytes)
	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	}
	fmt.Println(tags)
	view := "blog/new"
	helpers.RenderDynamic(res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"view":   view,
	})
}

// BlogEdit for edit blog Post
func BlogEdit(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blog, err := models.FindBlogByURL(params.ByName("url"))
	if err != nil {
		http.Redirect(res, req, "/404", 404)
		return
	}
	view := "blog/edit"
	helpers.RenderDynamic(res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blog,
		"view":   view,
	})
}

// BlogShow shows selected blog
func BlogShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blog, err := models.FindBlogByURL(params.ByName("url"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	user, err := models.FindUserByID(blog.UserID)
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	view := "blog/show"
	helpers.RenderDynamic(res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blog,
		"user":   user,
		"view":   view,
	})
}

// BlogImage shows selected blog
func BlogImage(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	b, err := models.GetImageByID(params.ByName("imageID"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}

	log.Println(b)

	res.Write(b)
}

// BlogSearch ...
func BlogSearch(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	blogs, err := models.GetBlogsByTags(params.ByName("searchTerm"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}

	log.Println(blogs)

	res.Write([]byte("ok"))
}
