package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "hello")
	})
}
