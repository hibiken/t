package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ids]",
	Short: "Deletes a todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, ids []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			fmt.Printf("fatal: todos delete %s: %v\n", strings.Join(ids, " "), err)
		}

		res := []*Todo{}
		for _, t := range todos {
			if !contains(ids, t.ID) {
				res = append(res, t)
			}
		}
		if err := writeToFile(res, filepath); err != nil {
			fmt.Printf("fatal: todos delete %s: %v\n", strings.Join(ids, " "), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
