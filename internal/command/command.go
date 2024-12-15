package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/internal/todo"
)

type CmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Adds a new task to the todo list.")
	flag.IntVar(&cf.Delete, "delete", -1, "Removes the task at the given index.")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the name of a task. id:new_title")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task to be completed. Will set its completion time.")
	flag.BoolVar(&cf.List, "list", false, "Prints the todo list to the console.")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch{
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) > 2 {
			fmt.Println("Error: invalid arguments. Please use id:new_title")
			os.Exit(1)
		}
		
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Not a valid index.")
			os.Exit(1)
		}
		
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.List:
		todos.Print()
	default:
		fmt.Println("Invalid command")
	}
}