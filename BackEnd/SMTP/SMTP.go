package SMTP

import (
	"gopkg.in/gomail.v2"
)

func ConnectSMTP() *gomail.Dialer {
	username := "19923a8df1073b"
	password := "9dbf6ca5188841"
	smtpHost := "sandbox.SMTP.mailtrap.io"
	smtpPort := 587
	dialer := gomail.NewDialer(smtpHost, smtpPort, username, password)
	return dialer
}
