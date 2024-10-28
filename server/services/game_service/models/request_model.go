package models

type GameRequest struct {
	UserName string `json:"username" bson:"username" validate:"required"`
}
