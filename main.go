package main

import (
	"fmt"

	"github.com/Kgl-01/Structs_Exercise-GO.git/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)

	return userInput
}
