package main

import (
	"fmt"
	"github.com/shimaro4/todo-cli/internal/todo"
)

func main() {
	// Create manager
	manager := todo.NewTodoManager()
	
	// Add some todos
	fmt.Println("=== Adding Todos ===")
	todo1 := manager.AddTodo("Learn Go interfaces")
	todo2 := manager.AddTodo("Build CLI application") 
	todo3 := manager.AddTodo("Deploy to production")
	
	fmt.Printf("Added: %s\n", todo1.String())
	fmt.Printf("Added: %s\n", todo2.String())
	fmt.Printf("Added: %s\n", todo3.String())
	
	// List all todos
	fmt.Println("\n=== Listing All Todos ===")
	todos := manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}
	
	// Find specific todo
	fmt.Println("\n=== Finding Specific Todo ===")
	found, err := manager.GetTodoByID(2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found: %s\n", found.String())
	}
	
	// Try to find non-existent todo
	_, err = manager.GetTodoByID(99)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	fmt.Println("\n✅ All TodoManager methods working!")
}