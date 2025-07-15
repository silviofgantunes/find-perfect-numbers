package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/silviofgantunes/find-perfect-numbers/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/perfect", handler.CheckPerfectNumbers)

	e.Logger.Fatal(e.Start(":8080"))
}
