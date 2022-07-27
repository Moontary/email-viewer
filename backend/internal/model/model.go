package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email,omitempty" validate:"required,email"`
}
