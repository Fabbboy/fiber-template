package middleware

import "github.com/gofiber/fiber/v2"

func InjectItem[T any](key string, val T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(key, val)
		return c.Next()
	}
}
