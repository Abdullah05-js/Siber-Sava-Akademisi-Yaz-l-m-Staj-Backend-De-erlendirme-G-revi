package models

import (
	"time"
)

type TodoStep struct {
	Id         string     `json:"id"`
	TodoListID string     `json:"todolistid"` // Hangi listeye aitse o listeyi işaret eder
	Content    string     `json:"content"`
	IsComplete bool       `json:"iscomplete"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  time.Time  `json:"updatedat"`
	DeletedAt  *time.Time `json:"deletedat"` // time.Time her zaman bir değer ister, nil olabilmesi için pointer yapıldı
}

// Mock todoSteps table
var todoSteps = map[string]TodoStep{}

func CreateTodoListStep(Id string, TodoListID string, Content string, IsComplete bool, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt *time.Time) {
	todoSteps[Id] = TodoStep{Id: Id, TodoListID: TodoListID, Content: Content, IsComplete: IsComplete, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, DeletedAt: DeletedAt}
	UpdateTodoListCompletion(TodoListID)
}

func FindStepAndDeleteById(Id string, userId string) bool {
	step, ok := todoSteps[Id]
	if !ok || userId != todoLists[step.TodoListID].UserID {
		return false
	} else if step.DeletedAt != nil {
		return true
	}
	now := time.Now()
	step.DeletedAt = &now
	step.UpdatedAt = now
	todoSteps[Id] = step
	UpdateTodoListCompletion(todoSteps[Id].TodoListID) // değer(kopya) olarak gönder
	return true
}

func UpdateTodoListStep(Id string, content string, iscomplete bool, userId string) bool {
	step, ok := todoSteps[Id]

	if !ok || step.DeletedAt != nil || userId != todoLists[step.TodoListID].UserID {
		return false
	}

	if step.IsComplete != iscomplete {
		step.IsComplete = iscomplete
	}

	if content != "" {
		step.Content = content
	}

	todoList := todoLists[step.TodoListID]

	now := time.Now()
	todoList.UpdatedAt = now
	step.UpdatedAt = now

	todoLists[step.TodoListID] = todoList
	todoSteps[Id] = step
	UpdateTodoListCompletion(todoSteps[Id].TodoListID)
	return true
}
