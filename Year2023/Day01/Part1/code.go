package y2023d01p01

import (
	"regexp"
	"strconv"
	"strings"
)

func calculateLine(inputLine string) int {

	digitRegex := regexp.MustCompile("[0-9]")
	digits := digitRegex.FindAllString(inputLine, -1)

	firstDigitIndex := 0
	lastDigitIndex := len(digits) - 1
	var answerString []string
	answerString = append(answerString, digits[firstDigitIndex])
	answerString = append(answerString, digits[lastDigitIndex])
	answer, _ := strconv.Atoi(strings.Join(answerString, ""))
	return answer
}

func Calculate(input []string) int {

	answer := 0
	for lineNumber := range input {
		answer = answer + calculateLine(input[lineNumber])
	}

	return answer
}
