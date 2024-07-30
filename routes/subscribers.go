package routes

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/subscribers")
	v1.Get("/", controllers.GetSubscriber)
	v1.Get("/:id", controllers.GetSubscriberById)
	v1.Post("/", controllers.CreateSubscriber)
	v1.Put("/:id", controllers.UpdateSubscriber)
	v1.Delete("/:id", controllers.DeleteSubscriber)
}
