package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undoneCmd)
}

var undoneCmd = &cobra.Command{
	Use:   "undone [ids]",
	Short: "Mark todos as undone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, ids []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			fmt.Printf("fatal: todos undone %s: %v\n", strings.Join(ids, " "), err)
		}
		for _, t := range todos {
			if contains(ids, t.ID) {
				t.Done = false
			}
		}
		if err := writeToFile(todos, filepath); err != nil {
			fmt.Printf("fatal: todos undone %s: %v\n", strings.Join(ids, " "), err)
		}
	},
}
