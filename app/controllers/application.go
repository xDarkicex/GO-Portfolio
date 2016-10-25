package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"

	"github.com/julienschmidt/httprouter"
	"github.com/scorredoira/email"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Application Controller.
type Application helpers.Controller

//Index New index function
func (c Application) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	blogs, err := models.AllBlogs()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if len(blogs) >= 5 {
		blogs = blogs[0:5]
	}
	view := "application/index"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"blog":   blogs,
	})
}

//About About me Pages
func (c Application) About(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	session, err := helpers.Store().Get(a.Request, "user-session")
	if err != nil {
		helpers.Logger.Println(err)
		http.Redirect(a.Response, a.Request, "/", 302)
		return
	}
	user, err := models.FindUserByName("xDarkicex")
	if err != nil {
		fmt.Println("There was an error")
	}
	view := "application/about"
	helpers.Render(a, view, map[string]interface{}{
		"UserID": session.Values["UserID"],
		"user":   user,
	})
}

//Contact form function
func (c Application) Contact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
	name := (a.Request.FormValue("contactName"))
	address := (a.Request.FormValue("contactAddress"))
	body := (a.Request.FormValue("contactBody"))
	subject := "Message From " + name + " - " + address
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: "From", Address: config.EMAIL}
	m.To = []string{"grolofson@bitdev.io"}
	auth := smtp.PlainAuth("", config.EMAIL, config.SMTPPASSWORD, config.SMTPHOST)
	gmailSMTP := config.SMTPHOST + ":" + config.SMTPPORT
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		log.Fatal(err)
	}
	http.Redirect(a.Response, a.Request, "/", 302)
}
