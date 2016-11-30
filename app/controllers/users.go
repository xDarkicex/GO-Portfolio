package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Users Controller!
type Users helpers.Controller

// Index function
func (c Users) Index(a helpers.RouterArgs) {
	users, err := models.AllUsers()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// fmt.Printf("Users %s", users)
	helpers.Render(a, "users/index", map[string]interface{}{
		"users": users,
	})
}

// Create a new user
func (c Users) Create(a helpers.RouterArgs) {
	session := a.Session
	success, createErr := models.CreateUser(a.Request.FormValue("email"), a.Request.FormValue("name"), a.Request.FormValue("password"))
	if success {
		user, err := models.Login(a.Request.FormValue("name"), a.Request.FormValue("password"))
		if err != nil {
			fmt.Println(err)
		} else {
			session.Values["UserID"] = user.ID.Hex()
			helpers.AddFlash(a, helpers.Flash{Type: "success", Message: "User Created Successful"})
			err = session.Save(a.Request, a.Response)
			if err != nil {
				helpers.Logger.Println(err)
			}
			http.Redirect(a.Response, a.Request, "/users/"+user.Name, 302)
			return
		}
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	helpers.AddFlash(a, helpers.Flash{Type: "danger", Message: createErr})
	err := session.Save(a.Request, a.Response)
	if err != nil {
		helpers.Logger.Println(err)
	}
	http.Redirect(a.Response, a.Request, "/register", http.StatusFound)
}

// Show Show page for users
func (c Users) Show(a helpers.RouterArgs) {
	user, err := models.FindUserByName(a.Params.ByName("name"))
	if err != nil {
		http.Redirect(a.Response, a.Request, "/404", 302)
	}
	helpers.Render(a, "users/show", map[string]interface{}{
		"user": user,
	})
}

// New ...
func (c Users) New(a helpers.RouterArgs) {
	if a.Session.Values["UserID"] == nil {
		helpers.Render(a, "users/new", map[string]interface{}{})
	} else {
		http.Redirect(a.Response, a.Request, "/", 302)
	}
}

// Update ...
func (c Users) Update(a helpers.RouterArgs) {
	if len(a.Request.FormValue("_method")) > 0 && string(a.Request.FormValue("_method")) == "DELETE" {
		user, err := models.FindUserByName(a.Params.ByName("name"))
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		// Actually update
		err = models.UserDestroy(user.ID)
		if err != nil {
			helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
			return
		}
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	user, err := models.FindUserByName(a.Params.ByName("name"))
	if err != nil {
		fmt.Println(err)
	}
	newUser := map[string]interface{}{}
	for _, key := range []string{"fullname", "age", "skills", "experiance", "bio", "file", "country", "state", "city", "street", "zip"} {
		value := a.Request.FormValue(key)
		if len(value) > 0 {
			newUser[key] = value
		}
	}
	// Get file
	file, _, err := a.Request.FormFile("file")
	if err == nil {
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		} else {
			newUser["Avatar"] = fileBytes
		}
	}
	// Actually update
	err = models.UserUpdate(user.ID.Hex(), newUser)
	if err != nil {
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	http.Redirect(a.Response, a.Request, "/users/"+user.Name, 302)
}

// Edit shows selected user profile
func (c Users) Edit(a helpers.RouterArgs) {
	user, err := models.FindUserByName(a.Params.ByName("name"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return

	}
	if a.User == user {
		helpers.Render(a, "users/edit", map[string]interface{}{
			"user": user,
		})
	} else {
		http.Redirect(a.Response, a.Request, "/", 302)
	}
}

//Image ..
func (c Users) Image(a helpers.RouterArgs) {
	b, err := models.GetImageByID(a.Params.ByName("imageID"))
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	a.Response.Write(b)
}
