package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ypm-llc/postgrid/message"
	"github.com/ypm-llc/postgrid/service"

	"net/http"
)

type MailHandler struct {
	SMTPService *service.SMTPService
}

func (h *MailHandler) Send(c echo.Context) error {
	mail := &message.Mail{}
	if err := c.Bind(mail); err != nil {
		return err
	}
	err := h.SMTPService.Send(mail)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")
}
