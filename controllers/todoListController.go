package controllers

import (
	"Todo-list/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTodoList(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)

	InputModel := new(struct {
		Name string `json:"name"`
	})

	err := c.BodyParser(InputModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body format is wrong",
		})
	}
	id := uuid.NewString()
	models.CreateTodoList(
		id,
		InputModel.Name,
		time.Now(),
		time.Now(),
		nil,
		0.0,
		userId,
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":     id,
		"status": "Created successfully",
	})
}

func DeleteTodoList(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)
	todoListId := c.Params("todoListId")

	if todoListId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "todoListId parameter is required",
		})
	}

	isSuccess := models.FindListAndDeleteById(todoListId, userId)

	if !isSuccess {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or already deleted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Deleted successfully",
	})
}

func CreateTodoListStep(c *fiber.Ctx) error {
	todoListId := c.Params("todoListId")

	if todoListId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "todoListId parameter is required",
		})
	}

	InputModel := new(struct {
		Content string `json:"content"`
	})

	err := c.BodyParser(InputModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body format is wrong",
		})
	}
	id := uuid.NewString()
	models.CreateTodoListStep(
		id,
		todoListId,
		InputModel.Content,
		false,
		time.Now(),
		time.Now(),
		nil,
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":     id,
		"status": "Created successfully",
	})
}

func DeleteTodoListStep(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)
	todoListStepId := c.Params("stepId")

	if todoListStepId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "todoListStepId parameter is required",
		})
	}

	isSuccess := models.FindStepAndDeleteById(todoListStepId, userId) // bunlara mesaj ekle yani isSuccess,message

	if !isSuccess {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or already deleted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Deleted successfully",
	})
}

func UpdateTodoListStep(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)
	todoListId := c.Params("todoListId")
	todoListStepId := c.Params("stepId")

	if todoListStepId == "" || todoListId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "parameters is required",
		})
	}

	InputModel := new(struct {
		Content    string `json:"content"`
		IsComplete bool   `json:"iscomplete"`
	})

	err := c.BodyParser(InputModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body format is wrong" + err.Error(),
		})
	}
	isSuccess := models.UpdateTodoListStep(todoListStepId, InputModel.Content, InputModel.IsComplete, userId)

	if !isSuccess {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or deleted",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Updated successfully",
	})
}

func GetTodoLists(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)

	user, ok := models.GetUserById(userId)

	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found or deleted",
		})
	}

	todoLists := models.GetTodoListsByUserId(userId, user.IsAdmin)

	return c.Status(fiber.StatusOK).JSON(todoLists)
}
