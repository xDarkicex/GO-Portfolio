package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"

	"gopkg.in/mgo.v2/bson"

	"github.com/scorredoira/email"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Application Controller.
type Application helpers.Controller

//Index New index function
func (c Application) Index(a helpers.RouterArgs) {
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
	user, err := models.FindUserByID(bson.ObjectIdHex("580c1a8e0d89b87abd33df91"))
	if err != nil {
		fmt.Println("There was an error")
	}
	helpers.Render(a, "application/about", map[string]interface{}{
		"user": user,
	})
}

//Contact form function
func (c Application) Contact(a helpers.RouterArgs) {
	name := (a.Request.FormValue("contactName"))
	address := (a.Request.FormValue("contactAddress"))
	body := (a.Request.FormValue("contactBody"))
	subject := "Message From " + name + " - " + address
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: "From", Address: config.EMAIL}
	m.To = []string{"grolofson@bitdev.io"}
	auth := smtp.PlainAuth("", config.EMAIL, config.SMTPPASSWORD, config.SMTPHOST)
	gmailSMTP := config.SMTPHOST + ":" + string(config.SMTPPORT)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		log.Fatal(err)
	}
	http.Redirect(a.Response, a.Request, "/", 302)
}
