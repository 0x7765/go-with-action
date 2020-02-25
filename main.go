package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/0x7765/go-with-action/routers"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	routers.Register(e)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
