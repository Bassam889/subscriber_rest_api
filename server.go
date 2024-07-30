package main

import (
	"app/initializers"
	"app/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	routes.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	app.Listen(":3000")
}
