package cmd

import (
	"fmt"

	"github.com/jeevansegu/Go_Projects/ToDoList/task"
)

// CompleteTask marks a task as completed by its ID
func CompleteTask(taskID int, filepath string) error {
	tasks, err := task.LoadTasks(filepath)
	if err != nil {
		return fmt.Errorf("could not load tasks: %v", err)
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].IsComplete = true
			break
		}
	}

	return task.SaveTasks(filepath, tasks)
}
