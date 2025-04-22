package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %s has the following content:\n\n%s\n", note.title, note.content)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	newNote := Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}

	return newNote, nil
}
