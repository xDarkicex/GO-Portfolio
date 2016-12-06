package models

import (
	"fmt"
	"time"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"gopkg.in/mgo.v2/bson"
)

// dbProject struct
type dbProject struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserID bson.ObjectId `bson:"user_id"`
	Image  string        `bson:"Image"`
	Title  string        `bson:"title"`
	Body   string        `bson:"body"`
	URL    string        `bson:"url"`
	Tags   []string      `bson:"tags"`
	Time   time.Time     `bson:"time"`
}

// Project struct
type Project struct {
	ID     bson.ObjectId
	Author User
	Image  string
	Title  string
	Body   string
	URL    string
	Tags   []string
	Time   time.Time
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
		Author: author,
		ID:     e.ID,
		Image:  e.Image,
		Body:   e.Body,
		Tags:   e.Tags,
		Time:   e.Time,
		URL:    e.URL,
	}
	return project
}
