package notificator

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/xhit/go-simple-mail/v2"
)

type EmailNotificator struct {
	SMTPPort  int
	SMTPHost  string
	SMTPPass  string
	SMTPUser  string
	EmailFrom string

	TO  []string
	MSG string
}

func NewEmailNotificator(
	SMTPPort int,
	SMTPHost string,
	SMTPPass string,
	SMTPUser string,
	EmailFrom string,
) *EmailNotificator {

	return &EmailNotificator{
		SMTPPort:  SMTPPort,
		SMTPHost:  SMTPHost,
		SMTPPass:  SMTPPass,
		SMTPUser:  SMTPUser,
		EmailFrom: EmailFrom,
	}
}

type EmailData struct {
	To      string
	Token   string
	Name    string
	Subject string
	Host    string
}

var _ Notificator = (*EmailNotificator)(nil)

func (en *EmailNotificator) Send(ctx context.Context, data interface{}) error {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = en.SMTPHost
	server.Port = en.SMTPPort
	server.Username = en.SMTPUser
	server.Password = en.SMTPPass
	server.Encryption = mail.EncryptionSSL

	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		return err
	}

	emailData := data.(*EmailData)
	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(en.EmailFrom).
		AddTo(emailData.To).
		SetSubject(emailData.Subject)

	const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
	<p>Hi,%s</p>
	<p>Please verify your account to be able to login</p>
	<a href="https://%s/email/verify/%s" target="_blank">Verify your account</a>
	</body>
</html>`

	body := fmt.Sprintf(htmlBody, emailData.Name, emailData.Host, emailData.Token)

	email.SetBody(mail.TextHTML, body)

	err = email.Send(smtpClient)
	if err != nil {
		return fmt.Errorf("cannot send email: %w", err)
	} else {
		log.Println("Email Sent")
	}

	return nil
}
