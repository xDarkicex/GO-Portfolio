package models

import (
	// "github.com/xDarkicex/Portfolienfig"

	"errors"
	"fmt"

	s "github.com/xDarkicex/GO-CLASS/lazy"
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
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	// Address  Address
}

// CreateUser create a new user in the database
func CreateUser(email string, name string, password string) (bool, string) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hashing Password incomplete")
		return false, "Encryption Failure"
	}
	fmt.Println("Success")
	session := db.Session()
	defer session.Close()
	c := session.DB(config.ENV).C("User")
	amount, _ := c.Find(bson.M{"name": name}).Count()
	if amount > 0 {
		return false, "Username already exists."
	}
	// Insert Datas
	err = c.Insert(&User{
		Email:    email,
		Name:     name,
		Password: string(hashedPass),
	})
	if err != nil {
		s.Say("Can not Insert User")
		return false, "Can't tell if in yet"
	}
	return true, "User created"
}

// GetUser for things
func GetUser(name string, password string) (user User, err error) {
	fmt.Println(name)
	fmt.Println(password)
	s := db.Session()
	defer s.Close()
	err = s.DB(config.ENV).C("User").Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return user, errors.New("Here is no user with this name/password combination")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Here is no user with this name/password combination")
	} else {
		return user, nil
	}
}
