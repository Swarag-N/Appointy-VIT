package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// Meeting Model of Mongo DB
type Meeting struct {
	ID    primitive.ObjectID
	title string
}

//Movie Model
type Movie struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Year      string             `json:"year" bson:"year"`
	Directors []string           `json:"directors" bson:"directors"`
	Writers   []string           `json:"writers" bson:"writers"`
}
