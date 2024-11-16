package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Genre     string             `json:"genre,omitempty" bson:"genre,omitempty" validate:"required"`
	Developer string             `json:"developer,omitempty" bson:"developer,omitempty" validate:"required"`
	Platform  string             `json:"platform,omitempty" bson:"platform,omitempty" validate:"required"`
	Price     float64            `json:"price,omitempty" bson:"price,omitempty" validate:"required,gt=0"`
	Stock     int                `json:"stock,omitempty" bson:"stock,omitempty" validate:"required,gte=0"`
}
