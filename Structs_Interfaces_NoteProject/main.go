package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs_noteproject/note"
	"example.com/structs_noteproject/todo"
)

// define certain constraints that any struct that implements the interface must have the methods on the inerface
// Common naming convention:- since the interface has only one method, we can name with the method name and a suffix  "er"
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// Embedded interface
type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Welcome")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

// interface{} can be used ensure any value can be passed or any
// func printSomething(value any) {
func printSomething(value interface{}) {

	//outside a switch, type checking
	intVal, isOfType := value.(int) // this checks if the value is of type int

	if isOfType {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, isOfType := value.(float64)

	if isOfType {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, isOfType := value.(string)

	if isOfType {
		fmt.Println("String:", stringVal)
		return
	}

	// Example also showing switching on types
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

// Go scans the structs passed should have the methods as defined in the interface wheneever we want to use it, not necessarily need to explicity implement
// the interface ont struct itself
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	data.Save()

	return nil

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
