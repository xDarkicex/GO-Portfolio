package helpers

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/scorredoira/email"
	irc "github.com/thoj/go-ircevent"
	"github.com/xDarkicex/PortfolioGo/config"
)

type errorLog struct {
}

func (e errorLog) Write(p []byte) (n int, err error) {
	if config.Data.Verbose {
		fmt.Println("Error: " + string(p))
	}
	file, _ := os.OpenFile("log/"+config.Data.Env+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(string(p))
	sendMSG(string(p))
	sendSMS(string(p))
	// Close the file when the surrounding function exists
	defer file.Close()

	return n, err
}

type shutDownLog struct {
}

func (e shutDownLog) Write(p []byte) (n int, err error) {
	if config.Data.Verbose {
		fmt.Println("Server: " + string(p))
	}
	file, _ := os.OpenFile("log/"+config.Data.Env+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(string(p))
	sendMSG(string(p))
	sendSMS(string(p))
	// Close the file when the surrounding function exists
	defer file.Close()

	return n, err
}

// Logger is a helpper method to print out a more useful error message
var Logger = log.New(errorLog{}, "", log.Lmicroseconds|log.Lshortfile)

// ShutDown log of server shutdown
var ShutDown = log.New(shutDownLog{}, "", log.Ltime)

func sendMSG(msg string) {
	ircobj := irc.IRC("PortfolioGo", "golang") //Create new ircobj
	ircobj.Connect("irc.bitdev.io:6667")       //Connect to server
	ircobj.AddCallback("001", func(e *irc.Event) {
		ircobj.Join("#notifier")
		ircobj.Privmsg("#notifier", msg)
		ircobj.Disconnect()
	})
}
func sendSMS(msg string) {
	name := "PortfolioGo"
	address := "127.168.0.1"
	body := msg
	subject := ("Message From " + name + " - " + string(address))
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: name, Address: address}
	m.To = []string{"5596760527@txt.att.net"}
	auth := smtp.PlainAuth("", config.Data.Email, config.Data.SMTP.Password, config.Data.SMTP.Host)
	gmailSMTP := config.Data.SMTP.Host + ":" + string(config.Data.SMTP.Port)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		log.Fatal(err)
	}
}
