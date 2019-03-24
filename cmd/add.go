package cmd

import (
	"log"

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
		todo := Todo{Title: args[0]}
		if err := addTodo(todo, filename); err != nil {
			log.Fatal(err)
		}
	},
}
