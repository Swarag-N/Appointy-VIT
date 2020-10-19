package user

import "go.mongodb.org/mongo-driver/bson/primitive"

//User DB Model.
type User struct {
	Name           string               `json:"name"`
	Email          string               `json:"email"`
	FutureMeetings []primitive.ObjectID `json:"future_meetings"`
}

//Users return List
type Users []User
