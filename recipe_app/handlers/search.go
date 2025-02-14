package handlers

import (
	"net/http"
	"recipe_app/models"
	"strings"

	"github.com/labstack/echo/v4"
)

func SearchRecipesHandler(c echo.Context) error {
	query := c.QueryParam("query")

	recipes, err := models.SearchRecipes(strings.ToLower(query))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error searching recipes")
	}

	// Return only the results template for HTMX requests
	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Render(http.StatusOK, "results", map[string]interface{}{
			"recipes": recipes,
			"query":   query,
		})
	}

	// Return full page for regular requests
	return c.Render(http.StatusOK, "base", map[string]interface{}{
		"Page":    "index",
		"recipes": recipes,
		"query":   query,
	})
}
