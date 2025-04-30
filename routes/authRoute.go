package routes

import (
	"Todo-list/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	app.Post("/login", controllers.GetUsers)
}
