package main

import (
	"fmt"

	"github.com/shimaro4/todo-cli/internal/todo"
)

func main() {
	manager := todo.NewTodoManager()

	fmt.Println("=== Testing Complete CLI Simulation ===")

	fmt.Println("\n1. Adding todos...")
	todo1 := manager.AddTodo("Learn Go interfaces")
	todo2 := manager.AddTodo("Build CLI application")
	todo3 := manager.AddTodo("Deploy to production")
	fmt.Printf("Added: %s\n", todo1.String())
	fmt.Printf("Added: %s\n", todo2.String())
	fmt.Printf("Added: %s\n", todo3.String())

	fmt.Println("\n2. Listing all todos...")
	todos := manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n3. Completing todo 2...")
	err := manager.CompleteTodo(2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Completed todo 2")
	}

	fmt.Println("\n4. Updating todo 3...")
	err = manager.UpdateTodo(3, "Deploy to production (updated)")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Updated todo 3")
	}

	fmt.Println("\n5. Listing after changes...")
	todos = manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n6. Removing todo 1...")
	err = manager.RemoveTodo(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Removed todo 1")
	}

	fmt.Println("\n7. Final list...")
	todos = manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n CLI simulation complete")
	fmt.Println("\nYour CLI commands are ready:")
	fmt.Println("  todo-cli add \"task\"")
	fmt.Println("  todo-cli list")
	fmt.Println("  todo-cli complete [id]")
	fmt.Println("  todo-cli remove [id]")
	fmt.Println("  todo-cli update [id] \"new title\"")
}
