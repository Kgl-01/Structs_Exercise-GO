package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %s has the following content:\n\n%s\n", note.Title, note.Content)
}
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fmt.Sprintf("%s.json", fileName)
	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	newNote := Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return newNote, nil
}
