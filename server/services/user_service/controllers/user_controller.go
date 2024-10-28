package controllers

import (
	"context"
	"net/http"
	"time"
	"user_service/configs"
	"user_service/models"
	"user_service/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user_data")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	hashedPassword, hash_err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if hash_err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "Could not hash password", Data: &fiber.Map{"data": hash_err.Error()}})
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Country:  user.Country,
		Role:     user.Role,
		Username: user.Username,
		Password: string(hashedPassword),
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "User Added", Data: &fiber.Map{"data": result}})
}

// Secret key for JWT signing
var secretKey = []byte("chama123")

func LoginUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Message: "Invalid request"})
	}

	// Find the user in the database
	var newUser models.User
	err := userCollection.FindOne(c.Context(), bson.M{"username": user.Username}).Decode(&newUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{Message: "Invalid credentials"})
		}
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Message: "Error fetching user"})
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{Message: "Invalid credentials"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": newUser.Username,                      // can use newUser.Id
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Message: "Could not generate token"})
	}

	// Return the token
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
