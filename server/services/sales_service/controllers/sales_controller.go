package controllers

import (
	"context"
	"fmt"
	"net/http"
	"sales_service/configs"
	"sales_service/models"
	"sales_service/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// "go.mongodb.org/mongo-driver/bson"
	"sales_service/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var salesCollection *mongo.Collection = configs.GetCollection(configs.DB, "sales_data")
var validate = validator.New()

func AddSalesData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var sale models.Sales
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&sale); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SalesResponse{Status: http.StatusBadRequest, Message: "error"})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&sale); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SalesResponse{Status: http.StatusBadRequest, Message: "error"})
	}

	newSale := models.Sales{
		Id:     primitive.NewObjectID(),
		Title:  sale.Title,
		Stock:  sale.Stock,
		Sales:  sale.Sales,
		Profit: sale.Profit,
	}

	_, err := salesCollection.InsertOne(ctx, newSale)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SalesResponse{Status: http.StatusInternalServerError, Message: "error"})
	}

	return c.Status(http.StatusCreated).JSON(responses.SalesResponse{Status: http.StatusCreated, Message: "success"})
}

func GetAllSales(c *fiber.Ctx) error {
	fmt.Println("Req came")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var sales []models.Sales
	defer cancel()

	results, err := salesCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SalesResponse{Status: http.StatusInternalServerError, Message: "error"})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var oneSale models.Sales
		if err = results.Decode(&oneSale); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SalesResponse{Status: http.StatusInternalServerError, Message: "error"})
		}

		sales = append(sales, oneSale)
	}

	title_array := services.ExtractNames(sales)
	stocks_array := services.ExtractSales(sales)
	sales_array := services.ExtractSales(sales)
	// iterate sales and collect titles to slice

	return c.Status(http.StatusOK).JSON(
		responses.SalesResponse{Status: http.StatusOK, Message: "success", Titles: title_array, Stocks: stocks_array, Sales: sales_array},
	)
}
