package middleware

import (
	"github.com/gofiber/fiber/v2"
	"schaub-dev.xyz/fabrice/fiber-template/pkg"
)

func ReqLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := c.Locals("mw_logger").(*pkg.Logger)
		logger.Info("Request received: %s %s", c.Method(), c.Path())
		return c.Next()
	}
}
