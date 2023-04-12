package api

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Health endpoint call
func Health(c *fiber.Ctx) error {
	// Health endpoint call
	log.Info("checking health endpoint")

	return c.SendString("All is well 👋!")
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Fiber instance
	app := fiber.New()

	// hello Handler
	app.Get("/health", Health)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Return the configured app
	return app
}
