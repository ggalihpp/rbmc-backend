package main

import (
	"net/http"

	example "github.com/ggalihpp/go-echo-boilerplate/route-example"
	"github.com/labstack/echo"
)

func setupHandlers(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	exampleRoute := e.Group("/example")
	example.SetupHandler(exampleRoute)
}
