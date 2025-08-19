package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

func NewTodo(id int, title string) *Todo {
	return &Todo{
		ID:        id,
		Title:     title,
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
	todos  []Todo
	nextID int
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

func (tm *TodoManager) AddTodo(title string) *Todo {
	todo := NewTodo(tm.nextID, title)
	tm.todos = append(tm.todos, *todo)
	tm.nextID++
	// Return pointer to the element in the slice, not the original
	return &tm.todos[len(tm.todos)-1]
}

func (tm *TodoManager) ListTodos() []Todo {
	return tm.todos
}

func (tm *TodoManager) GetTodoByID(id int) (*Todo, error) {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			return &tm.todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo with ID %d not found", id)
}

func (tm *TodoManager) CompleteTodo(id int) error {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos[i].Complete()
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

func (tm *TodoManager) RemoveTodo(id int) error {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			// Safer slice removal
			copy(tm.todos[i:], tm.todos[i+1:])
			tm.todos[len(tm.todos)-1] = Todo{} // Clear the last element
			tm.todos = tm.todos[:len(tm.todos)-1]
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

func (tm *TodoManager) UpdateTodo(id int, newTitle string) error {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos[i].Title = newTitle
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

func (tm *TodoManager) GetCompletedTodos() []Todo {
	var completed []Todo
	for _, todo := range tm.todos {
		if todo.Completed {
			completed = append(completed, todo)
		}
	}
	return completed
}

func (tm *TodoManager) GetPendingTodos() []Todo {
	var pending []Todo
	for _, todo := range tm.todos {
		if !todo.Completed {
			pending = append(pending, todo)
		}
	}
	return pending
}

func (tm *TodoManager) ClearCompleted() int {
	var remaining []Todo
	removedCount := 0

	for _, todo := range tm.todos {
		if todo.Completed {
			removedCount++
		} else {
			remaining = append(remaining, todo)
		}
	}

	tm.todos = remaining
	return removedCount
}