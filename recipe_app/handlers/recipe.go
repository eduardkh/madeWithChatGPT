package handlers

import (
	"recipe_app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecipe(client *mongo.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		collection := client.Database("test").Collection("recipes")
		var recipe models.Recipe
		err = collection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&recipe)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Render("recipe", fiber.Map{"recipe": recipe})
	}
}
