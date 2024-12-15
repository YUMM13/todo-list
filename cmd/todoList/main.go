package main

import (
	"todo-list/internal/todo"
	"todo-list/internal/storage"
	"todo-list/internal/command"
)

// main function, will run on execute
func main() {
	// initialize an empty todo list
	todos := todo.Todos{}

	// load data from json
	storage := storage.NewStorage[todo.Todos]("../../todos.json")
	storage.Load(&todos)

	// launch cli
	cmdFlags := command.NewCmdFlags()
	cmdFlags.Execute(&todos)

	// save the data
	storage.Save(todos)
}