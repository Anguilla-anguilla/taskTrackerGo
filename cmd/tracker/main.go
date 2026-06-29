package main

import (
	"errors"
	"flag"
	"internal/strconv"
	"strings"
	"task_tracker/internal/task"

	"github.com/fatih/color"
)

// тут должны остаться только ошибки пользовательского ввода
var (
	WrongInputError  = errors.New("WrongInputError")
	EmptyStringError = errors.New("EmptyStringError")
)

func main() {
	var err error
	flag.Parse()
	cmd := flag.Arg(0)

	switch cmd {
	case "add":
		title := strings.Join(flag.Args()[1:], " ")
		err = addTask(title)
	case "list":
		status := flag.Arg(1)
		err = listTask(status)
	case "mark":
		id := flag.Arg(1)
		status := flag.Arg(2)
		err = updateTask(id, "status", status)
	case "update":
		id := flag.Arg(1)
		field := flag.Arg(2)
		newValue := flag.Arg(3)
		err = updateTask(id, field, newValue)
	case "delete":
		id := flag.Arg(1)
		err = clearTask(id)
	default:
		err = WrongInputError
	}
	if err != nil {
		handleError(err)
	}
}

func addTask(title string) (err error) {
	if title == "" {
		return EmptyStringError
	}
	err = task.CreateTask(title)
	if err == nil {
		color.Green("Saved")
	}
	return
}

func listTask(status string) (err error) {
	tasks, err := task.ListFilteredByStatus(status)
	if err != nil {
		return
	}
	color.Cyan("List of tasks:")
	for _, t := range tasks {
		color.Blue("%d %s %s %s", t.ID, t.Task, t.Status, t.CreatedAt.Format("02-01-2006 15:04:05"))
	}
	return
}

func updateTask(id string, field string, newValue string) (err error) {
	if newValue == "" || id == "" || field == "" {
		return EmptyStringError
	}
	intID, err = convertToInt(id)
	if err != nil {
		return
	}

	err = task.RewriteField(intID, field, newValue)
	if err == nil {
		color.Green("Task has been updated")
	}
	return
}

func clearTask(id string) (err error) {
	intID, err = convertToInt(id)
	if err != nil {
		return
	}
	err = task.DeleteTask(intID)
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