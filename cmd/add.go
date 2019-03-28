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
		todos, err := readFromFile(filepath)
		if err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
		t, err := NewTodo(args[0])
		if err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
		todos = append(todos, t)
		if err := writeToFile(todos, filepath); err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
