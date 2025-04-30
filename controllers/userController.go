package controllers

import (
	"github.com/gofiber/fiber/v2"
	"Todo-list/models"
)

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(models.Users)
}

