package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

type SendMailModel struct {
	SendTo          string
	User            string
	UserID          int
	Token           string
	ArticleUser     string
	ArticleName     string
	LikeUser        string
	CommentUser     string
	ChilCommentUser string
}

func SendMail(templatePath string, data SendMailModel, subject string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(&body, data)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SEND_FROM_ADDRESS"))
	m.SetHeader("To", data.SendTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SEND_FROM_ADDRESS"), os.Getenv("MAIL_PASSWORD"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
