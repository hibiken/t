package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
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

func printTodos(todos []Todo, all bool) {
	for _, todo := range todos {
		if !all && todo.Done {
			continue
		}
		fmt.Println(todo.String())
	}
}

// generate pseudo unique id
func genID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String()[:3], nil
}
