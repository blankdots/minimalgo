package api

import (
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Handler
func Hello(c *fiber.Ctx) {
	log.Info("something")
	c.Send("Hello, World 👋!")
}
