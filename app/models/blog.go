package models

import (
	"fmt"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

//Blog struct for mongoDB structure
type Blog struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    string        `bson:"userID"`
	BlogImage []byte        `bson:"blogImage"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	URL       string        `bson:"url"`
}

//AllBlogs finds all blog posts
func AllBlogs() (blogs []Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{}).All(&blogs)
	if err != nil {
		fmt.Println("error in all blogs")
	}
	return blogs, err
}

// FindBlogByURL Returns blog by Title
func FindBlogByURL(url string) (blog Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{"url": url}).One(&blog)

	return blog, err
}

// BlogCreate creates a new blog post
func BlogCreate(title string, body string, userID string, url string, blogImage []byte) (string, error) {
	session := db.Session()
	defer session.Close()
	c := session.DB(config.ENV).C("Blog")
	// Insert Datas
	err := c.Insert(&Blog{
		Title:     title,
		Body:      body,
		UserID:    userID,
		URL:       url,
		BlogImage: blogImage,
	})
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
		return "This didnt work", err
	}
	return "Blog Post created", nil
}
