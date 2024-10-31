package cmd

import (
	"fmt"

	"github.com/jeevansegu/Go_Projects/ToDoList/task"
)

// ListTasks displays tasks based on the showAll flag
func ListTasks(showAll bool, filepath string) error {
	tasks, err := task.LoadTasks(filepath)
	if err != nil {
		return fmt.Errorf("could not load tasks: %v", err)
	}

	for _, task := range tasks {
		if showAll || !task.IsComplete {
			fmt.Printf("%d\t%s\t%s\t%v\n", task.ID, task.Description, task.CreatedAt.Format("2006-01-02 15:04:05"), task.IsComplete)
		}
	}
	return nil
}
