package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	ID           string       `bson:"_id,omitempty"` // Include the ID field to uniquely identify recipes
	Slug         string       `bson:"slug"`
	Image        string       `bson:"image"`
	Author       string       `bson:"author"`
	Cuisine      string       `bson:"cuisine"`
	Title        string       `bson:"title"`
	Introduction string       `bson:"introduction"`
	PrepTime     int          `bson:"prep_time"`
	CookTime     int          `bson:"cook_time"`
	Instructions string       `bson:"instructions"`
	Notes        string       `bson:"notes,omitempty"` // Use omitempty for optional fields
	PublishDate  string       `bson:"publish_date"`
	CreateDate   string       `bson:"create_date"`
	Categories   []Category   `bson:"categories"`
	Techniques   []Technique  `bson:"techniques"`
	Ingredients  []Ingredient `bson:"ingredients"`
	Difficulty   Difficulty   `bson:"difficulty"`
}

type Category struct {
	Name string `bson:"name"`
}

type Technique struct {
	Name string `bson:"name"`
}

type Ingredient struct {
	Name     string  `bson:"name"`
	Unit     string  `bson:"unit"`
	Quantity float64 `bson:"quantity"`
}

type Difficulty struct {
	Level int `bson:"level"`
}

var client *mongo.Client

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", mongoUser, mongoPassword)

	clientOptions := options.Client().ApplyURI(connectionString)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetRecipes fetches a paginated list of recipes from the database
func GetRecipes(page, pageSize int) ([]Recipe, int, error) {
	collection := client.Database("recipeDB").Collection("recipes")

	// Find the total number of recipes
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Calculate the skipping number
	skip := (page - 1) * pageSize

	// Find the recipes with pagination
	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize)))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())

	var recipes []Recipe
	if err = cursor.All(context.TODO(), &recipes); err != nil {
		return nil, 0, err
	}

	return recipes, int(total), nil
}

// GetRecipe fetches a single recipe by its ID
func GetRecipe(id string) (*Recipe, error) {
	collection := client.Database("recipeDB").Collection("recipes")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	var recipe Recipe
	err = collection.FindOne(context.TODO(), filter).Decode(&recipe)
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func GetAllRecipes() ([]Recipe, error) {
	collection := client.Database("recipeDB").Collection("recipes")

	// Find all recipes without any filter
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var recipes []Recipe
	if err = cursor.All(context.TODO(), &recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

func SearchRecipes(query string) ([]Recipe, error) {
	var matchedRecipes []Recipe

	// Get all recipes first
	recipes, err := GetAllRecipes()
	if err != nil {
		return nil, err
	}

	// Filter recipes based on search query
	for _, recipe := range recipes {
		if strings.Contains(strings.ToLower(recipe.Title), query) ||
			strings.Contains(strings.ToLower(recipe.Introduction), query) ||
			strings.Contains(strings.ToLower(recipe.Author), query) {
			matchedRecipes = append(matchedRecipes, recipe)

			// Debug prints (optional - remove in production)
			// fmt.Println("Recipe match found:")
			// fmt.Printf("Title: %s\n", recipe.Title)
			// fmt.Printf("Author: %s\n", recipe.Author)
			// fmt.Println("-------------------")
		}
	}

	return matchedRecipes, nil
}
