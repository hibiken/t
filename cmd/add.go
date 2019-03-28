package cmd

import (
	"log"
	"time"

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
		todos, err := readFromFile(filepath)
		if err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
		todos = append(todos, &Todo{ID: id, Title: args[0], CreatedAt: time.Now()})
		if err := writeToFile(todos, filepath); err != nil {
			log.Fatalf("todos add %s: %v\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
