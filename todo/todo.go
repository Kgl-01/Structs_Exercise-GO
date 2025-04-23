package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("ivalid input")
	}

	newTodo := Todo{
		Text: text,
	}

	return newTodo, nil
}
