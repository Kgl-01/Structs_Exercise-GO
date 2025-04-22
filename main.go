package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kgl-01/Structs_Exercise-GO.git/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()
	saveError := userNote.Save()

	if saveError != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded.")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
