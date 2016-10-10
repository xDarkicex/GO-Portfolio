package models

import (
	// "github.com/xDarkicex/Portfolienfig"

	"errors"
	"fmt"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	"github.com/xDarkicex/PortfolioGo/helpers"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

////////////////////////////////////////////////////////////
//Address of user plan to implement at some time in future
////////////////////////////////////////////////////////////

// Address This will be for user profiles geolocation
type Address struct {
	Zip    string `bson:"zip"`
	State  string `bson:"state"`
	City   string `bson:"city"`
	Street string `bson:"street"`
}

//User Struct
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Admin    bool          `bson:"admin"`
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
	session := db.Session()
	defer session.Close()
	admin := false
	c := session.DB(config.ENV).C("User")
	amount, _ := c.Count()
	if amount == 0 {
		admin = true
	}
	amount, _ = c.Find(bson.M{"name": name}).Count()
	if amount > 0 {
		return false, "Username already exists."
	}
	// Insert Datas
	err = c.Insert(&User{
		Email:    email,
		Name:     name,
		Admin:    admin,
		Password: string(hashedPass),
	})
	if err != nil {
		fmt.Println("Can not Insert User")
		return false, "Can't tell if in yet"
	}
	return true, "User created"
}

// Login as a user
func Login(name string, password string) (user User, err error) {
	// fmt.Println(name)
	// fmt.Println(password)
	// Not sure if I should close DB here?
	// defer s.Close()
	s := db.Session()
	defer s.Close()
	err = s.DB(config.ENV).C("User").Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return user, errors.New("Here is no user with this name/password combination")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Here is no user with this name/password combination")
	}
	return user, nil
}

//FindUserByName finds a user by name
func FindUserByName(name string) (user User, err error) {
	err = db.Session().DB(config.ENV).C("User").Find(bson.M{"name": name}).One(&user)
	return user, err
}

// FindUserByID ...
func FindUserByID(userID string) (user User, err error) {
	err = db.Session().DB(config.ENV).C("User").FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		helpers.Logger.Println(err)
	}
	return user, err
}

//AllUsers finds all the users
func AllUsers() (users []User, err error) {
	err = db.Session().DB(config.ENV).C("User").Find(bson.M{}).All(&users)
	return users, err
}

// GetUser Authenticates User access
// func GetUser(req *http.Request) (user User, err error) {
// 	s := db.Session()
// 	idCookie, err := req.Cookie("id")
// 	if err != nil {
// 		return user, err
// 	}
// 	err = s.DB(config.ENV).C("User").FindId(bson.ObjectIdHex(idCookie.Value)).One(&user)
// 	return user, err
// }
