package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undoneCmd)
}

var undoneCmd = &cobra.Command{
	Use:   "undone [id]",
	Short: "Mark a todo as undone",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatalf("todos undone %s: %v\n", args[0], err)
		}
		for _, todo := range todos {
			if todo.ID == args[0] {
				todo.Done = false
			}
		}
		if err := writeToFile(todos, filename); err != nil {
			log.Fatalf("todos undone %s: %v\n", args[0], err)
		}
	},
}
