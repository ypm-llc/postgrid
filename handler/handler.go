package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/ypm-llc/postgrid/service"
)

func ApplyRoutes(e *echo.Echo) {
	pingHandler := &PingHandler{}
	mailHandler := &MailHandler{
		SMTPService: &service.SMTPService{Server: "localhost:25"},
	}

	e.GET("/ping", pingHandler.ping)
	e.POST("/mail", mailHandler.Send)
}
