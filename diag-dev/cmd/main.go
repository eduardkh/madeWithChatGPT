package main

import (
	"diagportal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// serve static assets
	e.Static("/static", "static")

	// load templates
	e.Renderer = handlers.NewRenderer("templates/*.html")

	// tab panel endpoints
	e.GET("/", handlers.Index)
	e.GET("/diag/ping", handlers.PingPanel)
	e.GET("/diag/dns", handlers.DNSPanel)
	e.GET("/diag/trace", handlers.TracePanel)

	// action endpoints
	e.POST("/ping", handlers.Ping)
	e.POST("/dns", handlers.DNS)
	e.POST("/trace", handlers.Trace)

	e.Logger.Fatal(e.Start(":8080"))
}
