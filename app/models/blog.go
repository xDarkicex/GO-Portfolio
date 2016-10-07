package models

import (
	s "github.com/xDarkicex/GO-CLASS/lazy"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"gopkg.in/mgo.v2/bson"
)

//Blog struct for mongoDB structure
type Blog struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"title"`
	Body  string        `bson:"body"`
}

//AllBlogs finds all the users
func AllBlogs() (blogs []Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{}).All(&blogs)
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
		s.Say("Can not Create New Blog post")
		return false, "This didnt work"
	}
	return true, "Blog Post created"
}
