package main

import (
	"fmt"
	"strconv"

	"github.com/shrayolacrayon/aoc2020/shared"
)

type Seat struct {
	Name   string
	Value  int64
	Row    int64
	Column int64
}

// FBFBBFFRLR
// 0101100
// F is lower B is higher
func parseCode(code string) (interface{}, error) {
	binaryStr := ""
	for i := 0; i < 7; i++ {
		if code[i] == 'F' {
			binaryStr += "0"
		} else {
			binaryStr += "1"
		}
	}
	frontBack, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return 0, nil
	}

	binaryStr = ""
	for i := 7; i < 10; i++ {
		if code[i] == 'L' {
			binaryStr += "0"
		} else {
			binaryStr += "1"
		}
	}
	leftRight, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return 0, nil
	}
	//fmt.Println((frontBack * 8) + leftRight)
	return Seat{Name: code, Value: (frontBack * 8) + leftRight, Row: frontBack, Column: leftRight}, nil
}

func main() {
	outputs, err := shared.MapLines("data.txt", parseCode)
	if err != nil {
		fmt.Println(err)
	}
	maxVal := int64(0)
	maxCode := ""
	takenSeats := make([][]int, 128)
	for i := 0; i < 128; i++ {
		takenSeats[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			takenSeats[i][j] = 0
		}
	}

	for _, output := range outputs {
		seat, ok := output.(Seat)
		if !ok {
			fmt.Println("couldnt convert output")
		}
		if seat.Value > maxVal {
			maxVal = seat.Value
			maxCode = seat.Name
		}
		takenSeats[seat.Row][seat.Column] = 1
	}

	for i := 5; i < len(takenSeats)-14; i++ {
		for j := range takenSeats[i] {
			if takenSeats[i][j] == 0 {
				fmt.Println(i*8 + j)
			}
		}
	}

}
