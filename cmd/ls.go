package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all todos including completed ones")
}

var allFlag bool

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all todo items.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatalf("todo ls: %v\n", err)
		}
		printTodos(todos, allFlag)
	},
}
