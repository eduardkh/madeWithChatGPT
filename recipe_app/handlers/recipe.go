package handlers

import (
	"net/http"
	"recipe_app/models"

	"github.com/labstack/echo/v4"
)

func GetRecipeHandler(c echo.Context) error {
	id := c.Param("id")
	recipe, err := models.GetRecipe(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.Render(http.StatusOK, "recipe.html", recipe)
}
