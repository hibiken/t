package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Deletes only done todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filename)
		if err != nil {
			log.Fatalf("todos prune: %v\n", err)
		}
		var res []*Todo
		for _, t := range todos {
			if !t.Done {
				res = append(res, t)
			}
		}
		if err := writeToFile(res, filename); err != nil {
			log.Fatalf("todos prune: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}
