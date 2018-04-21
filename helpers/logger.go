package helpers

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"bytes"

	"github.com/scorredoira/email"
	"github.com/xDarkicex/PortfolioGo/config"
)

// simple empty struct to link to Wrtie functionality
type errorLog struct{}
type silentLog struct{}

var errBuf = bytes.NewBuffer([]byte{})
var silentBuf = bytes.NewBuffer([]byte{})

// FlushLog Flush queue to file, sms and msg.
func FlushLog() {
	length := errBuf.Len()
	if length > 0 {
		file, _ := os.OpenFile("log/"+config.Data.Env+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		buffed := errBuf.String()
		file.WriteString(buffed)
		sendSMS(buffed)
		errBuf.Reset()
		file.Close()
	}
}

// FlushSilentLog Flush queue to file, sms and msg.
func FlushSilentLog() {
	length := silentBuf.Len()
	if length > 0 {
		file, _ := os.OpenFile("log/network.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		buffed := silentBuf.String()
		file.WriteString(buffed)
		silentBuf.Reset()
		file.Close()
	}
}

// over write default write behavior
func (e errorLog) Write(p []byte) (n int, err error) {
	if config.Data.Verbose {
		fmt.Println("Error: " + string(p))
	}
	errBuf.Write(p)
	return n, err
}

// over write default write behavior
func (e silentLog) Write(p []byte) (n int, err error) {
	if config.Data.Verbose {
		fmt.Println(string(p))
	}
	silentBuf.Write(p)
	return n, err
}

// Logger is a helpper method to print out a more useful error message
var Logger = log.New(errorLog{}, "", log.Lmicroseconds|log.Lshortfile)

// SilentLogger is for logging without SMS
var SilentLogger = log.New(silentLog{}, "", log.Lmicroseconds|log.Lshortfile)

func sendSMS(msg string) {
	name := "PortfolioGo"
	address := "rolofson.me"
	body := msg
	subject := ("Message From " + name + " - " + string(address))
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: name, Address: address}
	m.To = []string{config.Data.SMTP.Number}
	auth := smtp.PlainAuth("", config.Data.Email, config.Data.SMTP.Password, config.Data.SMTP.Host)
	gmailSMTP := config.Data.SMTP.Host + ":" + fmt.Sprintf("%d", config.Data.SMTP.Port)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		log.Println(err)
	}
}
