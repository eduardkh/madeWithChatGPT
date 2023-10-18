package models

type Recipe struct {
	Title        string `bson:"title"`
	Introduction string `bson:"introduction"`
}
