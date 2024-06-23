package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs_noteproject/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded!")

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') //use signle quote for single characters like this

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // On windows, line breaks is not just created "\n" but could also include "\r"

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}
