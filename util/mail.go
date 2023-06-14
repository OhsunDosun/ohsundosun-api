package util

import (
	"net/smtp"
	"os"
)

func SendMail(email string, title string, content string) {

	auth := smtp.PlainAuth("", os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PWD"), "smtp.gmail.com")

	from := os.Getenv("EMAIL_ADDRESS")
	to := []string{email}

	msgFrom := "From: 오순도순 <" + os.Getenv("EMAIL_ADDRESS") + ">\r\n"
	msgTo := "To: " + email + "\r\n"
	msgHeaderSubject := "Subject: [오순도순] " + title + "\r\n"
	msgHeaderBlank := "\r\n"
	msgBody := content + "\r\n"
	msg := []byte(msgFrom + msgTo + msgHeaderSubject + msgHeaderBlank + msgBody)

	smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
}
