package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

// path to the file the program reads from and writes to.
// The value gets assinged inside the pre-run function before all commands.
var filepath string

// allFlag used for this root command.
var allFlag bool

var rootCmd = &cobra.Command{
	Use:   "t",
	Short: "t is a CLI application to track your daily todos",
	Long:  "A simple todo application in command line.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fatal: cannot get current user: %v\n", err)
			os.Exit(1)
		}
		filepath = path.Join(usr.HomeDir, "todos.json")

		// check if the file exist
		if _, err := os.Stat(filepath); err == nil {
			return // file already exist, go ahead and execute Run function for the command
		} else if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "fatal: cannot get fileinfo : %v\n", err)
			os.Exit(1)
		}

		f, err := os.Create(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fatal: cannot create file : %v\n", err)
			os.Exit(1)
		}
		if _, err := f.WriteString("[]"); err != nil {
			fmt.Fprintf(os.Stderr, "fatal: cannot write to %s : %v\n", filepath, err)
			os.Exit(1)
		}
		if err := f.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "fatal: cannot write to %s : %v\n", filepath, err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := readFromFile(filepath)
		if err != nil {
			log.Fatalf("todos ls: %v\n", err)
		}
		printTodos(todos, allFlag)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "List all todos including completed ones")
}

// Execute runs the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("fatal: %v\n", err)
		os.Exit(1)
	}
}
