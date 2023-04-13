package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PingHandler struct{}

func (_ *PingHandler) ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
