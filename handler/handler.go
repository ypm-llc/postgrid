package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/ypm-llc/postgrid/service"
)

func ApplyRoutes(e *echo.Echo) {
	pingHandler := &PingHandler{}
	mailHandler := &MailHandler{SMTPService: service.NewSMTPService()}

	e.GET("/ping", pingHandler.Ping)
	e.POST("/mail", mailHandler.Send)
}
