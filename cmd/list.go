package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `Display all todos with their IDs, completion status, and titles.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos := todoManager.ListTodos()
		if len(todos) == 0 {
			fmt.Println("No todos found. Add some with: todo-cli add \"your todo\"")
			return
		}
		
		fmt.Println("Your todos:")
		for _, todo := range todos {
			fmt.Println(todo.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}