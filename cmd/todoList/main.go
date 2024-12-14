package main

import (
	"todo-list/internal/todo"
	"todo-list/internal/storage"
)

// main function, will run on execute
func main() {
	// initialize an empty todo list
	todos := todo.Todos{}

	// load data from json
	storage := storage.NewStorage[todo.Todos]("todos.json")

	// populate the todos list with the data in the json file
	storage.Load(&todos)
	todos.Add("Get Milk")
	todos.Add("Cry")
	todos.Toggle(1)
	todos.Print()

	// save the data
	storage.Save(todos)
}