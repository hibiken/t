package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var clearallCmd = &cobra.Command{
	Use:   "clearall",
	Short: "Deletes all todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if err := writeToFile([]*Todo{}, filepath); err != nil {
			log.Fatalf("todos clearall: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearallCmd)
}
