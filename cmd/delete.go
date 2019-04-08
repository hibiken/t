package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a todo",
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		filtered := filter(todos, func(t *Todo) bool { return !t.Done })
		sortTodos(filtered)
		prompt := promptui.Select{
			Label:    "Select todo to delete",
			Items:    titles(filtered),
			HideHelp: true,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		res := []*Todo{}
		for _, t := range todos {
			if t.Title != selected {
				res = append(res, t)
			}
		}
		if err := writeToFile(res, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
