package controllers

import (
	"Todo-list/models"
	"Todo-list/utils"
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	user := new(struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	})

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // kötü istek çünkü bodyinin formatı yanlış = kod 400
			"error": "body format is wrong",
		})
	}

	targetUser, found := models.GetUserByName(user.Name)

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ // kullanıcı adı veya şifre hatalı kod 401
			"error": "password or username is wrong",
		})
	}

	if targetUser.Pass != user.Pass {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ // kullanıcı adı veya şifre hatalı kod 401
			"error": "password or username is wrong",
		})
	}

	token, err := utils.GenerateToken(&targetUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ // dahili hatta kod 500
			"error": "error occurred while generating token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
