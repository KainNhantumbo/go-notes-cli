package main

import (
	"github.com/KainNhantumbo/go-notes-cli/commands"
	"github.com/KainNhantumbo/go-notes-cli/storage"
	"github.com/KainNhantumbo/go-notes-cli/todo"
)

func main() {
	todosList := todo.TodosList{}
	storage := storage.NewStorage[todo.TodosList]("internal-data.json")
	storage.Load(&todosList)
	cmdFlags := commands.CommandFlags{}
	cmdFlags.Exec(&todosList)
	storage.Save(todosList)
}
