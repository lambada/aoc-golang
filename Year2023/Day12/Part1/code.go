package y2023d12p01

import (
	"regexp"
	"strconv"
	"strings"
)

func TestCombination(layout string, quantities []int) bool {
	// layout = "###.##.####"
	// quantities = [ 2 1 ... ]

	segment := regexp.MustCompile("#+")
	matches := segment.FindAllString(layout, -1)

	if len(matches) != len(quantities) {
		return false
	}

	for i := range quantities {
		if quantities[i] != len(matches[i]) {
			return false
		}
	}

	return true

}

func ComputeCombinations(layout string, quantities []int) int {
	// layout = "##???#?"
	// quantities = [ 2 1 ... ]

	if !strings.Contains(layout, "?") {
		if TestCombination(layout, quantities) {
			return 1
		}
		return 0
	}

	matchingCombinations := 0

	layoutCandidate := strings.Replace(layout, "?", "#", 1)
	matchingCombinations = matchingCombinations + ComputeCombinations(layoutCandidate, quantities)
	layoutCandidate = strings.Replace(layout, "?", ".", 1)
	matchingCombinations = matchingCombinations + ComputeCombinations(layoutCandidate, quantities)

	return matchingCombinations
}
func ParseRow(input string) int {

	inputParts := strings.Fields(input)

	sizesString := strings.Split(inputParts[1], ",")
	var sizes []int
	for _, str := range sizesString {
		sizeInt, _ := strconv.Atoi(str)
		sizes = append(sizes, sizeInt)
	}

	layout := inputParts[0]

	// duplicate empty spaces are redundant
	manyDots := regexp.MustCompile("[.]+")
	layout = manyDots.ReplaceAllString(layout, ".")

	// At this point we can guarentee that len(sizes) >= len(potentialSpringSegments)

	return ComputeCombinations(layout, sizes)
}

func Calculate(input string) int {

	rows := strings.Split(input, "\n")

	total := 0
	for _, row := range rows {
		total = total + ParseRow(row)
	}
	return total
}
