package controllers

import (
	"github.com/gofiber/fiber/v2"

	"file_server_service/models"
	"file_server_service/services"
	"log"
)

// Serve the file
func ServeFile(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	filePath := "./public/docs/" + fileName

	if err := c.SendFile(filePath); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}
	return nil
}

func CreateDocument(c *fiber.Ctx) error {
	var letter models.LetterDocument

	// Parse the JSON body into the LetterDocument struct
	if err := c.BodyParser(&letter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate the struct using validate service
	if err := services.ValidateLetterDocument(letter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate the document
	fileName, err := services.GenerateDocx(letter)
	if err != nil {
		log.Println("Error generating document:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate document",
		})
	}

	// Construct the file URL
	fileURL := c.BaseURL() + "/docs/" + fileName

	// Return the file URL
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Document created successfully",
		"file_url": fileURL,
	})
}
