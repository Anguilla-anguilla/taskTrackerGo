package storage

import (
	"encoding/json"
	"os"
	"task_tracker/internal/models"
)

type Storage interface {
	Save(tasks []models.Task) error
	Load() ([]models.Task, error)
}

type JSONStorage struct {
	fileName string
}

func NewJSONStorage(fileName string) *JSONStorage {
	return &JSONStorage{fileName: fileName}
}

func (js *JSONStorage) Save(tasks []models.Task) (err error) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(js.fileName, data, 0644)
	if err != nil {
		return err
	}
	return
}

func (js *JSONStorage) Load() (tasks []models.Task, err error) {
	data, err := os.ReadFile(js.fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return
	}

	return tasks, err
}
