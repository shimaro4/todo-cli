package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [todo-id] [new-title]",
	Short: "Update a todo's title",
	Long:  `Update a todo's title by providing its ID and the new title.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid todo ID\n", args[0])
			return
		}

		newTitle := strings.Join(args[1:], " ")
		err = todoManager.UpdateTodo(id, newTitle)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Updated todo %d to: %s\n", id, newTitle)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
