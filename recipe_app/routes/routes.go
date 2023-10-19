package routes

import (
	"recipe_app/handlers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, client *mongo.Client) {
	app.Get("/", handlers.GetRecipes(client))
	app.Get("/recipe/:id", handlers.GetRecipe(client))
}
