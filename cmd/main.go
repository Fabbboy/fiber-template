package main

import (
	"github.com/gofiber/fiber/v2"
	"schaub-dev.xyz/fabrice/fiber-template/pkg"
	"schaub-dev.xyz/fabrice/fiber-template/pkg/database"
	"schaub-dev.xyz/fabrice/fiber-template/pkg/middleware"
)

func main() {
	cfg := pkg.NewConfig()
	app_logger := pkg.NewLogger("App", cfg.LogLevel)

	app := fiber.New()
	db, err := database.NewDbClient(cfg, app)
	if err != nil {
		panic(err)
	}

	app.Use(middleware.InjectItem("app_logger", app_logger))
	app.Use(middleware.InjectItem("config", cfg))
	app.Use(middleware.ReqLog())
	app.Use(middleware.InjectItem("db", db))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the homepage!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Oops! That page doesn't exist. Try again!")
	})

	app_logger.Info("Starting server on %s", cfg.Host)
	app.Listen(cfg.Host)
}
