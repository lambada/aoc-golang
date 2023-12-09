package y2023d09p01

import (
	"slices"
	"strconv"
	"strings"
)

func FindNextValue(input []int) int {
	if slices.Max(input) == 0 && slices.Min(input) == 0 {
		return 0
	}

	var diffArray []int

	for i := 0; i < len(input)-1; i++ {
		diffArray = append(diffArray, input[i+1]-input[i])
	}

	nextDiff := FindNextValue(diffArray)

	return input[len(input)-1] + nextDiff

}

func ParseRow(input string) []int {
	stringNumbers := strings.Fields(input)

	var numbers []int

	for _, v := range stringNumbers {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)

	}

	return numbers
}

func Calculate(input string) int {

	rows := strings.Split(input, "\n")

	total := 0
	for _, r := range rows {
		parsedRow := ParseRow(r)
		nextVal := FindNextValue(parsedRow)
		total = total + nextVal
	}

	return total
}
