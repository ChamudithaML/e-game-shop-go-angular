package services

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// Serve images from specific directory
func ServeImage(c *fiber.Ctx, fileName string) error {
	filePath := filepath.Join("./public", fileName)
	return c.SendFile(filePath, true)
}
