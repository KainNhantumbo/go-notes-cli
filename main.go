package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/KainNhantumbo/go-notes-cli/commands"
	"github.com/KainNhantumbo/go-notes-cli/storage"
	"github.com/KainNhantumbo/go-notes-cli/todo"
)

func main() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Could not find the user home directory: ", err)
	}

	var storagePath string = filepath.Join(homeDir, ".todos-data.json")

	todosList := todo.TodosList{}
	storage := storage.NewStorage[todo.TodosList](storagePath)
	storage.Load(&todosList)
	cmdFlags := commands.CommandFlags{}
	cmdFlags.Exec(&todosList)
	storage.Save(todosList)
}
