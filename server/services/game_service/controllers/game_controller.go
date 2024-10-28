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

	"go.mongodb.org/mongo-driver/bson"
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

func GetGame(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	gameId := c.Params("gameId")
	var game models.Game
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(gameId)

	err := gameCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&game)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.GameResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": game}})
}

func EditGame(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	gameId := c.Params("gameId")
	var game models.Game
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(gameId)

	//validate the request body
	if err := c.BodyParser(&game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&game); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"title":     game.Title,
		"genre":     game.Genre,
		"developer": game.Developer,
		"platform":  game.Platform,
		"price":     game.Price,
		"stock":     game.Stock,
	}

	result, err := gameCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated game details
	var updatedGame models.Game
	if result.MatchedCount == 1 {
		err := gameCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedGame)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.GameResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedGame}})
}

func DeleteGame(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	gameId := c.Params("gameId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(gameId)

	result, err := gameCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.GameResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Game with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.GameResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Game successfully deleted!"}},
	)
}

func GetAllGames(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var games []models.Game
	defer cancel()

	results, err := gameCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var oneGame models.Game
		if err = results.Decode(&oneGame); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		games = append(games, oneGame)
	}

	return c.Status(http.StatusOK).JSON(
		responses.GameResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": games}},
	)
}
