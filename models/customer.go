package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	IDCustomer    string             `json:"idCustomer" bson:"id_customer"`
	Name          string             `json:"name" bson:"name"`
	CurrentPoints int                `json:"currentPoints" bson:"current_points"`
}
