package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiter applies a rate limit based on the UserName passed in the request header.
func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Extract UserName from the request header
		userName := c.Get("userName")

		// Check if UserName is provided
		if userName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Missing username in headers",
			})
		}

		// Set up the rate limiter
		return limiter.New(limiter.Config{
			Max:        50,               // Max requests
			Expiration: 30 * time.Second, // Time frame for limit

			// Note: if Key is removed, the rate limiter will default to monitoring requests by IP address.
			Key: func(c *fiber.Ctx) string { // Use UserName as key
				return userName
			},
		})(c) // Call the limiter handler
	}
}
