package api

import (
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Health endpoint call
func Health(c *fiber.Ctx) {
	// Health endpoint call
	log.Info("checking health endpoint")
	c.Send("All is well 👋!")
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Fiber instance
	app := fiber.New()

	// hello Handler
	app.Get("/health", Health)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// Return the configured app
	return app
}
