package models

import (
	"fmt"
	"time"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"gopkg.in/mgo.v2/bson"
)

// dbExample struct
type dbExample struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	UserID       bson.ObjectId `bson:"user_id"`
	ExampleImage string        `bson:"blogImage"`
	Title        string        `bson:"title"`
	Body         string        `bson:"body"`
	URL          string        `bson:"url"`
	Tags         []string      `bson:"tags"`
	Time         time.Time     `bson:"time"`
}

// Example struct
type Example struct {
	ID           bson.ObjectId
	Author       User
	ExampleImage string
	Title        string
	Body         string
	URL          string
	Tags         []string
	Time         time.Time
}

// AllExamples Find all examples in mongoDB
func AllExamples() (examples []Example, err error) {
	var rawexamples []dbExample
	err = db.Session().DB(config.ENV).C("Example").Find(bson.M{}).All(&rawexamples)
	for _, e := range rawexamples {
		examples = append(examples, examplify(e))
	}
	if err != nil {
		fmt.Println("error in all examples")
	}
	return examples, err
}
func examplify(e dbExample) (example Example) {
	author, _ := FindUserByID(e.UserID)
	example = Example{
		Author:       author,
		ID:           e.ID,
		ExampleImage: e.ExampleImage,
		Body:         e.Body,
		Tags:         e.Tags,
		Time:         e.Time,
		URL:          e.URL,
	}
	return example
}
