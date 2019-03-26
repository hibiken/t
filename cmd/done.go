package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task to be done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
