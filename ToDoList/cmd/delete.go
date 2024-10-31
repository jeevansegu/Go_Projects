package cmd

import (
	"fmt"

	"github.com/jeevansegu/Go_Projects/ToDoList/task"
)

// DeleteTask removes a task by its ID
func DeleteTask(taskID int, filepath string) error {
	tasks, err := task.LoadTasks(filepath)
	if err != nil {
		return fmt.Errorf("could not load tasks: %v", err)
	}

	var updatedTasks []task.Task
	for _, task := range tasks {
		if task.ID != taskID {
			updatedTasks = append(updatedTasks, task)
		}
	}

	return task.SaveTasks(filepath, updatedTasks)
}
