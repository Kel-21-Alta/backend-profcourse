package smtp

import (
	"context"
	"gopkg.in/gomail.v2"
)

type SmtpEmail struct {
	ConfigSmtpHost     string
	ConfigSmtpPort     int
	ConfigSenderName   string
	ConfigAuthEmail    string
	ConfigAuthPassword string
}

func (s SmtpEmail) SendEmail(ctx context.Context, to string, subject string, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.ConfigAuthEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		s.ConfigSmtpHost,
		s.ConfigSmtpPort,
		s.ConfigAuthEmail,
		s.ConfigAuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func NewSmtpEmail(config SmtpEmail) *SmtpEmail {
	return &SmtpEmail{
		config.ConfigSmtpHost,
		config.ConfigSmtpPort,
		config.ConfigSenderName,
		config.ConfigAuthEmail,
		config.ConfigAuthPassword,
	}
}
