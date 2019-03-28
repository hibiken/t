package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var allFlag bool

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			log.Fatalf("todos ls: %v\n", err)
		}
		printTodos(todos, allFlag)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all todos including completed ones")
}
