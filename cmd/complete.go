package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [todo-id]",
	Short: "Mark a todo as completed",
	Long:  `Mark a todo as completed by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid todo ID\n", args[0])
			return
		}

		err = todoManager.CompleteTodo(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Completed todo %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
