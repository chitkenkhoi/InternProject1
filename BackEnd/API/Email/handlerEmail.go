package Email

import (
	"gopkg.in/gomail.v2"
	"project/SMTP"
)

func SendEmail(list_recipient []string, subject string, body string) bool {
	// Tạo đối tượng Gomail
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "chitkenkhoi@gmail.com") // Địa chỉ email người gửi   // Địa chỉ email người nhận
	mailer.SetHeader("To", list_recipient...)         // Địa chỉ email người nhận
	mailer.SetHeader("Subject", subject)              // Tiêu đề email
	mailer.SetBody("text/plain", body)                // Nội dung email

	// Cấu hình kết nối SMTP
	dialer := SMTP.ConnectSMTP()

	// Gửi email
	if err := dialer.DialAndSend(mailer); err != nil {
		return false
	} else {
		return true
	}
}
