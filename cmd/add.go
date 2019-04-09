package cmd

import (
	"errors"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var priorityFlag int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
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
			Label:     "Enter your todo",
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
			Label:     "Enter priority [default = 2]",
			Templates: templates,
			Validate: func(input string) error {
				if input == "" {
					return nil
				}
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
		p := 2
		if priority != "" {
			p, _ = strconv.Atoi(priority) // ignoring error since it's already validated
		}

		t, err := NewTodo(title, p)
		if err != nil {
			printErrorAndExit(err)
		}
		todos = append(todos, t)
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
