package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Capitalize these fields(make them publicly available) in order to be able to use in the JSON Marshall method
// Struct tags are essentialy metadata that can be added to struct fields, added by using backTicks after field name
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("%v\n", todo.Text)

}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("no value provided")
	}
	return Todo{
		Text: text,
	}, nil
}
