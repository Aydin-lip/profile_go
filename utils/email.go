package utils

import (
	"bytes"
	"html/template"
	"log"
	"userProfile/config"

	"gopkg.in/gomail.v2"
)

type Email struct {
	To           string
	Subject      string
	Body         string
	TemplatePath string
	Template     *template.Template
	Data         any
	Attach       []string
	IsHTML       bool
}

func SendEmail(data Email) error {
	template := data.Template

	if data.TemplatePath != "" {
		tmpl, err := template.ParseFiles(data.TemplatePath)
		if err != nil {
			return err
		}
		template = tmpl
	}

	return sendGomail(Email{
		To:           data.To,
		Subject:      data.Subject,
		Body:         data.Body,
		TemplatePath: data.TemplatePath,
		Template:     template,
		Data:         data.Data,
		Attach:       data.Attach,
		IsHTML:       data.IsHTML,
	})
}

func sendGomail(data Email) error {
	glUsername := config.GetEnv("GMAIL_USERNAME", "")
	glPassword := config.GetEnv("GMAIL_PASSWORD", "")
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := gomail.NewMessage()
	m.SetHeader("From", glUsername)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", data.Subject)

	var bodyContent string
	if data.Template != nil {
		var buf bytes.Buffer
		if err := data.Template.Execute(&buf, data.Data); err != nil {
			return err
		}
		bodyContent = buf.String()
	} else {
		bodyContent = data.Body
	}
	contentType := "text/plain"
	if data.IsHTML {
		contentType = "text/html"
	}
	m.SetBody(contentType, bodyContent)

	for _, file := range data.Attach {
		m.Attach(file)
	}

	d := gomail.NewDialer(smtpHost, smtpPort, glUsername, glPassword)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent to", data.To)
	return nil
}
