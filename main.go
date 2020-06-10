package main

import (
	"minimalgo/api"
	"minimalgo/utils"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func init() {
	utils.LogSettting()
}

func main() {
	// Fiber instance
	app := fiber.New()

	// hello Handler
	app.Get("/", api.Hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// Start server
	log.Fatal(app.Listen(5430))
}
