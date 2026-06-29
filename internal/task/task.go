package task

import (
	"errors"
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

func (s *TaskManager) CreateTask(taskText string) (err error) {
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}

	newTask := Task{
		ID:        len(tasks) + 1,
		Task:      taskText,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, newTask)
	err = s.storage.Save()
	return err
}

func (s *TaskManager) GetTask(id int) (task Task, err error) {
	tasks, err := s.storage.Load()
	if err != nil {
		return
	}
	for _, v := range tasks {
		if v.ID == id {
			return v, err
		}
	}
	return Task{}, NoTaskHasBeenFoundError
}

func (s *TaskManager) ListFilteredByStatus(status string) (list []Task, err error) {
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
	task, err := s.GetTask(id)
	if err != nil {
		return
	}

	switch field {
	case "task":
		task.Task = new
		err = s.storage.Save()
	case "status":
		task.Status = new
		err = s.storage.Save()
	default:
		err = CantModifyTheTaskError
	}
	return err
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
			err = s.storage.Save()
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
