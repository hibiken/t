package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [string of task]",
	Short: "Adds a new todo item.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adding a new todo...")
		todos = append(todos, Todo{Title: args[0]})
		for i := 0; i < len(todos); i++ {
			fmt.Printf("%d: %s\n", i+1, todos[i].String())
		}
	},
}
