package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Readline(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Could not open File!")
		fmt.Println(err)
		return nil, errors.New("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in file")
	}

	file.Close()

	return lines, nil

}
