package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	content := string(data)
	value, err := strconv.ParseFloat(content, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(filename string, value float64) {
	content := fmt.Sprint(value)
	os.WriteFile("balance.txt", []byte(content), 0644)
}
