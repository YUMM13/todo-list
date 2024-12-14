package todo

import (
	"errors"
	"fmt"
	"time"
)

// struct for todo list items, will bundle necessary vars together
type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time // is a pointer because if the task is not completed, it will be null
}

// create a slice (dynamically size array) for todo items
type Todos []Todo

// method associated with type Todos. Now when we declare a Todos variable, we can call todos.add(...)
func (todos *Todos) Add(title string) {
	// create a new todo item to add to our slice
	newTask := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	// add to our slice, need to dereference it since todos is a pointer to our main slice
	*todos = append(*todos, newTask)
}

// helper method that will check if indexes for operations like remove are invalid
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// remove a todo item
func (todos *Todos) Delete(index int) error {
	t := *todos

	// return an error if there is one
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// remove the item by splitting the list around it and having it point to that new list
	*todos = append(t[:index], t[index+1:]...)

	// no need to return a value, so just return nill
	return nil
}