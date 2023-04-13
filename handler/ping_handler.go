package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingHandler struct{}

func (_ *PingHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
