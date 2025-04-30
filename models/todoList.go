package models

import (
	"time"
)

type ToDoList struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  time.Time  `json:"updatedat"`
	DeletedAt  *time.Time `json:"deletedat"`
	Completion float64    `json:"completion"`
	UserID     string     `json:"userid"` // Bu listeyi oluşturan kullanıcının ID'si (jwt tokendan alınır)
}

//Mock ToDoLists table
var ToDoLists = []ToDoList{}
