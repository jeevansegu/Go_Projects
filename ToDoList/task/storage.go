package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex

// LoadTasks reads tasks from the CSV file and returns them as a slice of Task structs
func LoadTasks(filepath string) ([]Task, error) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var tasks []Task
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		isComplete, _ := strconv.ParseBool(record[3])
		task := Task{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// SaveTasks writes the current tasks to the CSV file
func SaveTasks(filepath string, tasks []Task) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write task to CSV: %w", err)
		}
	}
	return nil
}
