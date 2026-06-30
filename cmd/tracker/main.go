package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"
	"task_tracker/internal/storage"
	"task_tracker/internal/task"

	"github.com/fatih/color"
)

var (
	WrongInputError  = errors.New("WrongInputError")
	EmptyStringError = errors.New("EmptyStringError")
)

func main() {
	jsonStorage := storage.NewJSONStorage("../../tasks.json")
	taskManager := task.NewTaskManaher(jsonStorage)

	var err error
	flag.Parse()
	cmd := flag.Arg(0)

	switch cmd {
	case "add":
		title := strings.Join(flag.Args()[1:], " ")
		err = addTask(taskManager, title)
	case "list":
		status := flag.Arg(1)
		err = listTask(taskManager, status)
	case "mark":
		id := flag.Arg(1)
		status := flag.Arg(2)
		err = updateTask(taskManager, id, "status", status)
	case "update":
		id := flag.Arg(1)
		field := flag.Arg(2)
		newValue := flag.Arg(3)
		err = updateTask(taskManager, id, field, newValue)
	case "delete":
		id := flag.Arg(1)
		err = clearTask(taskManager, id)
	default:
		err = WrongInputError
	}
	if err != nil {
		handleError(err)
	}
}

func addTask(taskManager *task.TaskManager, title string) (err error) {
	if title == "" {
		return EmptyStringError
	}
	err = taskManager.CreateTask(title)
	if err == nil {
		color.Green("Saved")
	}
	return
}

func listTask(taskManager *task.TaskManager, status string) (err error) {
	tasks, err := taskManager.ListFilteredByStatus(status)
	if err != nil {
		return
	}
	color.Cyan("List of tasks:")
	for _, t := range tasks {
		color.Blue("%d %s %s %s", t.ID, t.Task, t.Status, t.CreatedAt.Format("02-01-2006 15:04:05"))
	}
	return
}

func updateTask(taskManager *task.TaskManager, id string, field string, newValue string) (err error) {
	if newValue == "" || id == "" || field == "" {
		return EmptyStringError
	}
	intID, err := convertToInt(id)
	if err != nil {
		return
	}

	err = taskManager.RewriteField(intID, field, newValue)
	if err == nil {
		color.Green("Task has been updated")
	}
	return
}

func clearTask(taskManager *task.TaskManager, id string) (err error) {
	intID, err := convertToInt(id)
	if err != nil {
		return
	}
	err = taskManager.DeleteTask(intID)
	if err == nil {
		color.Green("Task has been deleted")
	}
	return
}

func handleError(err error) {
	switch err {
	case EmptyStringError:
		color.Red("You provided an empty string")
	case WrongInputError:
		color.Red("Wrong input")
	case task.CantModifyTheTaskError:
		color.Red("This field does not exist or can't be modified")
	case task.NoTaskHasBeenFoundError:
		color.Red("No task has been found")
	default:
		color.Red("%s\n", err)
	}
}

func convertToInt(num string) (id int, err error) {
	id, err = strconv.Atoi(num)
	if err != nil {
		err = WrongInputError
	}
	return
}
