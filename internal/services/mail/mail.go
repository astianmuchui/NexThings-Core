package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/gomail.v2"

	"github.com/astianmuchui/nexthings-core/internal/env"
	"github.com/astianmuchui/nexthings-core/internal/models"

)

var SenderEmail, AppPassword string
var envError error

func init() {
	env.Load()

	SenderEmail, envError = env.GetSenderEmail()
	if envError != nil {
		log.Error("Unable to load sender email from environment")
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

	send_err := d.DialAndSend(m)

	if send_err != nil {
		return send_err
	}

	return nil
}

func (e *Email) SendUserVerificationEmail(url string, u *models.User) error {
	tmplPath := filepath.Join("internal", "templates", "emails", "verify-email.html")
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		log.Infof("failed to parse template: %v", err)
		return err
	}

	var body bytes.Buffer
	data := map[string]interface{}{
		"User": &u,
		"Link": fmt.Sprintf("%s/accounts/verify/%s/%s", url, u.Uuid, u.EmailVerifyToken),
	}
	if err := tmpl.Execute(&body, data); err != nil {
		log.Error("failed to execute template: %v", err)
		return err
	}
	e.Body = body.String()
	e.Subject = fmt.Sprintf("Verify your account, %s", u.FirstName)
	e.Recepients = []string{u.Email}
	e.Mu = &sync.Mutex{}

	return e.Send()
}
