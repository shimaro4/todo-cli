package main

import (
	"fmt"

	"github.com/shimaro4/todo-cli/internal/todo"
)

func main() {
	manager := todo.NewTodoManager()

	fmt.Println("=== Adding Todos ===")
	todo1 := manager.AddTodo("Learn Go interfaces")
	todo2 := manager.AddTodo("Build CLI application")
	todo3 := manager.AddTodo("Deploy to production")

	fmt.Printf("Added: %s\n", todo1.String())
	fmt.Printf("Added: %s\n", todo2.String())
	fmt.Printf("Added: %s\n", todo3.String())

	fmt.Println("\n=== Listing All Todos ===")
	todos := manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n=== Finding Specific Todo ===")
	found, err := manager.GetTodoByID(2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found: %s\n", found.String())
	}

	_, err = manager.GetTodoByID(99)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	fmt.Println("\n=== Testing CompleteTodo ===")
	err = manager.CompleteTodo(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Successfully completed todo 1")
	}

	fmt.Println("\n=== Testing Pointer Fix ===")
	fmt.Printf("Before completing todo2: %s\n", todo2.String())
	todo2.Complete()
	fmt.Printf("After completing todo2: %s\n", todo2.String())

	fmt.Println("\n=== After Completions ===")
	todos = manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n=== Testing UpdateTodo ===")
	err = manager.UpdateTodo(3, "Deploy to production (updated)")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Successfully updated todo 3")
	}

	fmt.Println("\n=== Testing GetCompletedTodos ===")
	completed := manager.GetCompletedTodos()
	if len(completed) == 0 {
		fmt.Println("No completed todos")
	} else {
		for _, t := range completed {
			fmt.Println(t.String())
		}
	}

	fmt.Println("\n=== Testing GetPendingTodos ===")
	pending := manager.GetPendingTodos()
	if len(pending) == 0 {
		fmt.Println("No pending todos")
	} else {
		for _, t := range pending {
			fmt.Println(t.String())
		}
	}

	fmt.Println("\n=== Testing ClearCompleted ===")
	removedCount := manager.ClearCompleted()
	fmt.Printf("Removed %d completed todos\n", removedCount)

	fmt.Println("\n=== After ClearCompleted ===")
	todos = manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	fmt.Println("\n=== Testing RemoveTodo ===")
	fmt.Println("Before removal:")
	todos = manager.ListTodos()
	for _, t := range todos {
		fmt.Println(t.String())
	}

	if len(todos) > 0 {
		todoToRemove := todos[0].ID
		fmt.Printf("Attempting to remove todo ID: %d\n", todoToRemove)
		err = manager.RemoveTodo(todoToRemove)
		if err != nil {
			fmt.Printf("Error removing todo: %v\n", err)
		} else {
			fmt.Printf("Successfully removed todo %d\n", todoToRemove)
		}
	}

	fmt.Println("\nAfter removal:")
	todos = manager.ListTodos()
	if len(todos) == 0 {
		fmt.Println("No todos remaining")
	} else {
		for _, t := range todos {
			fmt.Println(t.String())
		}
	}

	fmt.Println("\n=== Testing Error Cases ===")
	err = manager.CompleteTodo(999)
	if err != nil {
		fmt.Printf("Expected error for CompleteTodo: %v\n", err)
	}

	err = manager.UpdateTodo(999, "Non-existent")
	if err != nil {
		fmt.Printf("Expected error for UpdateTodo: %v\n", err)
	}

	err = manager.RemoveTodo(999)
	if err != nil {
		fmt.Printf("Expected error for RemoveTodo: %v\n", err)
	}

	fmt.Println("\nAll TodoManager methods tested successfully")
}
