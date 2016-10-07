package models

import (
	"encoding/json"
	"fmt"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

//Blog struct for mongoDB structure
type Blog struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"title"`
	Body  string        `bson:"body"`
}

//AllBlogs finds all blog posts
func AllBlogs() (blogs []Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{}).All(&blogs)
	json.Marshal(blogs)
	fmt.Println("error in all blogs")
	return blogs, err
}

// FindBlogByTitle Returns blog by Title
func FindBlogByTitle(title string) (blog Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{"title": title}).One(&blog)

	return blog, err
}

// BlogCreate creates a new blog post
func BlogCreate(title string, body string) (bool, string) {
	session := db.Session()
	defer session.Close()
	c := session.DB(config.ENV).C("Blog")
	// Insert Datas
	err := c.Insert(&Blog{
		Title: title,
		Body:  body,
	})
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
		return false, "This didnt work"
	}
	return true, "Blog Post created"
}
