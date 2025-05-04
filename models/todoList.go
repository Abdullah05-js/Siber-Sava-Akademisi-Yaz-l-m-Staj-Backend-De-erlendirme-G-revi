package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type TodoList struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  time.Time  `json:"updatedat"`
	DeletedAt  *time.Time `json:"deletedat"`
	Completion float64    `json:"completion"`
	UserID     string     `json:"userid"` // Bu listeyi oluşturan kullanıcının ID'si (jwt tokendan alınır)
}

// Mock todoList table
var todoLists = map[string]TodoList{}

func CreateTodoList(Id string, Name string, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt *time.Time, Completion float64, UserID string) {
	todoLists[Id] = TodoList{Id: Id, Name: Name, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, DeletedAt: DeletedAt, Completion: Completion, UserID: UserID}
}

func FindListAndDeleteById(Id string, userId string) bool {
	todoList, ok := todoLists[Id]
	if !ok || userId != todoList.UserID {
		return false
	}

	for _, todoStep := range todoSteps {
		if todoStep.TodoListID == Id {
			FindStepAndDeleteById(todoStep.Id, userId)
		}
	}

	now := time.Now()
	todoList.DeletedAt = &now
	todoList.UpdatedAt = now
	todoLists[Id] = todoList

	return true
}

func GetUserByTodoListId(Id string) string {
	todoList, ok := todoLists[Id]
	if !ok {
		return ""
	}
	return todoList.UserID
}

func UpdateTodoListCompletion(Id string) {
	for index, todoList := range todoLists {
		if todoList.Id == Id && todoList.DeletedAt == nil {
			countSteps := 0
			countStepsComplete := 0
			for _, todoListStep := range todoSteps {
				if todoListStep.TodoListID == todoList.Id {
					countSteps++
					if todoListStep.IsComplete {
						countStepsComplete++
					}
				}
			}
			if countSteps > 0 {
				newTodoList := todoLists[index]
				newTodoList.Completion = float64(countStepsComplete) / float64(countSteps) * 100
				newTodoList.UpdatedAt = time.Now()
				todoLists[index] = newTodoList
			}
			break
		}
	}
}

func GetTodoListsByUserId(userId string, isAdmin bool) []fiber.Map {
	var data []fiber.Map

	for _, todoList := range todoLists {
		if (todoList.UserID == userId || isAdmin) && todoList.DeletedAt == nil {

			StepsArr := []TodoStep{}
			for _, todoStep := range todoSteps {
				if todoStep.TodoListID == todoList.Id && todoStep.DeletedAt == nil {
					StepsArr = append(StepsArr, todoStep)
				}
			}

			data = append(data, fiber.Map{
				"todolist":  todoList,
				"todosteps": StepsArr,
			})
		}
	}

	return data
}

func IsTodoListExistById(Id string) bool {
	_, ok := todoLists[Id]

	if !ok {
		return false
	}

	return true
}
