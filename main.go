package main

import (
	"todo/cmd"
	"todo/helpers"
)

func main() {
    helpers.CreateTodosTable()
    helpers.RemoveLastDayCompletedTasks()
	cmd.Execute()
}
