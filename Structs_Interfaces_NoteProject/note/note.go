package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Capitalize these fields(make them publicly available) in order to be able to use in the JSON Marshall method
// Struct tags are essentialy metadata that can be added to struct fields, added by using backTicks after field name
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note Note) Display() {
	fmt.Printf("Note title %v has the following content:\n\n%v\n\n", note.Title, note.Content)

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("no value provided")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
