package utils

import (
	"gopkg.in/gomail.v2"
)

type email struct {
	FromUser     string
	FromPassword string
	Host         string
	Port         int
}

func NewEmail(fromUser, fromPassword, host string, port int) *email {
	return &email{
		fromUser,
		fromPassword,
		host,
		port,
	}
}

func (e email) SendMail(toUser []string, userName, subject, body string, attach []string) error {
	m := gomail.NewMessage()

	if len(userName) == 0 {
		m.SetHeader("From", e.FromUser)
	} else {
		m.SetHeader("From", m.FormatAddress(e.FromUser, userName))
	}
	m.SetHeader("To", toUser...)
	//抄送显示
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//添加附件
	for _, v := range attach {
		m.Attach(v)
	}

	d := gomail.NewDialer(e.Host, e.Port, e.FromUser, e.FromPassword)
	if err := d.DialAndSend(m); err != nil {
		Logger.Infoln("-----------SendMail", toUser[0], err)
		return err
	}

	return nil
}
