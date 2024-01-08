package routes

import (
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// GET /api/v1/
	// Health check
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
