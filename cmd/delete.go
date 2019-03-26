package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatalf("todos delete %s", args[0], err)
		}
		idx := -1
		for i, todo := range todos {
			if todo.ID == args[0] {
				idx = i
			}
		}
		if idx != -1 {
			todos = append(todos[:idx], todos[idx+1:]...)
		}
		if err := writeToFile(todos, filename); err != nil {
			log.Fatalf("todos delete %s: %v\n", args[0], err)
		}
	},
}
