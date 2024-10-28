package middleware

import (
	"net/http"
	"strings"
	"user_service/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("chama123") // Use the same secret key used for signing

// AuthMiddleware checks for a valid JWT token
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the Authorization header
	tokenString := c.Get("Authorization")

	// Check if the token is present
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid token",
		})
	}

	// Extract the token (bearer token format: "Bearer <token>")
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token format",
		})
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token's signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(http.StatusUnauthorized, "Invalid signing method")
		}
		return jwtSecret, nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// ---------------------------------------------------

	// Extract claims and validate user_id
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenUserName := claims["user_name"].(string) // Extract user_id from the token

		// Parse user_id from the request payload
		// var payload struct {
		// 	UserID string `json:"user_id"`
		// }

		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request payload",
			})
		}

		// Compare the user_id from the token and the payload
		if tokenUserName != user.Username {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "User Name mismatch",
			})
		}
	}

	// ---------------------------------------------------

	// If valid, proceed to the next handler
	return c.Next()
}
