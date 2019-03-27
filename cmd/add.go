package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [string of todo]",
	Short: "Adds a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := genID()
		if err != nil {
			log.Fatalf("todos add %s: failed to generate ID: %v\n", args[0], err)
		}
		todo := &Todo{ID: id, Title: args[0]}
		if err := addTodo(todo, filename); err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
