package cmd

import (
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Deletes only done todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		var res []*Todo
		for _, t := range todos {
			if !t.Done {
				res = append(res, t)
			}
		}
		if err := writeToFile(res, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}
