package service

import (
	"github.com/ypm-llc/postgrid/infra"
	"github.com/ypm-llc/postgrid/message"
)

type SMTPService struct {
	client *infra.SMTPClient
	Server string
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
