package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a todo as done",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatalf("todos done %s: %v\n", args[0], err)
		}
		for _, todo := range todos {
			if todo.ID == args[0] {
				todo.Done = true
			}
		}
		if err := writeToFile(todos, filename); err != nil {
			log.Fatalf("todos done %s: %v\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
