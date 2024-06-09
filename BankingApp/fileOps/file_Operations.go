package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Showing example of splitting code into multiple packages

// In order to export a function from another package, it must starts with a capital letter
func GetFloatFromFile(fileName string) (float64, error) {
	// Typically errors doesn't crash proejct in Go, rather returns a default value
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	// 0644 is a file permission notation for operating systems(especially) that allows owner to read and write but others can only read it
	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
	os.WriteFile(fileName, []byte(valueText), 0644)
}
