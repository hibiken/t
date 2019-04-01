package cmd

import (
	"github.com/spf13/cobra"
)

var priorityFlag int

var addCmd = &cobra.Command{
	Use:   "add [string of todo]",
	Short: "Adds a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		t, err := NewTodo(args[0], priorityFlag)
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
	addCmd.Flags().IntVarP(&priorityFlag, "priority", "p", 2, "Specify the priority of todo")
}
