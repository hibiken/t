package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const filename = "todos.json"

// Todo is a task to be done.
type Todo struct {
	ID    string
	Title string
	Done  bool
}

func (t *Todo) String() string {
	if t.Done {
		return fmt.Sprintf("%s \u2714\t%s", t.ID, t.Title)
	}
	return fmt.Sprintf("%s _\t%s", t.ID, t.Title)
}

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todo is a CLI application to track your daily todos",
	Long:  "A simple todo application in command line.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Do not run this pre-run for the following commands.
		var whitelist = []string{"init", "help"}
		for _, c := range whitelist {
			if cmd.Name() == c {
				return
			}
		}
		// check if a file exist, if not user hasn't run `todos init`.
		if _, err := os.Stat(filename); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Error: Todos not initialized")
				fmt.Println("Run 'todos init' at the root of your project.")
				os.Exit(1)
			}
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello my todo CLI!")
	},
}

// Execute runs the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("rootCmd failed: %v\n", err)
		os.Exit(1)
	}
}
