package main

import (
	"fmt"
	"sort"

	"github.com/shrayolacrayon/aoc2020/shared"
)

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
	mapped, err := shared.MapLines("report.txt", shared.ToInt)
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
