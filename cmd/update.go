package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var interactiveMode bool

var updateCmd = &cobra.Command{
	Use:     "update [id] [updated title]",
	Aliases: []string{"modify"},
	Short:   "Updates a todo",
	Run: func(cmd *cobra.Command, args []string) {
		if interactiveMode {
			executeInteractiveMode(cmd, args)
		} else {
			execute(cmd, args)
		}
	},
}

// Example commands:
// t update 123 "Updated title"
// t update 123 "Updated title" -p 9
func execute(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Error: Wrong number of arguments")
		fmt.Println("Usage: t update [id] [updated title]")
		return
	}
	todos, err := readFromFile(filepath)
	if err != nil {
		printErrorAndExit(err)
	}
	t := findByID(todos, args[0])
	if t == nil {
		printErrorAndExit(fmt.Sprintf("cannot find todo with ID = %q", args[0]))
	}
	t.Title = args[1]
	if cmd.Flags().Changed("priority") {
		t.Priority = priorityFlag
	}
	if err := writeToFile(todos, filepath); err != nil {
		printErrorAndExit(err)
	}
}

func executeInteractiveMode(cmd *cobra.Command, args []string) {
	const (
		idPrompt           = "Enter todo ID to update: "
		titlePromptTmpl    = "Enter title (current: %q): "
		priorityPromptTmpl = "Enter priority (current: %d): "
	)
	todos, err := readFromFile(filepath)
	if err != nil {
		printErrorAndExit(err)
	}
	printTodos(todos, false)
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
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&interactiveMode, "interactiveMode", "i", false, "use interactive mode")
	updateCmd.Flags().IntVarP(&priorityFlag, "priority", "p", 2, "Specify the priority of todo")
}
