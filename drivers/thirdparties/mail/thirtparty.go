package mail

import (
	"context"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Host       string
	Port       int
	SenderName string
	AuthEmail  string
	Password   string
}

func (s Email) SendEmail(ctx context.Context, to string, subject string, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.AuthEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		s.Host,
		s.Port,
		s.AuthEmail,
		s.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func NewSmtpEmail(config Email) *Email {
	return &Email{
		config.Host,
		config.Port,
		config.SenderName,
		config.AuthEmail,
		config.Password,
	}
}
