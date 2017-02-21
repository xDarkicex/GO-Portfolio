package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"net/smtp"

	"strconv"

	"github.com/scorredoira/email"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Application Controller.
type Application helpers.Controller

//Index New index function
func (c Application) Index(a helpers.RouterArgs) {
	fmt.Println(a.Request.Host)
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(blogs) >= 5 {
		blogs = blogs[0:5]
	}
	helpers.Render(a, "application/index", map[string]interface{}{
		"blog": blogs,
	})
}

//About About me Pages
func (c Application) About(a helpers.RouterArgs) {
	var user models.User
	if len(helpers.Cache["firstUser"]) == 0 {
		user, err := models.FirstUser()
		if err != nil {
			panic(err)
		}
		userify, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		helpers.Cache["firstUser"] = string(userify)
	}
	byteUser := []byte(helpers.Cache["firstUser"])
	json.Unmarshal(byteUser, &user)
	helpers.Render(a, "application/about", map[string]interface{}{
		"user":  user,
		"title": "About Me",
	})
}

//Contact form function
func (c Application) Contact(a helpers.RouterArgs) {

	name := (a.Request.FormValue("contactName"))
	address := (a.Request.FormValue("contactAddress"))
	body := (a.Request.FormValue("contactBody"))
	subject := "Message From " + name + " - " + address
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: "From", Address: config.Data.Email}
	m.To = []string{"grolofson@bitdev.io"}
	auth := smtp.PlainAuth("", config.Data.Email, config.Data.SMTP.Password, config.Data.SMTP.Host)
	gmailSMTP := config.Data.SMTP.Host + ":" + strconv.Itoa(config.Data.SMTP.Port)
	fmt.Println(config.Data.SMTP.Port)
	fmt.Println(gmailSMTP)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		fmt.Println(err)
	}
	http.Redirect(a.Response, a.Request, "/", 302)
}
