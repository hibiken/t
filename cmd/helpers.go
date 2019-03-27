package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
)

// writeToFile writes given list of todos to the specified file.
func writeToFile(todos []*Todo, filename string) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot encode todos to JSON: %v", err)
	}
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		return err
	}
	return nil
}

// readFromFile reads a file and return a list of todos.
func readFromFile(filename string) ([]*Todo, error) {
	bytes, err := ioutil.ReadFile(filename)
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
	for _, todo := range todos {
		if !all && todo.Done {
			continue
		}
		fmt.Println(todo.String())
	}
}

// genID generates pseudo unique id.
func genID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String()[:3], nil
}