package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// priority represents urgency of todo.
// p0 being the most urgent.
type priority int

const (
	p0 priority = iota
	p1
	p2
	p3
)

// Todo is a task to be done.
type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	Priority  priority
}

// NewTodo returns a new todo with a given title.
func NewTodo(title string) (*Todo, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Todo{
		ID:        id.String()[:3], // Just the first three characters.
		Title:     title,
		CreatedAt: time.Now(),
		Priority:  p2, // by default, the priority is set to p2
	}, nil
}

// Age returns a string representation of its age.
// if it's less than 1 day old => less than a day
// if older than a day => xxx days old
func (t *Todo) Age() string {
	d := time.Since(t.CreatedAt)
	if d.Hours() < 24 {
		return "less than a day"
	}
	return fmt.Sprintf("%d days old", int(d.Hours()/24))
}
