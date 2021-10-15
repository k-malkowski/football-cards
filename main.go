package main

import (
	card "footballCards/cards"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	card.RegisterRoutes(e)
	e.Start(":3000")
}
