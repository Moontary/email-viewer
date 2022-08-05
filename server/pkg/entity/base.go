package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Address string             `json:"address" bson:"address,omitempty" validate:"required,email"`
}
