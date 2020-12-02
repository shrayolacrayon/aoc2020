package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
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

func toInt(s string) (interface{}, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	return strconv.Atoi(strings.TrimSpace(s))
}

func getMultipleOfSumVals(input []int, target int) int {
	leftPointer := 0
	rightPointer := len(input) - 1
	for rightPointer > leftPointer {
		sum := input[rightPointer] + input[leftPointer]
		if sum == target {
			return input[rightPointer] * input[leftPointer]
		}
		if sum < target {
			leftPointer += 1
		} else {
			rightPointer -= 1
		}
	}
	return 0
}

func getThreeSums(input []int, target int) int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if (input[i] + input[j] + input[k]) == target {
					fmt.Printf("i %v, j %v, k %v \n", input[i], input[j], input[k])
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0

}

func main() {
	mapped, err := MapLines("report.txt", toInt)
	if err != nil {
		fmt.Printf("error getting data %s \n", err)
	}
	intMapped := []int{}
	for _, m := range mapped {
		i, ok := m.(int)
		if !ok {
			fmt.Println("not okay")
		}
		intMapped = append(intMapped, i)
	}
	sortedInts := sort.IntSlice(intMapped)
	sortedInts.Sort()
	fmt.Println(getMultipleOfSumVals(sortedInts, 2020))
	multiple := getThreeSums(sortedInts, 2020)
	fmt.Println(multiple)

}
