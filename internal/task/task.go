package task

import (
	"errors"
	"task_tracker/internal/models"
	"task_tracker/internal/storage"
	"time"
)

var (
	CantModifyTheTaskError  = errors.New("CantModifyTheTaskError")
	NoTaskHasBeenFoundError = errors.New("NoTaskHasBeenFoundError")
	WrongStatusError        = errors.New("WrongStatusError")
)

type TaskManager struct {
	storage storage.Storage
}

func NewTaskManaher(s storage.Storage) *TaskManager {
	return &TaskManager{storage: s}
}

func (s *TaskManager) CreateTask(taskText string) (err error) {
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}

	newTask := models.Task{
		ID:        len(tasks) + 1,
		Task:      taskText,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, newTask)
	err = s.storage.Save(tasks)
	return err
}

func (s *TaskManager) ListFilteredByStatus(status string) (list []models.Task, err error) {
	if !checkStatus(status) {
		return list, WrongStatusError
	}
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}
	if status == "" {
		return tasks, err
	}
	for _, v := range tasks {
		if v.Status == status {
			list = append(list, v)
		}
	}
	return list, err
}

func (s *TaskManager) RewriteField(id int, field string, new string) (err error) {
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			switch field {
			case "task":
				tasks[i].Task = new
			case "status":
				tasks[i].Status = new
			default:
				err = CantModifyTheTaskError
			}
			found = true
			break
		}
	}

	if !found {
		return NoTaskHasBeenFoundError
	}

	return s.storage.Save(tasks)
}

func (s *TaskManager) DeleteTask(id int) (err error) {
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}

	found := false
	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err = s.storage.Save(tasks)
			found = true
			break
		}
	}
	if !found {
		err = NoTaskHasBeenFoundError
	}
	return err
}

func checkStatus(status string) bool {
	return status == "pending" || status == "in progress" || status == "done" || status == ""
}
