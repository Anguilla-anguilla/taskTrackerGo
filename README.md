# Task Tracker
A simple CLI task tracker written in Go.

## Features
- Add a new task
- List all tasks
- Filter tasks by status (`pending`, `in progress`, `done`)
- Update task fields
- Delete a task
- Persistent storage in JSON file
- Colored output using `fatih/color`

## Installation
1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Build the binary:  `go build -o task-tracker ./cmd/tracker`

## Commands
`add` - Add a new task 
`list` - List all tasks 
`list [status]` - List tasks by status
`update [id] [field] [value]` - Update task field
`update [id] status [value]` - Update task status
`delete [id]` - Delete a task
