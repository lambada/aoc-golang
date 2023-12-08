package y2023d08p01

import (
	"regexp"
	"strings"
)

type Options struct {
	L string
	R string
}

func ParseNode(input string) (string, Options) {

	inputParts := strings.Split(input, " = ")
	nodeName := inputParts[0]

	optionsRegex := regexp.MustCompile(`\(([A-Za-z]+), ([A-Za-z]+)\)`)

	options := optionsRegex.FindAllStringSubmatch(inputParts[1], -1)

	return nodeName, Options{L: options[0][1], R: options[0][2]}

}

func Calculate(input string) int {

	rows := strings.Split(input, "\n")

	path := strings.Split(rows[0], "")

	network := map[string]Options{}

	for i := 2; i < len(rows); i++ {
		nodeName, options := ParseNode(rows[i])
		network[nodeName] = options
	}

	currentNode := "AAA"
	stepCount := 0
	for pathIndex := 0; currentNode != "ZZZ"; pathIndex = (pathIndex + 1) % len(path) {
		if path[pathIndex] == "R" {
			currentNode = network[currentNode].R
		} else if path[pathIndex] == "L" {
			currentNode = network[currentNode].L
		}
		stepCount++
	}

	return stepCount
}
