package task

import "time"

// Task struct to represent a task with basic details
type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}
