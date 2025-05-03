package routes

import (
	"Todo-list/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoListRoutes(app fiber.Router) {

	// Kullanıcının tüm yapılacaklar listelerini (ToDo listelerini) getirmek için
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	// Yeni bir yapılacaklar listesi oluşturmak için
	app.Post("/create", controllers.CreateTodoList)

	// Belirli bir yapılacaklar listesini silmek için
	app.Delete("/:todoListId", controllers.DeleteTodoList)

	// Belirli bir yapılacaklar listesine yeni bir madde (adım) eklemek için
	app.Post("/:todoListId/steps/create", controllers.CreateTodoListStep)

	// Belirli bir yapılacaklar listesi maddesini (adımı) silmek için
	app.Delete("/steps/:stepId", controllers.DeleteTodoListStep)

	// Belirli bir yapılacaklar listesi maddesini (adımı) düzenlemek/güncellemek için
	app.Patch("/:todoListId/steps/:stepId", controllers.UpdateTodoListStep)
}
