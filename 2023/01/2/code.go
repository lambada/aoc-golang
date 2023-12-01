package y2023d01p02

import (
	"regexp"
	"strconv"
	"strings"
)

func calculateLine(inputLine string) int {

	noMoreStringNumbers := inputLine

	stringNumberRegex := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
	matchLocation := stringNumberRegex.FindStringIndex(inputLine)
	for matchLocation != nil {
		match := noMoreStringNumbers[matchLocation[0]:matchLocation[1]]
		switch match {
		case "one":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "one", "1", 1)
		case "two":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "two", "2", 1)
		case "three":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "three", "3", 1)
		case "four":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "four", "4", 1)
		case "five":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "five", "5", 1)
		case "six":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "six", "6", 1)
		case "seven":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "seven", "7", 1)
		case "eight":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "eight", "8", 1)
		case "nine":
			noMoreStringNumbers = strings.Replace(noMoreStringNumbers, "nine", "9", 1)
		}
		matchLocation = stringNumberRegex.FindStringIndex(noMoreStringNumbers)
	}
	digitRegex := regexp.MustCompile("[0-9]")
	digits := digitRegex.FindAllString(noMoreStringNumbers, -1)

	firstDigitIndex := 0
	lastDigitIndex := len(digits) - 1
	var answerString []string
	answerString = append(answerString, digits[firstDigitIndex])
	answerString = append(answerString, digits[lastDigitIndex])
	answer, _ := strconv.Atoi(strings.Join(answerString, ""))
	return answer
}

func calculate(input []string) int {

	answer := 0
	for lineNumber := range input {
		answer = answer + calculateLine(input[lineNumber])
	}

	return answer
}
