package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Oluwaseun241/wallet/models"
	"github.com/go-mail/mail"
)

type EmailData struct {
  URL string
  FirstName string
  Subject string
}

func SendEmail(user *models.User, data *EmailData) {
  // Sender data.
	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
  smtpUser := os.Getenv("SMTP_USER")
	to := user.Email
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
  
  smtpPort, err := strconv.Atoi(smtpPortStr)
  if err != nil {
		panic(err)
	}

  m := mail.NewMessage()

  m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
  
  m.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nClick the following link to verify your email: %s", data.FirstName, data.URL))
	m.SetBody("text/html", fmt.Sprintf(`<html><body>
		<p>Hello %s,</p>
		<p>Click the following link to verify your email: <a href="%s">%s</a></p>
	</body></html>`, data.FirstName, data.URL, data.URL))

  d := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
  if err := d.DialAndSend(m); err != nil {
    panic(err)
  }
}
