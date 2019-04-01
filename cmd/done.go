package cmd

import (
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [ids]",
	Short: "Mark todos as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, ids []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		for _, t := range todos {
			if contains(ids, t.ID) {
				t.Done = true
			}
		}
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
