package mail

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
)



type QQMailSender struct {
	name              string  // 发件人（泛指）
	fromEmailAddress  string  // 发件人邮箱地址
	fromEmailPassword string  // 发件人邮箱密码
}


func NewQQMailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *QQMailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {

	m := mail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress))  // 发件人 e.g. Simple Bank <sjxiang2023@gmail.com>
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	for _, f := range attachFiles {
		m.Attach(f)
	}

	d := mail.NewDialer("smtp.qq.com", 465, sender.fromEmailAddress, sender.fromEmailPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	
	return d.DialAndSend(m)
}