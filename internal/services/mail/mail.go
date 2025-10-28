package mail

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/gomail.v2"

	"github.com/astianmuchui/nexthings-core/internal/env"

)

var SenderEmail, AppPassword string
var envError error

func init() {
	env.Load()

	SenderEmail, envError = env.GetSenderEmail()
	if envError != nil {
		log.Errorf("Unable to load sender email from environment")
	}

	AppPassword, envError = env.GetAppPassword()
	if envError != nil {
		log.Errorf("Unable to load Email App Password from environment")
	}
}

type EmailCC struct {
	Email string
	Name  string
}

type Email struct {
	Recepients []string
	Body       string
	Subject    string
	EmailCCs   []EmailCC
	Mu         *sync.Mutex
}

func (e *Email) Send() error {
	e.Mu.Lock()
	defer e.Mu.Unlock()

	m := gomail.NewMessage()

	m.SetHeader("From", SenderEmail)
	m.SetHeader("To", e.Recepients...)

	for _, cc := range e.EmailCCs {
		m.SetAddressHeader("Cc", cc.Email, cc.Name)
	}

	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, SenderEmail, AppPassword)

	err := make(chan error)

	go func() {
		send_err := d.DialAndSend(m)

		if send_err != nil {
			err <- send_err
		}
	}()

	return <-err
}
