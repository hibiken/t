package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func contains(elements []string, target string) bool {
	for _, elem := range elements {
		if elem == target {
			return true
		}
	}
	return false
}

func findByID(todos []*Todo, id string) *Todo {
	var t *Todo
	for _, todo := range todos {
		if todo.ID == id {
			t = todo
			break
		}
	}
	return t
}

func findByTitle(todos []*Todo, title string) *Todo {
	var t *Todo
	for _, todo := range todos {
		if todo.Title == title {
			t = todo
			break
		}
	}
	return t
}

// Sort todos by Priority (primary sort key) and CreatedAt (secondary sort key).
func sortTodos(todos []*Todo) {
	sort.Slice(todos, func(i, j int) bool {
		x, y := todos[i], todos[j]
		if x.Priority != y.Priority {
			return x.Priority < y.Priority
		}
		return x.CreatedAt.After(y.CreatedAt)
	})
}

func any(todos []*Todo, fn func(*Todo) bool) bool {
	for _, t := range todos {
		if fn(t) {
			return true
		}
	}
	return false
}

func titles(todos []*Todo) []string {
	var res []string
	for _, t := range todos {
		res = append(res, t.Title)
	}
	return res
}

func printErrorAndExit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
