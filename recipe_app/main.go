package main

import (
	"recipe_app/handlers"
	"recipe_app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = handlers.TemplateRenderer()
	routes.RegisterRoutes(e)

	// Serve static files
	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
