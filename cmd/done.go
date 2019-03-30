package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [ids]",
	Short: "Mark todos as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, ids []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			fmt.Printf("fatal: todos done %s: %v\n", strings.Join(ids, " "), err)
		}
		for _, t := range todos {
			if contains(ids, t.ID) {
				t.Done = true
			}
		}
		if err := writeToFile(todos, filepath); err != nil {
			fmt.Printf("fatal: todos done %s: %v\n", strings.Join(ids, " "), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
