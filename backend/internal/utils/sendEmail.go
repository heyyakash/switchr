package utils

import (
	"context"
	"crypto/tls"
	"log"
	"net/smtp"

	"gihtub.com/heyyakash/switchr/internal/modals"
	"github.com/resend/resend-go/v2"
)

var (
	smtpHost  = GetString("SMTP_HOST")
	smtpPort  = GetString("SMTP_PORT")
	smtpEmail = GetString("SMTP_EMAIL")
	smtpPass  = GetString("SMTP_PASSWORD")
)

func SendEmailold(mail *modals.Email) error {
	mode := GetString("ENV")
	if mode == "prod" {
		// Set up authentication information.
		auth := smtp.PlainAuth("", smtpEmail, smtpPass, smtpHost)

		// Create a TLS configuration.
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		}

		// Establish a TLS connection to the SMTP server.
		conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
		if err != nil {
			log.Printf("Failed to dial SMTP server: %v", err)
			return err
		}
		defer conn.Close()

		// Create a new SMTP client over the TLS connection.
		client, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			log.Printf("Failed to create SMTP client: %v", err)
			return err
		}
		defer client.Quit()

		// Authenticate with the SMTP server.
		if err = client.Auth(auth); err != nil {
			log.Printf("Failed to authenticate: %v", err)
			return err
		}

		// Set the sender and recipient addresses.
		if err = client.Mail(smtpEmail); err != nil {
			log.Printf("Failed to set sender email: %v", err)
			return err
		}

		if err = client.Rcpt(mail.To); err != nil {
			log.Printf("Failed to set recipient email: %v", err)
			return err
		}

		// Send the email content.
		writer, err := client.Data()
		if err != nil {
			log.Printf("Failed to send email data: %v", err)
			return err
		}

		_, err = writer.Write([]byte("Subject: " + mail.Subject + "\r\n" + mail.Content))
		if err != nil {
			log.Printf("Failed to write email content: %v", err)
			return err
		}

		err = writer.Close()
		if err != nil {
			log.Printf("Failed to close email writer: %v", err)
			return err
		}

		log.Print("Email sent successfully")
		return nil

	}

	auth := smtp.PlainAuth("", smtpEmail, smtpPass, smtpHost)
	message := []byte(mail.Subject + "\r\n" + mail.Content)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, []string{mail.To}, []byte(message))
	return err
}

func SendEmail(mail *modals.Email) error {
	apikey := GetString("RESEND_KEY")
	ctx := context.TODO()
	client := resend.NewClient(apikey)
	params := &resend.SendEmailRequest{
		From:    "Switchr <switchr@byakash.dev>",
		To:      []string{mail.To},
		Subject: mail.Subject,
		Html:    mail.Content,
	}

	_, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		log.Print("Failed to send email : ", err)
		return err
	}

	return nil
}
