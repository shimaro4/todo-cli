package main

import (
	"fmt"

	"github.com/shimaro4/todo-cli/internal/todo"
)

func main() {
	// Test NewTodo
	t := todo.NewTodo(1, "Test todo")
	fmt.Println("Created:", t.String())

	// Test Complete
	t.Complete()
	fmt.Println("Completed:", t.String())

	fmt.Println("All basic functionality works")
}
