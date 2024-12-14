package main

import (
	"todo-list/internal/todo"
)

// main function, will run on execute
func main() {
	// initialize an empty todo list
	todos := todo.Todos{}
	todos.Add("Get Milk")
	todos.Add("Cry")
	todos.Toggle(1)
	todos.Print()
}