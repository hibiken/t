package cmd

import (
	"fmt"

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
		fmt.Println("list all todos...")
		for i := 0; i < len(todos); i++ {
			fmt.Printf("%d: %s\n", i+1, todos[i].String())
		}
	},
}
