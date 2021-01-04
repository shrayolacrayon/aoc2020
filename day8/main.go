package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shrayolacrayon/aoc2020/shared"
)

type ActionItem struct {
	Action   string
	Positive bool
	Amount   int
}

func parseLine(s string) (interface{}, error) {
	actions := strings.Split(s, " ")
	var positive bool
	if actions[1][0] == '+' {
		positive = true
	}
	a := strings.TrimSpace(actions[1][1:])
	amt, err := strconv.Atoi(a)
	if err != nil {
		return nil, err
	}
	return ActionItem{
		Action:   strings.TrimSpace(actions[0]),
		Positive: positive,
		Amount:   amt,
	}, nil
}

func getAcc() int {
	actions, err := shared.MapLines("data.txt", parseLine)
	if err != nil {
		fmt.Println(err)
	}

	acc := 0
	executed := map[int]struct{}{}
	i := 0
	for {
		if i < 0 {
			break
		}
		action, ok := actions[i].(ActionItem)
		if !ok {
			fmt.Println("error not an action")
		}
		if _, ok := executed[i]; ok {
			break
		}
		executed[i] = struct{}{}
		switch action.Action {
		case "acc":
			//fmt.Printf("%v %d \n", action.Positive, action.Amount)
			if action.Positive {
				acc += action.Amount
			} else {
				acc -= action.Amount
			}
			i += 1

		case "jmp":
			if action.Positive {
				i += action.Amount
			} else {
				i -= action.Amount
			}
		case "nop":
			i += 1
		}
	}
	return acc
}
func findNopOrJump(start int, actions []ActionItem) (string, int) {
	for i := start; i < len(actions); i++ {
		a := actions[i]
		if a.Action == "jmp" {
			return "nop", i
		}
		if a.Action == "nop" {
			return "jmp", i
		}
	}
	return "", len(actions)
}

func getAcc(i int) {
	actions, err := shared.MapLines("data.txt", parseLine)
	if err != nil {
		fmt.Println(err)
	}

	acc := 0
	executed := map[int]struct{}{}
	i := 0
	for {
		if i < 0 {
			break
		}
		action, ok := actions[i].(ActionItem)
		if !ok {
			fmt.Println("error not an action")
		}
		if _, ok := executed[i]; ok {
			break
		}
		executed[i] = struct{}{}
		switch action.Action {
		case "acc":
			//fmt.Printf("%v %d \n", action.Positive, action.Amount)
			if action.Positive {
				acc += action.Amount
			} else {
				acc -= action.Amount
			}
			i += 1

		case "jmp":
			if action.Positive {
				i += action.Amount
			} else {
				i -= action.Amount
			}
		case "nop":
			i += 1
		}
	}
	return acc
}

func main() {

	actions, err := shared.MapLines("data.txt", parseLine)
	if err != nil {
		fmt.Println(err)
	}

	acc := 0
	executed := map[int]struct{}{}
	i := 0
	for {
		if i < 0 {
			break
		}
		action, ok := actions[i].(ActionItem)
		if !ok {
			fmt.Println("error not an action")
		}
		if _, ok := executed[i]; ok {
			break
		}
		executed[i] = struct{}{}
		switch action.Action {
		case "acc":
			//fmt.Printf("%v %d \n", action.Positive, action.Amount)
			if action.Positive {
				acc += action.Amount
			} else {
				acc -= action.Amount
			}
			i += 1

		case "jmp":
			if action.Positive {
				i += action.Amount
			} else {
				i -= action.Amount
			}
		case "nop":
			i += 1
		}
	}
	fmt.Println(acc)
}
