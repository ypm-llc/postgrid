package service

import (
	"os"

	"github.com/ypm-llc/postgrid/infra"
	"github.com/ypm-llc/postgrid/message"
)

type SMTPService struct {
	client *infra.SMTPClient
	Server string
}

func NewSMTPService() *SMTPService {
	smtpServer := "127.0.0.1:25"
	v := os.Getenv("SMTP_SERVER")
	if v != "" {
		smtpServer = v
	}
	return &SMTPService{Server: smtpServer}
}

func (s *SMTPService) Send(m *message.Mail) error {
	if s.client == nil {
		s.client = &infra.SMTPClient{Server: s.Server}
		err := s.client.Connect()
		if err != nil {
			return err
		}
	}
	defer s.client.Disconnect()
	return s.client.Send(m)
}

func (s *SMTPService) BulkSend(m []*message.Mail) []error {
	if s.client == nil {
		s.client = &infra.SMTPClient{Server: s.Server}
		err := s.client.Connect()
		if err != nil {
			return []error{err}
		}
	}
	defer s.client.Disconnect()
	errList := []error{}
	for _, mail := range m {
		err := s.client.Send(mail)
		if err != nil {
			errList = append(errList, err)
		}
	}
	return errList
}
