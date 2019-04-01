package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			idPrompt           = "Enter todo ID to update: "
			titlePromptTmpl    = "Enter title (current: %q): "
			priorityPromptTmpl = "Enter priority (current: %d): "
		)
		todos, err := readFromFile(filepath)
		if err != nil {
			printErrorAndExit(err)
		}
		printTodos(todos, false) // No need to show done todos
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(idPrompt)
		var t *Todo
		for scanner.Scan() {
			id := scanner.Text()
			t = findByID(todos, id)
			if t != nil {
				break
			}
			fmt.Printf("todo with ID %q not found. Try again.\n\n", id)
			fmt.Print(idPrompt)
		}
		fmt.Printf(titlePromptTmpl, t.Title)
		for scanner.Scan() {
			title := scanner.Text()
			if title == "" {
				fmt.Println("Skip updating title")
			} else {
				t.Title = title
			}
			break
		}
		fmt.Printf(priorityPromptTmpl, t.Priority)
		for scanner.Scan() {
			pStr := scanner.Text()
			p, err := strconv.Atoi(pStr)
			if err != nil {
				fmt.Printf("Could not parse %q. Please try again.\n\n", pStr)
				fmt.Printf(priorityPromptTmpl, t.Priority)
				continue
			}
			t.Priority = p
			break
		}
		if err := writeToFile(todos, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
