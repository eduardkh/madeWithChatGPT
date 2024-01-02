package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		// To work with partials (views)
		Views: html.New("./views", ".gohtml"), // configure the view engine
	})

	// Serve static files
	app.Static("/", "./public")

	// Search endpoint for HTMX the query
	app.Get("/search", func(c *fiber.Ctx) error {
		query := strings.ToLower(c.Query("query"))

		var filteredUsers []map[string]interface{}

		// Only proceed with the API call if the query is not empty
		if query != "" {
			// Query the API endpoint
			resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			var users []map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&users)

			// Loop through the response
			for _, user := range users {
				name := strings.ToLower(user["name"].(string))
				email := strings.ToLower(user["email"].(string))
				if strings.Contains(name, query) || strings.Contains(email, query) {
					filteredUsers = append(filteredUsers, user)
				}
			}
		}

		// Render response for the HTMX frontend using partials (results.gohtml)
		return c.Render("results", fiber.Map{
			"users": filteredUsers,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
