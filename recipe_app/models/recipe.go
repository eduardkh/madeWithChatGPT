package models

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
