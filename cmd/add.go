package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long: `Add a new todo item to your list.

The todo text should be provided as an argument.

Examples:
  todo-cli add "Learn Go interfaces"
  todo-cli add "Complete Phase 1 project"
  todo-cli add "Review error handling patterns"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Please provide a todo item")
			return
		}
		todoText := strings.Join(args, " ")
		todo := todoManager.AddTodo(todoText)
		fmt.Printf("Added: %s\n", todo.String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
