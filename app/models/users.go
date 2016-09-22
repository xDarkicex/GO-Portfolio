package models

import (
	// "github.com/xDarkicex/Portfolienfig"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//Address of user
type Address struct {
	Zip    string
	State  string
	City   string
	Street string
}

//User Struct
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string
	Email    string
	Password string
	// Address  Address
}

// CreateUser create a new user in the database
func CreateUser(email string, name string, password string) bool {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	c := db.Session.DB(config.ENV).C("user")
	// Insert Datas
	err = c.Insert(&User{
		Email:    email,
		Name:     name,
		Password: string(hashedPass),
	})

	if err != nil {
		panic(err)
	}

	return true
}
