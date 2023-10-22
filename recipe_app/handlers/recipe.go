package handlers

import (
	"net/http"
	"recipe_app/models"

	"github.com/labstack/echo/v4"
)

func GetRecipeHandler(c echo.Context) error {
	id := c.Param("id")
	recipe, err := models.GetRecipe(id)
	if err != nil || recipe == nil {
		return c.String(http.StatusNotFound, "Recipe not found")
	}
	return c.Render(http.StatusOK, "base.html", map[string]interface{}{
		"Page":   "recipe",
		"recipe": recipe,
	})
}
