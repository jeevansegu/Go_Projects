package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/jeevansegu/Go_Projects/ToDoList/cmd"
)

const filepath = "tasks.csv"

func main() {
	rootCmd := &cobra.Command{Use: "tasks"}

	// Define "add" command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(command *cobra.Command, args []string) { // changed 'cmd' to 'command'
			description := args[0]
			err := cmd.AddTask(description, filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			fmt.Println("Task added successfully")
		},
	})

	// Define "list" command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Run: func(command *cobra.Command, args []string) { // changed 'cmd' to 'command'
			showAll, _ := command.Flags().GetBool("all")
			err := cmd.ListTasks(showAll, filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		},
	})

	// Define "complete" command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "complete [taskID]",
		Short: "Mark a task as complete",
		Args:  cobra.MinimumNArgs(1),
		Run: func(command *cobra.Command, args []string) { // changed 'cmd' to 'command'
			taskID, _ := strconv.Atoi(args[0])
			err := cmd.CompleteTask(taskID, filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			fmt.Println("Task marked as complete")
		},
	})

	// Define "delete" command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "delete [taskID]",
		Short: "Delete a task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(command *cobra.Command, args []string) { // changed 'cmd' to 'command'
			taskID, _ := strconv.Atoi(args[0])
			err := cmd.DeleteTask(taskID, filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			fmt.Println("Task deleted successfully")
		},
	})

	// Add flag for listing all tasks
	rootCmd.PersistentFlags().BoolP("all", "a", false, "Show all tasks")

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
