package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shrayolacrayon/aoc2020/shared"
)

func parseAdjs(s string) string {
	s = strings.Join(strings.Split(s, " ")[:2], " ")
	return s
}

func parseLine(s string) (interface{}, error) {
	output := map[string]map[string]int{}
	containSplit := strings.Split(s, "contain")
	key := containSplit[0]
	key = parseAdjs(s)

	vals := map[string]int{}
	for _, content := range strings.Split(containSplit[1], ",") {
		content = strings.TrimSpace(content)
		if content == "no other bags." {
			break
		}
		content = parseAdjs(content)
		amt, err := strconv.Atoi(string(content[0]))
		if err != nil {
			return nil, err
		}
		vals[content[2:]] = amt
	}
	output[key] = vals
	return output, nil

}

func main() {
	outputs, err := shared.MapLines("sample.txt", parseLine)
	if err != nil {
		fmt.Printf("error: %v \n", err)
	}
	//count := 0
	allVals := map[string]map[string]int{}
	for _, output := range outputs {
		o, ok := output.(map[string]map[string]int)
		if !ok {
			fmt.Println("error converting output")
		}
		for key, vals := range o {
			allVals[key] = vals
		}
	}
	for key, vals := range allVals {

	}

}
