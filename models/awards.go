package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Award struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Points int                `json:"points" bson:"points"`
}
