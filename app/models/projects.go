package models

import (
	"fmt"
	"time"

	"html/template"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

//URL data structure
type URL struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Hash     string        `bson:"hash"`
	Original string        `bson:"original"`
	New      string        `bson:"new"`
	Time     time.Time     `bson:"time"`
}

// dbProject struct
type dbProject struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    bson.ObjectId `bson:"user_id"`
	Image     string        `bson:"Image"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	Summary   string        `bson:"summary"`
	URL       string        `bson:"url"`
	CustomURL string        `bson:"customURL"`
	Tags      []string      `bson:"tags"`
	Time      time.Time     `bson:"time"`
}

// Project struct
type Project struct {
	ID        bson.ObjectId
	Author    User
	Image     string
	Title     string
	Body      template.HTML
	Summary   string
	URL       string
	CustomURL string
	Tags      []string
	Time      time.Time
}

// AllProjects Find all projects in mongoDB
func AllProjects() (projects []Project, err error) {
	var rawprojects []dbProject
	err = db.Session().DB(config.Data.Env).C("Project").Find(bson.M{}).All(&rawprojects)
	for _, e := range rawprojects {
		projects = append(projects, projectify(e))
	}
	if err != nil {
		fmt.Println("error in all projects")
	}
	return projects, err
}
func projectify(e dbProject) (project Project) {
	author, _ := FindUserByID(e.UserID)
	project = Project{
		Title:     e.Title,
		Author:    author,
		ID:        e.ID,
		Image:     e.Image,
		Body:      template.HTML(e.Body),
		Summary:   e.Summary,
		Tags:      e.Tags,
		Time:      e.Time,
		URL:       e.URL,
		CustomURL: e.CustomURL,
	}
	return project
}

//URLShortenerCreate entry into db for new short url
func URLShortenerCreate(hash string, original string, new string) (string, error) {
	session := db.Session()
	defer session.Close()
	new = "https://gro.de/1wd02"
	c := session.DB(config.Data.Env).C("URLShortener")
	// Insert Datas
	err := c.Insert(&URL{
		Hash:     hash,
		Original: original,
		New:      new,
		Time:     time.Now(),
	})
	if err != nil {
		helpers.Logger.Println(err)
		return "Shortener Failed", err
	}
	return "Shortened URL created", nil
}

// ProjectCreate creates a new project
func ProjectCreate(title string, body string, summary string, tags []string, userID bson.ObjectId, url string, Image []byte, customURL string) (string, error) {
	session := db.Session()
	defer session.Close()
	gridFS := session.DB(config.Data.Env).GridFS("fs")
	gridFile, err := gridFS.Create("")
	if err != nil {
		helpers.Logger.Println(err)
		return "This didnt work", err
	}
	defer helpers.Close(gridFile)
	_, err = gridFile.Write(Image)
	if err != nil {
		helpers.Logger.Println(err)
		return "This didnt work", err
	}
	c := session.DB(config.Data.Env).C("Project")
	// Insert Datas
	err = c.Insert(&dbProject{
		Title:     title,
		Body:      body,
		Summary:   summary,
		UserID:    userID,
		URL:       url,
		CustomURL: customURL,
		Image:     gridFile.Id().(bson.ObjectId).Hex(),
		Tags:      tags,
		Time:      time.Now(),
	})
	if err != nil {
		helpers.Logger.Println(err)
		return "This didnt work", err
	}
	return "Project Post created", nil
}

// DestroyProject ...
func DestroyProject(id bson.ObjectId) error {
	helpers.DeleteCache(string(id))
	session := db.Session()
	defer session.Close()
	return session.DB(config.Data.Env).C("Project").RemoveId(id)
}

// FindProjectByURL Returns blog by Title
func FindProjectByURL(url string) (project Project, err error) {
	projectURL := helpers.Get(url, func() *helpers.CacheObject {
		var rawproject dbProject
		session := db.Session()
		defer session.Close()
		err = session.DB(config.Data.Env).C("Project").Find(bson.M{"url": url}).One(&rawproject)
		project = projectify(rawproject)
		return helpers.NewCacheObject(project)
	})
	return projectURL.Object.(Project), err
}

// FindBlogByID ...
func findDbProjectByID(id string) (project dbProject, err error) {
	Projectdb := helpers.Get(id, func() *helpers.CacheObject {
		session := db.Session()
		defer session.Close()
		err = session.DB(config.Data.Env).C("Project").FindId(bson.ObjectIdHex(id)).One(&project)
		return helpers.NewCacheObject(project)
	})
	return Projectdb.Object.(dbProject), err
}

func findShortURLByID(id string) (url URL, err error) {
	Projectdb := helpers.Get(id, func() *helpers.CacheObject {
		session := db.Session()
		defer session.Close()
		err = session.DB(config.Data.Env).C("URLShortener").FindId(bson.ObjectIdHex(id)).One(&url)
		return helpers.NewCacheObject(url)
	})
	return Projectdb.Object.(URL), err
}

// ProjectUpdate Project Update!
func ProjectUpdate(id string, updated map[string]interface{}) error {
	helpers.DeleteCache(id)
	session := db.Session()
	defer session.Close()
	c := session.DB(config.Data.Env).C("Project")
	// Want to make each field optional how?
	newProject, err := findDbProjectByID(id)
	if err != nil {
		return err
	}
	for key, actual := range map[string]*string{
		"title":     &newProject.Title,
		"body":      &newProject.Body,
		"summary":   &newProject.Summary,
		"url":       &newProject.URL,
		"customURL": &newProject.CustomURL,
	} {
		if updated[key] != nil {
			*actual = updated[key].(string)
		}
	}
	if updated["tags"] != nil {
		newProject.Tags = updated["tags"].([]string)
	}
	if updated["Image"] != nil {
		gridFS := session.DB(config.Data.Env).GridFS("fs")
		gridFile, err := gridFS.Create("")
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		defer helpers.Close(gridFile)
		_, err = gridFile.Write(updated["Image"].([]byte))
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		newProject.Image = gridFile.Id().(bson.ObjectId).Hex()
	}
	err = c.UpdateId(bson.ObjectIdHex(id), newProject)
	if err != nil {
		helpers.Logger.Println(err)
		return err
	}
	return nil
}

// GetProjectsByTags ...
func GetProjectsByTags(searchTerm string) (projects []Project, err error) {
	projectTags := helpers.Get(searchTerm, func() *helpers.CacheObject {
		var rawProjects []dbProject
		err = db.Session().DB(config.Data.Env).C("Project").Find(bson.M{
			"$or": []bson.M{bson.M{
				"tags": searchTerm,
			},
				bson.M{
					"title": &bson.RegEx{Pattern: searchTerm, Options: "i"},
				}}}).All(&rawProjects)
		if err != nil {
			helpers.Logger.Println(err)
		}
		for _, e := range rawProjects {
			projects = append(projects, projectify(e))
		}
		if err != nil {
			helpers.Logger.Println(err)
		}
		return helpers.NewCacheObject(projects)
	})
	return projectTags.Object.([]Project), err
}
