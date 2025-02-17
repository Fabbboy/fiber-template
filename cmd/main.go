package main

import (
	"github.com/gofiber/fiber/v2"
	"schaub-dev.xyz/fabrice/fiber-template/pkg"
	"schaub-dev.xyz/fabrice/fiber-template/pkg/database"
	"schaub-dev.xyz/fabrice/fiber-template/pkg/middleware"
)

func main() {
	cfg := pkg.NewConfig()

	app := fiber.New()
	db, err := database.NewDbClient(cfg, app)
	if err != nil {
		panic(err)
	}

	app.Use(middleware.ReqLog())
	app.Use(middleware.InjectItem("config", cfg))
	app.Use(middleware.InjectItem("db", db))

	app.Listen(cfg.Host)
}
