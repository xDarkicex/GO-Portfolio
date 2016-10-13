package models

import (
	"fmt"
	"time"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

//Blog struct for mongoDB structure
type Blog struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    string        `bson:"userID"`
	BlogImage string        `bson:"blogImage"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	URL       string        `bson:"url"`
	Tags      []string      `bson:"tags"`
	Time      time.Time     `bson:"time"`
}

// VideoBlog ..
type VideoBlog struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	URL       string        `bson:"url"`
	BlogVideo string        `bson:"BlogVideo"`
	Time      time.Time     `bson:"time"`
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

// FindBlogByID ...
func FindBlogByID(id string) (blog Blog, err error) {
	err = db.Session().DB(config.ENV).C("Blog").FindId(bson.ObjectIdHex(id)).One(&blog)
	return blog, err
}

// BlogCreate creates a new blog post
func BlogCreate(title string, body string, tags []string, userID string, url string, blogImage []byte) (string, error) {
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
	err = c.Insert(&Blog{
		Title:     title,
		Body:      body,
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
	newPost, err := FindBlogByID(id)
	if err != nil {
		return err
	}
	for key, actual := range map[string]*string{
		"title":  &newPost.Title,
		"body":   &newPost.Body,
		"userID": &newPost.UserID,
		"url":    &newPost.URL,
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

// BlogUpdate creates a new blog post
// func BlogUpdate(title string, body string, tags []string, id string, userID string, url string, blogImage []byte) error {
// 	fmt.Println("I made it too update function")
// 	session := db.Session()
// 	defer session.Close()
// 	gridFS := session.DB(config.ENV).GridFS("fs")
// 	gridFile, err := gridFS.Create("")
// 	if err != nil {
// 		helpers.Logger.Println(err)
// 		return err
// 	}

// 	defer helpers.Close(gridFile)
// 	_, err = gridFile.Write(blogImage)
// 	if err != nil {
// 		helpers.Logger.Println(err)
// 		return err
// 	}

// 	c := session.DB(config.ENV).C("Blog")
// 	// Update Data currently is making new posts not updating, Also
// 	// Want to make each field optional how?

// 	err = c.UpdateId(bson.ObjectIdHex(id),
// 		bson.M{"$set": &Blog{
// 			Title:     title,
// 			Body:      body,
// 			UserID:    userID,
// 			URL:       url,
// 			BlogImage: gridFile.Id().(bson.ObjectId).Hex(),
// 			Tags:      tags,
// 		},
// 		})
// 	if err != nil {
// 		helpers.Logger.Println(err)
// 		return err
// 	}
// 	return nil
// }

// GetImageByID ...
func GetImageByID(imageID string) ([]byte, error) {
	gridFS := db.Session().DB(config.ENV).GridFS("fs")
	gridFile, err := gridFS.OpenId(bson.ObjectIdHex(imageID))
	if err != nil {
		helpers.Logger.Println(err)
		fmt.Println("Can not Create New Blog post")
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
