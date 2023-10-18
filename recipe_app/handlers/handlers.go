package handlers

import (
	"recipe_app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecipes(client *mongo.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := client.Database("test").Collection("recipes")
		cursor, err := collection.Find(c.Context(), bson.D{})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		var recipes []models.Recipe
		if err = cursor.All(c.Context(), &recipes); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Render("index", fiber.Map{"recipes": recipes})
	}
}
