package models

import (
	"time"
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
var todoLists = []TodoList{}

func CreateTodoList(Id string, Name string, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt *time.Time, Completion float64, UserID string) {
	todoLists = append(todoLists, TodoList{Id: Id, Name: Name, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, DeletedAt: DeletedAt, Completion: Completion, UserID: UserID})
}

func FindListAndDeleteById(Id string, userId string) bool {
	for index, todoList := range todoLists {
		if todoList.Id == Id && userId == todoList.UserID {
			now := time.Now()
			todoLists[index].DeletedAt = &now
			todoSteps[index].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

func GetUserByTodoListId(Id string) string {
	for _, todoList := range todoLists {
		if todoList.Id == Id {
			return todoList.UserID
		}
	}
	return ""
}

func UpdateTodoListCompletion(Id string) {
	for index, todoList := range todoLists {
		if todoList.Id == Id && todoList.DeletedAt == nil {
			countSteps := 1
			countStepsComplete := 0
			for _, todoListStep := range todoSteps {
				if todoListStep.TodoListID == todoList.Id {
					if todoListStep.IsComplete {
						countStepsComplete++
					}
					countSteps++
				}
			}
			todoLists[index].Completion = float64(countStepsComplete/countSteps) * 100
			todoLists[index].UpdatedAt = time.Now()
			break
		}
	}
}
