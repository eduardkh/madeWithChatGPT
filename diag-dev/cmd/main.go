package main

import (
	"diagportal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static assets (htmx, etc)
	e.Static("/static", "static")

	// HTML templates
	e.Renderer = handlers.NewRenderer("templates/*.html")

	// routes
	e.GET("/", handlers.Index)
	e.POST("/ping", handlers.Ping)
	e.POST("/dns", handlers.DNS)
	e.POST("/trace", handlers.Trace)

	e.Logger.Fatal(e.Start(":8080"))
}
