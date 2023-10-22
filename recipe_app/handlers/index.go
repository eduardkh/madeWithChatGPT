package handlers

import (
	"net/http"
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
		pageSize = 8
	}

	// Fetch the recipes from the database
	recipes, total, err := models.GetRecipes(page, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Calculate the total number of pages
	totalPages := (total + pageSize - 1) / pageSize

	// Render the index page with the fetched recipes and pagination data
	return c.Render(http.StatusOK, "base.html", map[string]interface{}{
		"Page":       "index",
		"recipes":    recipes,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": totalPages,
	})
}
