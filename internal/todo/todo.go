package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID int
	Title string
	Completed bool
	CreatedAt time.Time
}

func NewTodo(id int, title string) *Todo {
	return &Todo{
		ID: id,
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func (t *Todo) Complete() {
	t.Completed = true
}

func (t Todo) String() string {
	status := "[ ]"
	if t.Completed {
		status = "[X]"
	}
	return fmt.Sprintf("%d. %s %s", t.ID, status, t.Title)
}

type TodoManager struct {
	todos []Todo
	nextID int
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		todos: make([]Todo, 0),
		nextID: 1,
	}
}

