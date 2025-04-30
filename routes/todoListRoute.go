package routes

import (
	"Todo-list/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoListRoutes(app fiber.Router) {
	app.Post("/login", controllers.GetUsers)
}