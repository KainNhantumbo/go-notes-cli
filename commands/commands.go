package commands

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/KainNhantumbo/go-notes-cli/todo"
)

type CommandFlags struct {
	Add    string
	Remove int
	Toggle int
	Edit   string
	Show   bool
}

func Commands() *CommandFlags {
	cmd := CommandFlags{}
	flag.StringVar(&cmd.Add, "add", "", "Adds a new todo by providing content")
	flag.StringVar(&cmd.Edit, "edit", "", "Edit a todo by index & specify a new content. id:content")
	flag.IntVar(&cmd.Remove, "remove", -1, "Specify a todo by index to delete")
	flag.IntVar(&cmd.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cmd.Show, "show", false, "Show all saved todos")

	flag.Parse()

	return &cmd
}

func (cmd *CommandFlags) Exec(todosList *todo.TodosList) {
	switch {
	case cmd.Show:
		todosList.Print()

	case cmd.Add != "":
		todosList.Add(cmd.Add)

	case cmd.Edit != "":
		parts := strings.SplitN(cmd.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit content. Please use id:content")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit content")
			os.Exit(1)
		}
		todosList.UpdateContent(index, parts[1])

	case cmd.Toggle != -1:
		todosList.ToggleStatus(cmd.Toggle)

	case cmd.Remove != -1:
		todosList.Delete(cmd.Remove)

	default:
		fmt.Println("Invalid input command. Please check and try again.")
	}
}
