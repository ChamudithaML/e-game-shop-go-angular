package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sales struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Stock  int                `json:"stock,omitempty" bson:"stock,omitempty" validate:"required,gte=0"`
	Sales  int                `json:"sales,omitempty" bson:"sales,omitempty" validate:"required,gte=0"`
	Profit float64            `json:"profit,omitempty" bson:"profit,omitempty" validate:"required,gte=0"`
}
