package utils

import (
	"net/smtp"

	"gihtub.com/heyyakash/switchr/internal/modals"
)

var (
	smtpHost  = GetString("SMTP_HOST")
	smtpPort  = GetString("SMTP_PORT")
	smtpEmail = GetString("SMTP_EMAIL")
	smtpPass  = GetString("SMTP_PASSWORD")
)

func SendEmail(mail *modals.Email) error {
	auth := smtp.PlainAuth("", smtpEmail, smtpPass, smtpHost)
	message := []byte(mail.Subject + "\r\n" + mail.Content)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, []string{mail.To}, []byte(message))
	return err
}
