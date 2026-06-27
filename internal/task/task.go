package task

import "time"

type Task struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Тут про бизнес логику
// Создание, обновление, удаление задач
// Валидацию данных
// Поиск и фильтрацию
// Логику предметной области

func CreateTask(title string) (err error) {
	return err
}

func GetTask(id string) (task Task, found bool, err error) {
	return task, found, err
}

func ListAll() (err error) {
	return err
}

func ListFilteredByStatus(status string) (list []Task, err error) {
	return list, err
}

func RewriteField(task Task, field string, new string) (err error) {
	return err
}

func DeleteTask(task Task) (err error) {
	return err
}

// Тут будет основная логика, которая будет вызываться из парсерных функций

// type Status struct {
// 	Done bool
// 	InProgress bool
// 	NotDone bool
// }
