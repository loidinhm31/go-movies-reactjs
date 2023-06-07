package service

import (
	"context"
	"fmt"
	simple_mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"movies-service/config"
	"movies-service/internal/common/model"
	"movies-service/internal/control"
	"movies-service/internal/mail"
	"os"
	"strings"
	"time"
)

type mailService struct {
	config *config.Config

	mgmtCtrl control.Service
}

func NewMovieService(config *config.Config, mgmtCtrl control.Service) mail.Service {
	return &mailService{
		config:   config,
		mgmtCtrl: mgmtCtrl,
	}
}

func (ms *mailService) SendMessage(ctx context.Context, m *model.MailData) error {
	server := simple_mail.NewSMTPClient()
	server.Host = ms.config.Mail.Host
	server.Port = ms.config.Mail.Port
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		log.Println(err)
		return err
	}

	email := simple_mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.TemplateMail == "" {
		email.SetBody(simple_mail.TextHTML, m.Content)
	} else {
		var resourcesPath string
		if ms.config.Environment != "local" {
			resourcesPath = "./resources/email-templates"
		} else {
			resourcesPath = "/etc/resources/email-templates"
		}

		data, err := os.ReadFile(fmt.Sprintf("%s/%s", resourcesPath, m.TemplateMail))

		if err != nil {
			log.Println(err)
			return err
		}

		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(simple_mail.TextHTML, msgToSend)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Email sent")
	return nil
}
