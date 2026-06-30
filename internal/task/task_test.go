package task

import (
	"task_tracker/internal/models"
	"testing"
)

type MockStorage struct {
	tasks []models.Task
}

func (m *MockStorage) Save(tasks []models.Task) error {
	m.tasks = tasks
	return nil
}

func (m *MockStorage) Load() ([]models.Task, error) {
	return m.tasks, nil
}

func TestCreateTask(t *testing.T) {
	storage := &MockStorage{}
	manager := NewTaskManager(storage)

	err := manager.CreateTask("some task")
	if err != nil {
		t.Errorf("CreateTask error: %v", err)
	}

	loadedTasks, _ := storage.Load()
	if len(loadedTasks) != 1 {
		t.Errorf("Expected 1, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Task != "some task" {
		t.Errorf("Expected 'some task', got %s", loadedTasks[0].Task)
	}
}

func TestDeleteTask(t *testing.T) {
	storage := &MockStorage{
		tasks: []models.Task{
			{ID: 1, Task: "1"},
			{ID: 2, Task: "2"},
		},
	}
	manager := NewTaskManager(storage)

	err := manager.DeleteTask(1)
	if err != nil {
		t.Errorf("DeleteTask error %v", err)
	}

	loadedTasks, _ := storage.Load()
	if len(loadedTasks) != 1 {
		t.Errorf("Expected 1, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Task != "2" {
		t.Errorf("Expected '2', got %s", loadedTasks[0].Task)
	}

	err = manager.DeleteTask(9)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}

func TestRewriteField(t *testing.T) {
	storage := &MockStorage{
		tasks: []models.Task{
			{ID: 1, Task: "old", Status: "pending"},
		},
	}
	manager := NewTaskManager(storage)

	err := manager.RewriteField(1, "task", "new")
	if err != nil {
		t.Errorf("REwriteField error: %v", err)
	}

	loadedTasks, _ := storage.Load()
	if loadedTasks[0].Task != "new" {
		t.Errorf("Expected 'new', got %s", loadedTasks[0].Task)
	}

	err = manager.RewriteField(1, "status", "done")
	if err != nil {
		t.Errorf("RewriteField error: %v", err)
	}

	loadedTasks, _ = storage.Load()
	if loadedTasks[0].Status != "done" {
		t.Errorf("Expected 'done', got %s", loadedTasks[0].Status)
	}

	err = manager.RewriteField(1, "whatever", "whatever")
	if err == nil {
		t.Errorf("Expected error for unknown field")
	}
}
