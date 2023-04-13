package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ypm-llc/postgrid/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	handler.ApplyRoutes(e)

	e.Logger.Fatal(e.Start(":6888"))
}
