package storage

import (
	"encoding/json"
	"json"
	"task_tracker/internal/task"
)

// Отвечает за:

// Чтение и запись данных в файл (JSON, SQLite, etc.)

// Преобразование структур в байты и обратно

// Работу с файловой системой

type Storage interface {
	Save(tasks []task.Task)error
	Load() ([]task.Task, error)
}

type JSONStorage struct {
	fileName string
}

func NewJSONStorage(fileName string) *JSONStorage {
	return &JSONStorage{fileName: fileName}
}

func (js *JSONStorage) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	return err
}

func (js *JSONStorage) Load() (tasks []task.Task, err error) {
    tasks = json.Unmarshal(data, &tasks)
	return nil, nil
}