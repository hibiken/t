package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Todo is a task to be done.
type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	Priority  int
}

// NewTodo returns a new todo with a given title.
func NewTodo(title string, priority int) (*Todo, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Todo{
		ID:        id.String()[:3], // Just the first three characters.
		Title:     title,
		CreatedAt: time.Now(),
		Priority:  priority,
	}, nil
}

// CreatedTimeInWords returns a string that indicates the time t was created.
func (t *Todo) CreatedTimeInWords() string {
	const (
		Day   = time.Hour * 24
		Month = Day * 30 // approximate 30 days
		Year  = Month * 12
	)
	d := time.Since(t.CreatedAt)
	if d > Year {
		return fmt.Sprintf("%d years ago", int(d/Year))
	}
	if d > Month {
		return fmt.Sprintf("%d months ago", int(d/Month))
	}
	if d > 2*Day {
		return fmt.Sprintf("%d days ago", int(d/Day))
	}
	if d > Day {
		return "yesterday"
	}
	return "today"
}
