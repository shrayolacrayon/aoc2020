package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseFile(filepath string) ([]int, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var readErr error
	outputs := []int{}
	used := map[rune]struct{}{}
	for {
		line, readErr = reader.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			break
		}

		if line == "\n" {
			outputs = append(outputs, len(used))
			used = map[rune]struct{}{}
		} else {
			for _, char := range strings.TrimSpace(line) {
				used[char] = struct{}{}
			}
		}

		if readErr != nil {
			break
		}

	}
	outputs = append(outputs, len(used))

	if readErr != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}

	return outputs, nil
}

func parseFile2(filepath string) ([]int, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var readErr error
	outputs := []int{}
	used := map[rune]int{}
	lineLength := 0
	for {
		line, readErr = reader.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			break
		}
		if line == "\n" {
			fmt.Println(used)
			fmt.Println(lineLength)

			count := 0
			for _, use := range used {
				if use == lineLength {
					count += 1
				}
			}
			outputs = append(outputs, count)
			used = map[rune]int{}
			lineLength = 0
		} else {
			lineLength += 1
			for _, char := range strings.TrimSpace(line) {
				if lineLength > 1 {
					if _, ok := used[char]; ok {
						used[char] += 1
					}
				} else {
					used[char] = 1
				}

			}
		}

		if readErr != nil {
			break
		}

	}

	count := 0
	for _, use := range used {
		if use == lineLength {
			count += 1
		}
	}
	outputs = append(outputs, count)

	if readErr != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}

	return outputs, nil
}
func main() {
	outputs, err := parseFile2("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, count := range outputs {
		sum += count
	}
	fmt.Println(sum)
}
