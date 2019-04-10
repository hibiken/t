package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undoneCmd)
}

var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark todos as undone",
	Run: func(_ *cobra.Command, _ []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}

		filtered := filter(todos, func(t *Todo) bool { return t.Done })
		if len(filtered) == 0 {
			fmt.Println("You have no completed todos!")
			return
		}

		sortTodos(filtered)
		prompt := promptui.Select{
			Label:    "Select todo to mark as undone",
			Items:    titles(filtered),
			HideHelp: true,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		for _, t := range todos {
			if t.Title == selected {
				t.Done = false
			}
		}
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}
