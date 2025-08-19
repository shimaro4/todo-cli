package cmd

import (
	"os"

	"github.com/shimaro4/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var todoManager *todo.TodoManager

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A production-quality CLI todo manager",
	Long: `A professional todo list manager built with Go and Cobra.

Manage your tasks efficiently from the command line with features like:
  • Adding new todos
  • Listing all todos  
  • Completing todos
  • Deleting todos

Examples:
  todo-cli add "Learn Go interfaces"
  todo-cli list
  todo-cli complete 1
  todo-cli delete 2`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	todoManager = todo.NewTodoManager()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
