package cmd

import (
	"fmt"
	"time"

	"github.com/jeevansegu/Go_Projects/ToDoList/task"
)

// AddTask adds a new task with the provided description
func AddTask(description string, filepath string) error {
	tasks, err := task.LoadTasks(filepath)
	if err != nil {
		return fmt.Errorf("could not load tasks: %v", err)
	}

	newTask := task.Task{
		ID:          len(tasks) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	tasks = append(tasks, newTask)
	return task.SaveTasks(filepath, tasks)
}
