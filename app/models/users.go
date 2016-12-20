package models

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

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

// dbUser Struct
type dbUser struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name,omitempty"`
	Avatar       string        `bson:"avatar,omitempty"`
	FullName     string        `bson:"fullname,omitempty"`
	Skills       string        `bson:"skills,omitempty"`
	Experience   string        `bson:"experience,omitempty"`
	Bio          string        `bson:"bio,omitempty"`
	Admin        bool          `bson:"admin"`
	Email        string        `bson:"email"`
	Password     string        `bson:"password"`
	Country      string        `bson:"country,omitempty"`
	Language     string        `bson:"language,omitempty"`
	Zip          string        `bson:"zip,omitempty"`
	State        string        `bson:"state,omitempty"`
	City         string        `bson:"city,omitempty"`
	Street       string        `bson:"street,omitempty"`
	LoginAttempt int           `bson:"loginattempt,omitempty"`
}

//User Struct
type User struct {
	ID           bson.ObjectId
	Name         string
	Avatar       string
	FullName     string
	Skills       string
	Experience   string
	Bio          template.HTML
	Admin        bool
	Email        string
	Password     string
	Country      string
	Language     string
	Zip          string
	State        string
	City         string
	Street       string
	LoginAttempt int
}

func userify(e dbUser) (user User) {
	user = User{
		ID:           e.ID,
		Name:         e.Name,
		Avatar:       e.Avatar,
		FullName:     e.FullName,
		Skills:       e.Skills,
		Experience:   e.Experience,
		Bio:          template.HTML(e.Bio),
		Admin:        e.Admin,
		Email:        e.Email,
		Password:     e.Password,
		Country:      e.Country,
		Language:     e.Language,
		Zip:          e.Zip,
		State:        e.State,
		City:         e.City,
		Street:       e.Street,
		LoginAttempt: e.LoginAttempt,
	}
	return user
}

// CreateUser create a new user in the database
func CreateUser(email string, name string, password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters!"
	}
	if strings.Contains(password, " ") {
		return false, "Password must not contain spaces!"
	}
	if !strings.ContainsAny(password, "1234567890") {
		return false, "Password must contain at least one number!"
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		code := helpers.RandStringRunes(5)
		err := fmt.Sprintf("Error Hashing Password incomplete: %s\nHash Code: %s", err, code)
		helpers.Logger.Println(err)
		return false, fmt.Sprintf("Contact Site webmaster internal server error %s", code)
	}
	session := db.Session()
	defer session.Close()
	admin := false
	c := session.DB(config.Data.Env).C("User")
	amount, err := c.Count()
	if err != nil {
		helpers.Logger.Println(err)
	}
	if amount == 0 {
		admin = true
	}
	amount, err = c.Find(bson.M{"name": name}).Count()
	if err != nil {
		// Hash CODE for internal server errors are for things I need to refrence later cant show users.
		code := helpers.RandStringRunes(5)
		err := fmt.Sprintf("Error: %s\nHash Code: %s", err, code)
		helpers.Logger.Println(err)
		return false, fmt.Sprintf("Contact Site webmaster internal server error %s", code)
	}
	if amount > 0 {
		err := fmt.Errorf("Username taken: %s", err)
		helpers.Logger.Println(err)
		return false, err.Error()
	}
	expression, err := regexp.Compile("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$")
	if !expression.MatchString(email) {
		err := fmt.Errorf("Not vaild Email: %s", err)
		helpers.Logger.Println(err)
		return false, err.Error()
	}
	github, err := http.Get("http://github.com/" + name)
	if github.StatusCode == 404 {
		err := fmt.Errorf("Not vaild Github username: %s", err)
		helpers.Logger.Println(err)
		return false, err.Error()
	}
	// Insert Datas
	err = c.Insert(&dbUser{
		Email:    email,
		Name:     name,
		Admin:    admin,
		Password: string(hashedPass),
	})
	if err != nil {
		err := fmt.Errorf("Error inserting new User into User collection: %s", err)
		helpers.Logger.Println(err)
		return false, err.Error()
	}
	return true, "User created"
}

// Login as a user
func Login(name string, password string) (user User, err error) {
	s := db.Session()
	defer s.Close()
	var rawUser dbUser
	err = s.DB(config.Data.Env).C("User").Find(bson.M{"name": name}).One(&rawUser)
	user = userify(rawUser)
	if err != nil {
		helpers.Logger.Println(err)
		err := fmt.Errorf("There is no user with this username/password combination %s", err)
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		helpers.Logger.Println(err)
		err := fmt.Errorf("There is no user with this Username/password combination %s", err)
		return user, err
	}
	return user, nil
}

//FindUserByName finds a user by name
func FindUserByName(name string) (user User, err error) {
	var rawUser dbUser
	err = db.Session().DB(config.Data.Env).C("User").Find(bson.M{"name": name}).One(&rawUser)
	if err != nil {
		helpers.Logger.Println(err)
		helpers.Logger.Println(string(debug.Stack()))
	}
	user = userify(rawUser)
	return user, err
}

// FindUserByID ...
func FindUserByID(userID bson.ObjectId) (user User, err error) {
	var rawUser dbUser
	err = db.Session().DB(config.Data.Env).C("User").FindId(userID).One(&rawUser)
	if err != nil {
		helpers.Logger.Println(err)
	}
	user = userify(rawUser)
	return user, err
}

// FinddbUserByID ...
func finddbUserByID(id string) (user dbUser, err error) {
	err = db.Session().DB(config.Data.Env).C("User").FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		helpers.Logger.Println(err)
	}
	return user, err
}

//AllUsers finds all the users
func AllUsers() (users []User, err error) {
	var rawUsers []dbUser
	err = db.Session().DB(config.Data.Env).C("User").Find(bson.M{}).All(&rawUsers)
	if err != nil {
		helpers.Logger.Println(err)
	}
	for _, e := range rawUsers {
		users = append(users, userify(e))
	}
	return users, err
}

// UserDestroy Blog Destroy
func UserDestroy(id bson.ObjectId) error {
	session := db.Session()
	defer session.Close()
	return session.DB(config.Data.Env).C("User").RemoveId(id)
}

// UserUpdate Update!
func UserUpdate(id string, updated map[string]interface{}) error {
	session := db.Session()
	defer session.Close()
	c := session.DB(config.Data.Env).C("User")
	// Update Data
	newUser, err := finddbUserByID(id)
	if err != nil {
		helpers.Logger.Println(err)
		return err
	}
	for key, actual := range map[string]*string{
		"fullname":   &newUser.FullName,
		"skills":     &newUser.Skills,
		"bio":        &newUser.Bio,
		"experience": &newUser.Experience,
		"password":   &newUser.Password,
		"zip":        &newUser.Zip,
		"state":      &newUser.State,
		"city":       &newUser.City,
		"street":     &newUser.Street,
	} {
		if updated[key] != nil {
			*actual = updated[key].(string)
		}
	}
	if updated["Avatar"] != nil {
		gridFS := session.DB(config.Data.Env).GridFS("fs")
		gridFile, err := gridFS.Create("")
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		defer helpers.Close(gridFile)
		_, err = gridFile.Write(updated["Avatar"].([]byte))
		if err != nil {
			helpers.Logger.Println(err)
			return err
		}
		newUser.Avatar = gridFile.Id().(bson.ObjectId).Hex()
	}
	err = c.UpdateId(bson.ObjectIdHex(id), newUser)
	if err != nil {
		helpers.Logger.Println(err)
		return err
	}
	return nil
}

// FirstUser This returns the first user created
func FirstUser() (user User, err error) {
	var rawUser dbUser
	err = db.Session().DB(config.Data.Env).C("User").Find(bson.M{}).One(&rawUser)
	if err != nil {
		helpers.Logger.Println(err)
	}
	user = userify(rawUser)
	return user, err
}
