package controllers

import (
	"fmt"
	"net/http"

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
	}
	helpers.RenderDynamic(res, "blogs/index", map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
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
	helpers.RenderDynamic(res, "blog/new", map[string]interface{}{
		"UserID": session.Values["UserID"],
	})
}

// BlogNew for new post
func BlogNew(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	success, _ := models.BlogCreate(req.FormValue("title"), req.FormValue("body"))
	if success != true {
		http.Redirect(res, req, "/", 302)
	}
	// blog, err := models.FindBlogByTitle(params.ByName("title"))
	// if err != nil {
	// 	fmt.Println("There was an error")
	// }
	helpers.RenderDynamic(res, "blog/new", map[string]interface{}{
		"UserID": session.Values["UserID"],
		// 	"blog":   blog,
	})
}

// helpers.RenderDynamic(res, "blog/new", map[string]interface{}{})

// BlogEdit for edit blog Post
func BlogEdit(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blog, err := models.FindBlogByTitle(params.ByName("title"))
	if err != nil {
		http.Redirect(res, req, "/404", 404)
	} else {
		helpers.RenderDynamic(res, "blog/edit", map[string]interface{}{
			"UserID": session.Values["UserID"],
			"blog":   blog,
		})
	}
}

// BlogShow shows selected blog
func BlogShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	session, err := Store.Get(req, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(res, req, "/", 302)
		return
	}
	blog, err := models.FindBlogByTitle(params.ByName("title"))
	if err != nil {
		defer fmt.Println("/////////////////////////////")
		fmt.Println("/////////////////////////////")
		helpers.Logger.Println(blog)
		http.Redirect(res, req, "/404", 404)
	} else {
		helpers.RenderDynamic(res, "blogs/show", map[string]interface{}{
			"UserID": session.Values["UserID"],
			"blog":   blog,
		})
	}
}
