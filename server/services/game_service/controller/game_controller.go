package controllers

import (
	"context"
	"game_service/configs"
	"game_service/models"
	"game_service/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var gameCollection *mongo.Collection = configs.GetCollection(configs.DB, "game_data")
var validate = validator.New()

func AddGame(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var game models.Game
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&game); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newGame := models.Game{
		Id:        primitive.NewObjectID(),
		Title:     game.Title,
		Genre:     game.Genre,
		Developer: game.Developer,
		Platform:  game.Platform,
		Price:     game.Price,
		Stock:     game.Stock,
	}

	result, err := gameCollection.InsertOne(ctx, newGame)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.GameResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
