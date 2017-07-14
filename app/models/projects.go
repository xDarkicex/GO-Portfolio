package models

import (
	"fmt"
	"time"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"gopkg.in/mgo.v2/bson"
)

// dbProject struct
type dbProject struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  bson.ObjectId `bson:"user_id"`
	Image   string        `bson:"Image"`
	Title   string        `bson:"title"`
	Body    string        `bson:"body"`
	Summary string        `bson:"summary"`
	URL     string        `bson:"url"`
	Tags    []string      `bson:"tags"`
	Time    time.Time     `bson:"time"`
}

// Project struct
type Project struct {
	ID      bson.ObjectId
	Author  User
	Image   string
	Title   string
	Body    string
	Summary string
	URL     string
	Tags    []string
	Time    time.Time
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
		Author:  author,
		ID:      e.ID,
		Image:   e.Image,
		Body:    e.Body,
		Summary: e.Summary,
		Tags:    e.Tags,
		Time:    e.Time,
		URL:     e.URL,
	}
	return project
}

// ProjectCreate creates a new project
func ProjectCreate(title string, body string, summary string, tags []string, userID bson.ObjectId, url string, Image []byte) (string, error) {
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
		Title:   title,
		Body:    body,
		Summary: summary,
		UserID:  userID,
		URL:     url,
		Image:   gridFile.Id().(bson.ObjectId).Hex(),
		Tags:    tags,
		Time:    time.Now(),
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
