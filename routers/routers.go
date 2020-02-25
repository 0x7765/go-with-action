package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register all
func Register(echo *echo.Echo) {
	echo.GET("/", home)
	echo.GET("/ping", ping)
}

func ping(context echo.Context) error {
	return context.JSON(http.StatusOK, "pong")
}

func home(context echo.Context) error {
	return context.HTML(http.StatusOK, "This is index page...")
}
