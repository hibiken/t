package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

// writeToFile writes given list of todos to the specified file.
func writeToFile(todos []*Todo, filepath string) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot encode todos to JSON: %v", err)
	}
	if err := ioutil.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

// readFromFile reads a file and return a list of todos.
func readFromFile(filepath string) ([]*Todo, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var todos []*Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, fmt.Errorf("cannot decode todos from JSON: %v", err)
	}
	return todos, nil
}

// printTodos prints todos. If all is false, it prints only the items with Done
// field set to false, otherwise it prints all items in the slice.
func printTodos(todos []*Todo, all bool) {
	if len(todos) == 0 {
		fmt.Println("There are no todos :)")
		return
	}
	undones := filter(todos, func(t *Todo) bool {
		return !t.Done
	})
	if len(undones) == 0 {
		fmt.Println("You'are all done 🎉")
		if !all {
			return
		}
		fmt.Println()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Status", "Age"})
	table.SetBorder(false)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})

	for _, t := range todos {
		if !all && t.Done {
			continue
		}
		status := ""
		if t.Done {
			status = "   \u2714   "
		}
		table.Append([]string{t.ID, t.Title, status, t.Age()})
	}
	table.Render()
}

// filter filters todos with a given predicate function p.
func filter(todos []*Todo, p func(*Todo) bool) []*Todo {
	var res []*Todo
	for _, t := range todos {
		if p(t) {
			res = append(res, t)
		}
	}
	return res
}
