package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	c.JSON(http.StatusInternalServerError, map[string]string{
		"message": err.Error(),
	})
}
