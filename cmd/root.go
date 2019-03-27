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

func init() {
	// check if the file exist.
	if _, err := os.Stat(filename); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
		// create a file and write an empty list.
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("[]"); err != nil {
			log.Fatal(err)
		}
		if err := f.Sync(); err != nil {
			log.Fatal(err)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todo is a CLI application to track your daily todos",
	Long:  "A simple todo application in command line.",
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
