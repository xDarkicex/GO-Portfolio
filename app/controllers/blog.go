package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// BlogIndex for indexing all blog posts
func BlogIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// fmt.Printf("Users %s", users)
	helpers.RenderDynamic(res, "blog/index", map[string]interface{}{
		"title": blogs,
	})
}

// BlogNew for new post
func BlogNew(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	helpers.RenderDynamic(res, "blog/new", map[string]interface{}{})
}

// BlogEdit for edit blog Post
func BlogEdit(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	blog, err := models.FindBlogByTitle(params.ByName("title"))
	if err != nil {
		http.Redirect(res, req, "/404", 404)
	} else {
		helpers.RenderDynamic(res, "blog/edit", map[string]interface{}{
			"blog": blog,
		})
	}
}

// BlogShow shows selected blog
func BlogShow(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	blog, err := models.FindBlogByTitle(params.ByName("title"))
	if err != nil {
		http.Redirect(res, req, "/404", 404)
	} else {
		helpers.RenderDynamic(res, "blog/show", map[string]interface{}{
			"blog": blog,
		})
	}
}
