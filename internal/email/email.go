package email

import (
	"log"
	"net/smtp"
	"os"
	"sync"
)

var (
	email Config
	once  sync.Once
)

// Config holds information related to sending an email.
type Config struct {
	Auth smtp.Auth
	Addr string
	To   string
}

// Send sends an email to one recipient.
func (e Config) Send(from, msg string) {
	msg = "To: " + e.To + "\r\n" + msg
	err := smtp.SendMail(e.Addr, e.Auth, from, []string{e.To}, []byte(msg))
	if err != nil {
		log.Printf("could not send email: %s", err)
	}
}

// IsValid verifies whether the Config struct is configured. If it is not, then
// the user did not set the MAIL_* variables under the .env file.
func (e Config) IsValid() bool {
	return e.Addr != "" && e.To != ""
}

// Email initializes the Email variable.
func Email() Config {
	once.Do(func() {
		host := os.Getenv("MAIL_HOST")
		to := os.Getenv("MAIL_TO")

		email = Config{
			Auth: smtp.PlainAuth("", to, os.Getenv("MAIL_PASSWORD"), host),
			Addr: host + ":" + os.Getenv("MAIL_PORT"),
			To:   to,
		}
	})
	return email
}