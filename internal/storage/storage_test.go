package storage

import (
	"os"
	"task_tracker/internal/models"
	"testing"
	"time"
)

func TestSaveAndLoad(t *testing.T) {
	file := "test.json"
	defer os.Remove(file)

	storage := NewJSONStorage(file)

	timeNow := time.Now()

	task := models.Task{
		ID:        1,
		Task:      "some task",
		Status:    "done",
		CreatedAt: timeNow,
	}
	tasks := []models.Task{task}

	err := storage.Save(tasks)
	if err != nil {
		t.Errorf("Save error: %v", err)
	}

	loadedTasks, err := storage.Load()
	if err != nil {
		t.Errorf("Load error: %v", err)
	}
	if len(loadedTasks) != 1 {
		t.Errorf("Expected 1, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Task != "some task" {
		t.Errorf("Expected 'some task', got %s", loadedTasks[0].Task)
	}
}

func TestEmptyFile(t *testing.T) {
	file := "test.json"
	defer os.Remove(file)

	storage := NewJSONStorage(file)

	loadedTasks, err := storage.Load()
	if err != nil {
		t.Errorf("Load error: %v", err)
	}
	if len(loadedTasks) != 0 {
		t.Errorf("Expected 0, got %d", len(loadedTasks))
	}
}
