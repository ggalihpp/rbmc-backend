package main

import (
	"net/http"

	"github.com/ggalihpp/rbmc-backend/auth"
	"github.com/ggalihpp/rbmc-backend/membership"
	"github.com/ggalihpp/rbmc-backend/primary"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setupHandlers(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	config := middleware.JWTConfig{
		Claims:     &primary.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	m := e.Group("/membership")

	m.Use(middleware.JWTWithConfig(config))
	membership.SetupHandler(m)

	auth.SetupHandler(e.Group("/auth"))
}
