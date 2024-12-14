package main

import (
	"fmt"
	"todo-list/internal/todo"
)

// main function, will run on execute
func main() {
	// initialize an empty todo list
	todos := todo.Todos{}
	todos.Add("Get Milk")
	todos.Add("Cry")
	fmt.Printf("%v\n", todos)
	todos.Delete(1)
	fmt.Printf("%v\n", todos)
}