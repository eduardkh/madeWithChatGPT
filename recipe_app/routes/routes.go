package routes

import (
	"recipe_app/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.GetRecipesHandler)
	e.GET("/recipe/:id", handlers.GetRecipeHandler)
	e.GET("/api/search", handlers.SearchRecipesHandler)
}
