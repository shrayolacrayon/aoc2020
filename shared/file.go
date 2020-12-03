package shared

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoMem(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	return string(data), nil
}

func MapLines(filepath string, lineFunc func(string) (interface{}, error)) ([]interface{}, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	outputs := []interface{}{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var readErr error
	for {
		line, readErr = reader.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			break
		}
		output, err := lineFunc(line)
		if err != nil {
			return nil, err
		}
		if output != nil {
			outputs = append(outputs, output)
		}

		if readErr != nil {
			break
		}

	}

	if readErr != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}

	return outputs, nil
}

func ToInt(s string) (interface{}, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	return strconv.Atoi(strings.TrimSpace(s))
}
