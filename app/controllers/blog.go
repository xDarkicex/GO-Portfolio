package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Blog Controller
type Blog helpers.Controller

// // Search ...
func (c Blog) Search(a *helpers.Params) {
	searchValue := strings.ToLower(a.Request.FormValue("search"))
	blogs, err := models.GetBlogsByTags(searchValue)
	if err != nil {
		helpers.Logger.Printf("Error: %s", err)
		return
	}

	helpers.Render(a, "blog/index", map[string]interface{}{
		"blog":  blogs,
		"title": searchValue,
	})

}

// //Index New index function
func (c Blog) Index(a *helpers.Params) {
	var blogs []models.Blog
	var err error
	if len(strings.ToLower(a.Request.FormValue("search"))) > 0 {
		blogs, err = models.GetBlogsByTags(strings.ToLower(a.Request.FormValue("search")))
		if err != nil {
			helpers.Logger.Printf("Error: %s", err)
		}
	} else {
		blogs, err = models.AllBlogs()
		if err != nil {
			helpers.Logger.Printf("Error: %s", err)
			return
		}
	}
	blogsTop, err := models.AllBlogs()
	if err != nil {
		helpers.Logger.Printf("Error: %s", err)
		return
	}
	if len(blogsTop) >= 5 {
		blogsTop = blogsTop[0:5]
	}
	if err != nil {
		helpers.Logger.Printf("Error: %s", err)
		return
	}
	// users, err := models.AllUsers()
	// if err != nil {
	// 	helpers.Logger.Printf("Error: %s", err)
	// }
	helpers.Render(a, "blog/index", map[string]interface{}{
		"blog":  blogs,
		"top":   blogsTop,
		"title": "Blog",
	})
}

// //New ....
func (c Blog) New(a *helpers.Params) {
	helpers.Render(a, "blog/new", map[string]interface{}{
		"blog": &models.Blog{
			Title:   "",
			Body:    "",
			Summary: "",
			Tags:    []string{},
			URL:     "",
		},
	})
}

// // Create ...
func (c Blog) Create(a *helpers.Params) {
	session := a.Session
	User := session.Values["UserID"]

	// File processing ...
	file, _, err := a.Request.FormFile("file")
	if err != nil {
		helpers.Logger.Printf("Error: %s\n", err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		helpers.Logger.Printf("Error: %s\n", err)
	}
	tags := strings.Split(strings.ToLower(a.Request.FormValue("tags")), ",")
	for k, v := range tags {
		tags[k] = strings.TrimSpace(v)
	}
	// URL Processing
	rawURL := a.Request.FormValue("title")
	URL := strings.Replace(rawURL, " ", "-", -1)
	err = models.BlogCreate(a.Request.FormValue("title"), a.Request.FormValue("body"), a.Request.FormValue("summary"), tags, bson.ObjectIdHex(User.(string)), URL, fileBytes)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/post/"+URL, 302)
}

// Update ...
func (c Blog) Update(a *helpers.Params) {
	url := strings.Split(a.Request.URL.Path, "/")[2]
	if len(a.Request.FormValue("_method")) > 0 && string(a.Request.FormValue("_method")) == "DELETE" {

		blog, err := models.FindBlogByURL(url)
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
	blog, err := models.FindBlogByURL(url)
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
		http.Redirect(a.Response, a.Request, "/post/"+url+"/edit", 302)
		return
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
	http.Redirect(a.Response, a.Request, "/post/"+url, 302)
}

// // Show shows selected blog
func (c Blog) Show(a *helpers.Params) {
	url := strings.Split(a.Request.URL.Path, "/")[2]
	blog, err := models.FindBlogByURL(url)
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(blogs) >= 5 {
		blogs = blogs[0:5]
	}
	helpers.Render(a, "blog/show", map[string]interface{}{
		"post":  blog,
		"posts": blogs,
	})
}

// // Edit shows selected blog
func (c Blog) Edit(a *helpers.Params) {
	url := strings.Split(a.Request.URL.Path, "/")[2]
	blog, err := models.FindBlogByURL(url)
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	helpers.Render(a, "blog/edit", map[string]interface{}{
		"blog": blog,
	})
}

// // Image shows selected blog
func (c Blog) Image(a *helpers.Params) {
	imageID := strings.Split(a.Request.URL.Path, "/")[4]
	b, err := models.GetImageByID(imageID)
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	a.Response.Write(b)
}

func (c Blog) APIIndex(a *helpers.Params) {
	indexBlogs, err := models.AllBlogs()
	if err != nil {
		helpers.Logger.Println(err)
	}
	data, err := json.Marshal(indexBlogs)
	if err != nil {
		helpers.Logger.Println(err)
	}
	a.Response.Header().Set("Content-Type", "application/json")
	a.Response.Write(data)
}
