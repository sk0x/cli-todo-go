## Todo CLI

A command-line todo application written in Go that helps you manage tasks with simplicity and efficiency.

## Features

- Add new tasks
- List all tasks with their status
- Mark tasks as complete
- Delete tasks
- Automatic cleanup of completed tasks from the previous day

## Installation

### Prerequisites

- Go 1.16 or higher
- SQLite3

### Steps

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/todo-cli.git
   cd todo-cli
   ```

2. Build the application
   ```bash
   go build -o todo
   ```

3. Move the binary to your PATH (optional)
   ```bash
   sudo mv todo /usr/local/bin/
   ```

## Usage

### Adding a task

```bash
todo add -t "Complete the project documentation"
```

### Listing all tasks

```bash
todo list
```

This will display a formatted table with columns for ID, task description, status, and creation date.

### Marking a task as complete

```bash
todo complete -i 1
```

Where `1` is the ID of the task you want to mark as complete.

### Deleting a task

```bash
todo delete -i 1
```

Where `1` is the ID of the task you want to delete.

## Database

The application uses SQLite3 as its database engine. The database file (`todo.db`) is stored in your home directory.

## Automatic Cleanup

When you start the application, it automatically removes all completed tasks from the previous day, helping you keep your task list clean and focused on current items.

## Command Details

| Command | Flag | Description |
|---------|------|-------------|
| add | -t, --task | Add a new task |
| list | | List all tasks |
| list | -F, --top | Show top five tasks (not fully implemented) |
| list | -B, --bottom | Show last five tasks (not fully implemented) |
| complete | -i, --id | Mark a task as complete |
| delete | -i, --id | Delete a task |

## Project Structure

- `main.go` - Entry point of the application
- `cmd/` - Contains command definitions using Cobra
  - `root.go` - Root command setup
  - `todo.go` - Command implementations
- `helpers/` - Helper functions
  - `sql_helpers.go` - Database operations

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - Command-line interface framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver for Go

## Future Improvements

- Implement filtering for list command (top/bottom functionality)
- Add due dates for tasks
- Add priority levels
- Implement task categories/tags cli-todo-go
