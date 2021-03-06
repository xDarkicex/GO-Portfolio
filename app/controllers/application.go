package controllers

import (
	"fmt"
	"net/http"
	"net/mail"
	"net/smtp"

	"strconv"

	"github.com/scorredoira/email"
	"github.com/xDarkicex/PortfolioGo/app/controllers/filter"
	"github.com/xDarkicex/PortfolioGo/app/models"
	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

// Application Controller.
type Application helpers.Controller

//Index New index function
func (c Application) Index(a helpers.RouterArgs) {
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	if a.Request.Method == "HEAD" {
		a.Response.WriteHeader(200)
		return
	}

	blogs, err := models.AllBlogs()
	if err != nil {
		helpers.Logger.Println(err)
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
	err := filter.IP(a.Request)
	if err != nil {
		http.Error(a.Response, err.Error(), 403)
		return
	}
	if a.Request.Method == "HEAD" {
		a.Response.WriteHeader(200)
		return
	}
	user, err := models.FirstUser()
	if err != nil {
		panic(err)
	}
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
	m.From = mail.Address{Name: name, Address: config.Data.Email}
	m.To = []string{"grolofson@bitdev.io"}
	auth := smtp.PlainAuth("", config.Data.Email, config.Data.SMTP.Password, config.Data.SMTP.Host)
	gmailSMTP := config.Data.SMTP.Host + ":" + strconv.Itoa(config.Data.SMTP.Port)
	fmt.Println(config.Data.SMTP.Port)
	fmt.Println(gmailSMTP)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		helpers.Logger.Println(err)
	}
	http.Redirect(a.Response, a.Request, "/", 302)
}
