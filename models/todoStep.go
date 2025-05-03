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
var todoSteps = []TodoStep{}

func CreateTodoListStep(Id string, TodoListID string, Content string, IsComplete bool, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt *time.Time) {
	todoSteps = append(todoSteps, TodoStep{Id: Id, TodoListID: TodoListID, Content: Content, IsComplete: IsComplete, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, DeletedAt: DeletedAt})
}

func FindStepAndDeleteById(Id string, userId string) bool {
	for index, todoListStep := range todoSteps {
		if todoListStep.Id == Id && userId == GetUserByTodoListId(todoListStep.TodoListID) {
			now := time.Now()
			todoSteps[index].DeletedAt = &now
			todoSteps[index].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

func UpdateTodoListStep(Id string, content string, iscomplete bool, userId string) bool {
	for index, todoListStep := range todoSteps {
		if todoListStep.Id == Id && userId == GetUserByTodoListId(todoListStep.TodoListID) && todoListStep.DeletedAt == nil {
			if todoListStep.IsComplete != iscomplete {
				todoSteps[index].IsComplete = iscomplete
				UpdateTodoListCompletion(todoListStep.TodoListID)
			}
			if content != "" {
				todoSteps[index].Content = content
			}
			todoSteps[index].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}
