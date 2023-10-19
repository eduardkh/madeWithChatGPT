package handlers

import (
	"recipe_app/models"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Index handler to render the index page with pagination
func GetRecipesHandler(c echo.Context) error {
	pageStr := c.QueryParam(strings.ToLower("page"))
	pageSizeStr := c.QueryParam(strings.ToLower("pageSize"))

	// Convert page and pageSize to integers
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// Provide default values if page or pageSize are not provided or are invalid
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 3
	}

	// Fetch the recipes from the database
	recipes, err := models.GetRecipes(page, pageSize)
	if err != nil {
		return c.String(500, "Internal Server Error")
	}

	// Render the index page with the fetched recipes
	return c.Render(200, "index.html", map[string]interface{}{
		"recipes": recipes,
	})
}
