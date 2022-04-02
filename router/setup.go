package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
)

func Setup(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "OK")
	})

	app := e.Group("")
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("[%v] [${time_rfc3339}] Got ${method} on ${uri} from ${remote_ip}. Responded with ${status} in ${latency_human}.\n", color.Green("INFO")), //nolint:lll,revive
	}))
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		//nolint:wrapcheck
		return c.String(http.StatusOK, "hello")
	})
}
