package infra

import (
	"io"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/ypm-llc/postgrid/message"
)

// Line Break
const LineBreak = "\n"

// SMTPClient is a SMTP client.
type SMTPClient struct {
	*smtp.Client
	Server string
}

// Connect connects to the SMTP server.
func (c *SMTPClient) Connect() error {
	client, err := smtp.Dial(c.Server)
	if err != nil {
		return err
	}
	client.CommandTimeout = 5 * time.Second
	client.SubmissionTimeout = 5 * time.Second

	c.Client = client
	return nil
}

// Disconnect disconnects from the SMTP server.
func (c *SMTPClient) Disconnect() error {
	return c.Quit()
}

// Send sends a mail.
func (c *SMTPClient) Send(m *message.Mail) error {
	err := c.writeSenderAndRecipients(m)
	if err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	err = c.writeHeaders(wc, m)
	if err != nil {
		return err
	}

	_, err = wc.Write([]byte(m.Body))
	return err
}

// write headers to the writeCloser
func (c *SMTPClient) writeHeaders(wc io.WriteCloser, m *message.Mail) error {
	_, err := wc.Write([]byte("Subject: " + m.Subject + LineBreak))
	if err != nil {
		return err
	}

	toList := strings.Join(m.To, ",")
	_, err = wc.Write([]byte("To: " + toList + LineBreak))
	if err != nil {
		return err
	}

	if len(m.Cc) > 0 {
		ccList := strings.Join(m.Cc, ",")
		_, err = wc.Write([]byte("Cc: " + ccList + LineBreak))
		if err != nil {
			return err
		}
	}

	if len(m.Bcc) > 0 {
		bccList := strings.Join(m.Bcc, ",")
		_, err = wc.Write([]byte("Bcc: " + bccList + LineBreak))
		if err != nil {
			return err
		}
	}

	_, err = wc.Write([]byte(LineBreak))
	return err
}

// write a sender and recipients to the writeCloser
func (c *SMTPClient) writeSenderAndRecipients(m *message.Mail) error {
	c.Mail(m.From, nil)
	for _, to := range m.To {
		err := c.Rcpt(to)
		if err != nil {
			return err
		}
	}
	for _, cc := range m.Cc {
		err := c.Rcpt(cc)
		if err != nil {
			return err
		}
	}
	for _, bcc := range m.Bcc {
		err := c.Rcpt(bcc)
		if err != nil {
			return err
		}
	}
	return nil
}
