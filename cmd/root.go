package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

const filename = "todos.json"

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

// writeToFile writes given list of todos to the specified file.
func writeToFile(todos []Todo, filename string) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot encode todos to json: %v", err)
	}
	if err := ioutil.WriteFile("todos.json", bytes, 0644); err != nil {
		return err
	}
	return nil
}

func readFromFile(filename string) ([]Todo, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, fmt.Errorf("cannot unmarshal to json: %v", err)
	}
	return todos, nil
}

func addTodo(todo Todo, filename string) error {
	todos, err := readFromFile(filename)
	if err != nil {
		return err
	}
	todos = append(todos, todo)
	if err := writeToFile(todos, filename); err != nil {
		return err
	}
	return nil
}

func printTodos(todos []Todo) {
	for i := 0; i < len(todos); i++ {
		fmt.Printf("%d: %s\n", i+1, todos[i].String())
	}
}
