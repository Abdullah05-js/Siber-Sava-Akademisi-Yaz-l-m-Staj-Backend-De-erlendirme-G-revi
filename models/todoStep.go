package models

import (
	"time"
)

type ToDoStep struct {
	ID         string     `json:"id"`
	ToDoListID string     `json:"todolistid"` // Hangi listeye aitse o listeyi işaret eder
	Content    string     `json:"content"`
	IsDone     bool       `json:"isdone"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  time.Time  `json:"updatedat"`
	DeletedAt  *time.Time `json:"deletedat"` // time.Time her zaman bir değer ister, nil olabilmesi için pointer yapıldı
}

var ToDoSteps = []ToDoStep{}
