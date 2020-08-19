package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail() error {
	auth := smtp.PlainAuth("",UserMail,AuthCode,MailSmtpHost)
	to := []string{MailTo}
	user := UserMail
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte(MailSubject+"\r\n"+MailBody+"\r\n"+UrlReply+"\r\n"+UrlZhihu+"\r\n"+
		content_type)

	err := smtp.SendMail(MailSmtpHost+MailSmtpPort,auth,user,to,msg)
	if err != nil {
		fmt.Printf("send mail failed:%v\n",err)
	}
	log.Println("send mail done.")
	return err
}
