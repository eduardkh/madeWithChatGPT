package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	ID           string       `bson:"_id,omitempty"` // Include the ID field to uniquely identify recipes
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
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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
	collection := client.Database("test").Collection("recipes")

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
	collection := client.Database("test").Collection("recipes")
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
