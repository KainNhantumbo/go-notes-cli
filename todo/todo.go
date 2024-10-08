package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Content   string
	Status    bool
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type TodosList []Todo

func (todosList *TodosList) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todosList) {
		err := errors.New("cannot find the specified todo")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todosList *TodosList) Add(content string) {
	newTodo := Todo{
		Content:   content,
		Status:    false,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	*todosList = append(*todosList, newTodo)
}

func (todosList *TodosList) UpdateContent(index int, content string) error {
	if err := todosList.ValidateIndex(index); err != nil {
		return err
	}

	time := time.Now()
	(*todosList)[index].Content = content
	(*todosList)[index].UpdatedAt = &time

	return nil
}

func (todosList *TodosList) Delete(index int) error {
	if err := todosList.ValidateIndex(index); err != nil {
		return err
	}
	*todosList = append((*todosList)[:index], (*todosList)[index+1:]...)
	return nil
}

func (todosList *TodosList) ToggleStatus(index int) error {
	if err := todosList.ValidateIndex(index); err != nil {
		return err
	}

	(*todosList)[index].Status = !(*todosList)[index].Status
	(*todosList)[index].UpdatedAt = &time.Time{}

	return nil
}

func (todosList *TodosList) GetTodo(index int) (Todo, error) {
	if err := todosList.ValidateIndex(index); err != nil {
		return Todo{}, err
	}

	return (*todosList)[index], nil
}

func (todosList *TodosList) Print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Status", "Created At", "Completed At")

	for index, todo := range *todosList {
		isCompleted := "❌"
		completedAt := ""

		if todo.Status {
			isCompleted = "✅"
			if todo.UpdatedAt != nil {
				completedAt = todo.UpdatedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), todo.Content, isCompleted, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
