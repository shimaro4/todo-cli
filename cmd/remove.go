package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [todo-id]",
	Short: "Remove a todo",
	Long:  `Remove a todo by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid todo ID\n", args[0])
			return
		}

		err = todoManager.RemoveTodo(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Removed todo %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
