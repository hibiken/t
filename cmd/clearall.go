package cmd

import (
	"github.com/spf13/cobra"
)

var clearallCmd = &cobra.Command{
	Use:   "clearall",
	Short: "Deletes all todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if err := writeToFile([]*Todo{}, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearallCmd)
}
