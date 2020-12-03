package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shrayolacrayon/aoc2020/shared"
)

type Password struct {
	Beginning int
	End       int
	Character string
	Password  string
}

func parseLine(line string) (interface{}, error) {
	p := Password{}
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("unknown format of line %v", line)
	}
	patterns := strings.Split(parts[0], "-")
	if len(patterns) < 2 {
		return nil, fmt.Errorf("unknown format of pattern %v", patterns)
	}

	beg, err := strconv.Atoi(string(patterns[0]))
	if err != nil {
		return nil, err
	}
	p.Beginning = beg

	subPattern := strings.Split(patterns[1], " ")
	if len(subPattern) < 2 {
		return nil, fmt.Errorf("unknown format of pattern %v", patterns)

	}
	end, err := strconv.Atoi(string(subPattern[0]))
	if err != nil {
		return nil, err
	}
	p.End = end

	p.Character = subPattern[1]
	p.Password = strings.TrimSpace(parts[1])
	fmt.Printf("%#v \n", p)
	return p, nil
}

func isValidPassword(p Password) bool {
	numChars := strings.Count(p.Password, p.Character)
	return numChars <= p.End && numChars >= p.Beginning
}

func isPositionValidPassword(p Password) bool {
	if p.Password[p.Beginning-1] == p.Character[0] {
		if p.Password[p.End-1] == p.Character[0] {
			return false
		}
		return true
	}
	if p.Password[p.End-1] == p.Character[0] {
		return true
	}
	return false
}

func main() {
	lines, err := shared.MapLines("data.txt", parseLine)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}
	count := 0
	count2 := 0
	for _, elt := range lines {
		p, ok := elt.(Password)
		if !ok {
			fmt.Println("was not a password")
		}
		if isValidPassword(p) {
			count += 1
		}
		if isPositionValidPassword(p) {
			count2 += 1
		}
	}
	fmt.Println(count)
	fmt.Println(count2)

}
