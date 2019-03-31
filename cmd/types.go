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

// Created returns a string that indicates the time t was created.
func (t *Todo) Created() string {
	year, month, day := time.Now().Date()
	createdY, createdM, createdD := t.CreatedAt.Date()
	if year != createdY {
		return fmt.Sprintf("%d years ago", year-createdY)
	}
	if month != createdM {
		return fmt.Sprintf("%d months ago", month-createdM)
	}
	switch day - createdD {
	case 0:
		return "today"
	case 1:
		return "yesterday"
	default:
		return fmt.Sprintf("%d days ago", day-createdD)
	}
}
