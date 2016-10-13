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
	session, err := helpers.Store().Get(req, "user-session")
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
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
	})
}

// BlogNew render blog post page
func BlogNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	helpers.RenderDynamic(req, res, "blog/new", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog": &models.Blog{
			Title: "",
			Body:  "",
			Tags:  []string{},
			URL:   "",
		},
	})
}

// BlogCreate for new post
func BlogCreate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
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
	http.Redirect(res, req, "/post/"+string(req.FormValue("url")), 302)
}

// BlogUpdate for edit blog Post
func BlogUpdate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_, err := helpers.Store().Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	if len(req.FormValue("_method")) > 0 && string(req.FormValue("_method")) == "DELETE" {
		blog, err := models.FindBlogByURL(params.ByName("url"))
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(res, req, "/", 302)
			return
		}
		// Actually update
		err = models.BlogDestroy(blog.ID)
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(res, req, "/", 302)
			return
		}
		http.Redirect(res, req, "/posts", 302)
		return
	}
	blog, err := models.FindBlogByURL(params.ByName("url"))
	tags := strings.Split(req.FormValue("tags"), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	newPost := map[string]interface{}{}
	for _, key := range []string{"title", "body", "url"} {
		value := req.FormValue(key)
		if len(value) > 0 {
			newPost[key] = value
		}
	}
	if tags != nil {
		newPost["tags"] = tags
	}
	// Get file
	file, _, err := req.FormFile("file")
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
		http.Redirect(res, req, "/", 302)
		return
	}
	http.Redirect(res, req, "/post/"+string(req.FormValue("url")), 302)
}

// BlogShow shows selected blog
func BlogShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
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

	author, err := models.FindUserByID(blog.UserID)
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	view := "blog/show"
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blog,
		"author": author,
	})
}

// BlogEdit shows selected blog
func BlogEdit(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := helpers.Store().Get(req, "user-session")
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
	view := "blog/edit"
	helpers.RenderDynamic(req, res, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blog,
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
