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
