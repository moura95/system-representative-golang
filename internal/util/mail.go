package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jordan-wright/email"

	"my-orders/internal/repository"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown fromServer")
		}
	}
	return nil, nil
}

type Data = struct {
	Title string
	Name  string
	Msg   string
}

type EmailSender interface {
	SendEmail(
		subject string,
		data Data,
		templateName string,
		to []string,
		cc []string,
		bcc []string,
		attachFile []byte,
	) error
}

type Credential struct {
	name              string
	fromEmailLogin    string
	fromEmailAddress  string
	fromEmailPassword string
	fromEmailHost     string
	fromEmailPort     int32
}

func NewSender(name string, smtp repository.Smtp) EmailSender {
	if !smtp.IsActive {
		smtp = repository.Smtp{
			Email:    "nao-responder@midasgestor.com.br",
			Password: "cER7hN4tvXUE",
			Server:   "smtp.zoho.eu",
			Port:     587,
		}
	}
	return &Credential{
		name:              name,
		fromEmailLogin:    smtp.Email,
		fromEmailAddress:  smtp.Email,
		fromEmailPassword: smtp.Password,
		fromEmailHost:     smtp.Server,
		fromEmailPort:     smtp.Port,
	}
}

func (sender *Credential) SendEmail(
	subject string,
	data Data,
	templateName string,
	to []string,
	cc []string,
	bcc []string,
	attachFile []byte,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject

	templatePath := filepath.Join("templates/" + templateName + ".html")
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read email template file: %w", err)
	}
	tmpl, err := template.New(templateName).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, data)
	if err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}
	e.HTML = content.Bytes()
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	if attachFile != nil {
		buffer := bytes.NewBuffer(attachFile)
		filename := "pedido.pdf"
		contentType := "application/pdf"

		_, err = e.Attach(buffer, filename, contentType)
		if err != nil {
			return fmt.Errorf("failed to execute e template: %w", err)
		}
	}

	smtpAuth := LoginAuth(sender.fromEmailLogin, sender.fromEmailPassword)

	return e.Send(sender.fromEmailHost+":"+strconv.Itoa(int(sender.fromEmailPort)), smtpAuth)
}
