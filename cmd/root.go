package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Todo is a task to be done.
type Todo struct {
	Title string
	Done  bool
}

func (t *Todo) String() string {
	if t.Done {
		return fmt.Sprintf("\u2714\t%s", t.Title)
	}
	return fmt.Sprintf("_\t%s", t.Title)
}

var todos []Todo

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todo is a CLI application to track your daily todos",
	Long:  "A simple todo application in command line.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello my todo CLI!")
	},
}

func init() {
	todos = []Todo{
		{"Go get groceries", false},
		{"Finish Go book ch5", false},
		{"Create my first CLI app with Cobra", true},
	}
}

// Execute runs the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("rootCmd failed: %v\n", err)
		os.Exit(1)
	}
}
