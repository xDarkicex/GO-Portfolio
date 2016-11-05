package models

import (
	"fmt"
	"time"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

//dbBlog struct for MongoDB structure
type dbBlog struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    bson.ObjectId `bson:"user_id"`
	BlogImage string        `bson:"blogImage"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	Summary   string        `bson:"summary"`
	URL       string        `bson:"url"`
	Tags      []string      `bson:"tags"`
	Time      time.Time     `bson:"time"`
}

//Blog struct for our structure
type Blog struct {
	ID        bson.ObjectId
	Author    User
	BlogImage string
	Title     string
	Summary   string
	Body      string
	URL       string
	Tags      []string
	Time      time.Time
}

//AllBlogs finds all blog posts
func AllBlogs() (blogs []Blog, err error) {
	var rawblogs []dbBlog
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{}).All(&rawblogs)
	for _, e := range rawblogs {
		blogs = append(blogs, blogify(e))
	}
	if err != nil {
		fmt.Println("error in all blogs")
	}
	return blogs, err
}

// So blogify will turn a database blog structure into the more code friendly Blog structure with relationships resolved to actual objects.
// Your user_id will instead be Author
func blogify(e dbBlog) (blog Blog) {
	author, err := FindUserByID(e.UserID)
	if err != nil {
		fmt.Println("Error in Finding an author for blog: " + e.URL)
	}
	blog = Blog{
		Author:    author,
		ID:        e.ID,
		Title:     e.Title,
		BlogImage: e.BlogImage,
		Body:      e.Body,
		Summary:   e.Summary,
		Tags:      e.Tags,
		Time:      e.Time,
		URL:       e.URL,
	}
	return blog
}

// FindBlogByURL Returns blog by Title
func FindBlogByURL(url string) (blog Blog, err error) {
	var rawblog dbBlog
	err = db.Session().DB(config.ENV).C("Blog").Find(bson.M{"url": url}).One(&rawblog)
	blog = blogify(rawblog)
	return blog, err
}

// FindBlogByID ...
func FindBlogByID(id string) (blog Blog, err error) {
	var rawblog dbBlog
	err = db.Session().DB(config.ENV).C("Blog").FindId(bson.ObjectIdHex(id)).One(&rawblog)
	blog = blogify(rawblog)
	return blog, err
}

// FindBlogByID ...
func findDbBlogByID(id string) (blog dbBlog, err error) {
	// var rawblog dbBlog
	err = db.Session().DB(config.ENV).C("Blog").FindId(bson.ObjectIdHex(id)).One(&blog)
	// blog = blogify(rawblog)
	return blog, err
}

// BlogCreate creates a new blog post
func BlogCreate(title string, body string, summary string, tags []string, userID bson.ObjectId, url string, blogImage []byte) (string, error) {
	session := db.Session()
	defer session.Close()
	gridFS := session.DB(config.ENV).GridFS("fs")
	gridFile, err := gridFS.Create("")
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
		return "This didnt work", err
	}
	defer helpers.Close(gridFile)
	_, err = gridFile.Write(blogImage)
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
		return "This didnt work", err
	}
	c := session.DB(config.ENV).C("Blog")
	// Insert Datas
	err = c.Insert(&dbBlog{
		Title:     title,
		Body:      body,
		Summary:   summary,
		UserID:    userID,
		URL:       url,
		BlogImage: gridFile.Id().(bson.ObjectId).Hex(),
		Tags:      tags,
		Time:      time.Now(),
	})
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
		return "This didnt work", err
	}
	return "Blog Post created", nil
}

// BlogDestroy Blog Destroy
func BlogDestroy(id bson.ObjectId) error {
	session := db.Session()
	defer session.Close()
	return session.DB(config.ENV).C("Blog").RemoveId(id)
}

// BlogUpdate Blog Update!
func BlogUpdate(id string, updated map[string]interface{}) error {
	session := db.Session()
	defer session.Close()
	c := session.DB(config.ENV).C("Blog")
	// Update Data currently is making new posts not updating, Also
	// Want to make each field optional how?
	newPost, err := findDbBlogByID(id)
	if err != nil {
		return err
	}
	for key, actual := range map[string]*string{
		"title":   &newPost.Title,
		"body":    &newPost.Body,
		"summary": &newPost.Summary,
		"url":     &newPost.URL,
	} {
		if updated[key] != nil {
			*actual = updated[key].(string)
		}
	}
	if updated["tags"] != nil {
		newPost.Tags = updated["tags"].([]string)
	}
	if updated["blogImage"] != nil {
		gridFS := session.DB(config.ENV).GridFS("fs")
		gridFile, err := gridFS.Create("")
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		defer helpers.Close(gridFile)
		_, err = gridFile.Write(updated["blogImage"].([]byte))
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		newPost.BlogImage = gridFile.Id().(bson.ObjectId).Hex()
	}
	err = c.UpdateId(bson.ObjectIdHex(id), newPost)
	if err != nil {
		helpers.Logger.Println(err)
		return err
	}
	return nil
}

// GetImageByID ...
func GetImageByID(imageID string) ([]byte, error) {
	gridFS := db.Session().DB(config.ENV).GridFS("fs")
	gridFile, err := gridFS.OpenId(bson.ObjectIdHex(imageID))
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can't find image")
		return nil, err
	}
	defer gridFile.Close()

	b := make([]byte, gridFile.Size())
	_, err = gridFile.Read(b)
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Error encoding image")
		return nil, err
	}
	return b, nil
}

// GetBlogsByTags ...
func GetBlogsByTags(searchTerm string) ([]Blog, error) {
	var blogs []Blog
	err := db.Session().DB(config.ENV).C("Blog").Find(bson.M{
		"tags": searchTerm,
	}).All(&blogs)
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Error locating blog by tag, no tag found")
		return nil, err
	}
	return blogs, err
}
