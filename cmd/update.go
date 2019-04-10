package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"modify"},
	Short:   "Updates a todo",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		filtered := filter(todos, func(t *Todo) bool { return !t.Done })
		if len(filtered) == 0 {
			fmt.Println("You have no todos")
			return
		}

		sortTodos(filtered)
		prompt := promptui.Select{
			Label:    "Select todo to update",
			Items:    titles(filtered),
			HideHelp: true,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		t := findByTitle(todos, selected)
		if t == nil {
			panic("t should not be nil")
		}

		titles := make(map[string]bool) // set of titles
		for _, t := range todos {
			titles[t.Title] = true
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}

		titlePrompt := promptui.Prompt{
			Label:     "Enter your updated todo name",
			Templates: templates,
			Validate: func(input string) error {
				if len(strings.TrimSpace(input)) == 0 {
					return errors.New("empty string")
				}
				if titles[input] {
					return errors.New("already exist")
				}
				return nil
			},
		}

		priorityPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("Enter priority [currently = %d]", t.Priority),
			Templates: templates,
			Validate: func(input string) error {
				if _, err := strconv.Atoi(input); err != nil {
					return err
				}
				return nil
			},
		}

		title, err := titlePrompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}

		priority, err := priorityPrompt.Run()
		if err != nil {
			printErrorAndExit(err)
		}
		p, err := strconv.Atoi(priority)
		if err != nil {
			printErrorAndExit(err)
		}
		t.Title = title
		t.Priority = p
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
