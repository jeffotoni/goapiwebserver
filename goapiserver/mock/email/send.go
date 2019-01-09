package main

import "github.com/jeffotoni/goapiwebserver/goapiserver/pkg/email"

func main() {

	fromName := "jeff"
	fromEmail := "jeff.otoni@gmail.com"
	toNames := []string{"Jefferson"}
	toEmails := []string{"jeff.otoni@s3wf.com"}
	subject := "This is the subject of your email"

	email.SendMail(fromName, fromEmail, toNames, toEmails, subject, "Hello, test jeff email local.")

}
