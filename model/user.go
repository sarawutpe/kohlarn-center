package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName" default:""`
	LastName  string             `bson:"lastName" json:"lastName" default:""`
	Email     string             `bson:"email" json:"email" default:""`
}
