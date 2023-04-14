package handler

import (
	"strings"

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

func (h *MailHandler) BulkSend(c echo.Context) error {
	mes := &message.Mails{}
	if err := c.Bind(mes); err != nil {
		return err
	}
	mails := mes.Items
	errList := h.SMTPService.BulkSend(mails)
	if len(errList) > 0 {
		errMesList := []string{}
		for _, err := range errList {
			errMesList = append(errMesList, err.Error())
		}
		errMes := strings.Join(errMesList, ", ")
		return c.String(http.StatusInternalServerError, errMes)
	}
	return c.String(http.StatusOK, "ok")
}
