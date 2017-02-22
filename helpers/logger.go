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

type errorLog struct {
}

var errBuf = bytes.NewBuffer([]byte{})

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

func (e errorLog) Write(p []byte) (n int, err error) {
	if config.Data.Verbose {
		fmt.Print("Error: " + string(p))
	}
	errBuf.Write(p)
	return n, err
}

// Logger is a helpper method to print out a more useful error message
var Logger = log.New(errorLog{}, "", log.Lmicroseconds|log.Lshortfile)

// func sendMSG(msg string) {
// 	ircobj := irc.IRC("PortfolioGo", "golang") //Create new ircobj
// 	ircobj.Connect("irc.bitdev.io:6667")       //Connect to server
// 	errors := strings.Split(msg, "\n")
// 	ircobj.AddCallback("001", func(e *irc.Event) {
// 		ircobj.Join("#notifier")
// 		time.Sleep(1 * time.Second)
// 		for k, v := range errors {
// 			if len(v) > 0 {
// 				ircobj.Privmsg("#notifier", fmt.Sprintf("Error %d: %s", k+1, v))
// 			}
// 		}
// 		time.Sleep(1 * time.Second)
// 		ircobj.Disconnect()
// 	})
// }
func sendSMS(msg string) {
	name := "PortfolioGo"
	address := "127.168.0.1"
	body := msg
	subject := ("Message From " + name + " - " + string(address))
	m := email.NewMessage(subject, body)
	m.From = mail.Address{Name: name, Address: address}
	m.To = []string{config.Data.SMTP.Number}
	auth := smtp.PlainAuth("", config.Data.Email, config.Data.SMTP.Password, config.Data.SMTP.Host)
	gmailSMTP := config.Data.SMTP.Host + ":" + fmt.Sprintf("%d", config.Data.SMTP.Port)
	if err := email.Send(gmailSMTP, auth, m); err != nil {
		log.Fatal(err)
	}
}
