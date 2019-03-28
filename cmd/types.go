package cmd

import (
	"fmt"
	"time"
)

// Todo is a task to be done.
type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
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
