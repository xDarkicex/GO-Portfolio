package models

import (
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"gopkg.in/mgo.v2/bson"
)

//Blog struct for mongoDB structure
type Blog struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"title"`
	Body  string        `bson:"body"`
	Admin bool          `bson:"admin"`
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
