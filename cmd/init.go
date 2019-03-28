package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes todos for the current directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
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

		// if gitignore file exists, append filename to it
		f, err = os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return
		}
		defer f.Close()
		if _, err := f.WriteString(filename); err != nil {
			// write failed, prompt the user to upate .gitignore file manually
			fmt.Printf("Add %s to .gitignore file", filename)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
