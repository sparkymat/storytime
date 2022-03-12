package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/storytime/router"
)

func main() {
	e := echo.New()

	router.Setup(e)

	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
