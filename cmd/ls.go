package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all todo items.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		printTodos(todos)
	},
}
