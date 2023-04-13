package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ypm-llc/postgrid/handler"
)

func main() {
	godotenv.Load()

	e := echo.New()
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = handler.ErrorHandler

	handler.ApplyRoutes(e)

	e.Logger.Fatal(e.Start(":6888"))
}
