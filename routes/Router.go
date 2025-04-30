package routes

import (
	"Todo-list/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api")

	// @/api/auth
	auth := api.Group("/auth")
	AuthRoutes(auth)

	// @/api/todoList
	todoList := api.Group("/todoList")
	todoList.Use(middleware.JWTMiddleware) // Tüm /api/todoList/* endpoint'lerinde JWT token doğrulaması yapılır
	TodoListRoutes(todoList)
}
