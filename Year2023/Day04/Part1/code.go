package y2023d04p01

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Scratchcard struct {
	winningNumbers []int
	cardNumbers    []int
}

func (s Scratchcard) getPoints() float64 {
	matchQty := 0
	for i := range s.cardNumbers {
		if slices.Contains(s.winningNumbers, s.cardNumbers[i]) {
			matchQty++
		}
	}

	if matchQty == 0 {
		return 0
	} else {
		// ar^n. 1=1. 2=2. 3=4. 4=8. 5=16.
		return 1 * math.Pow(2, float64(matchQty-1))
	}
}

func ParseScratchcard(input string) Scratchcard {
	var scratchcard Scratchcard
	numbers := strings.Split(input, ": ")[1]
	numbersByType := strings.Split(numbers, "|")
	winners := strings.Split(strings.TrimSpace(numbersByType[0]), " ")
	for i := range winners {
		if winners[i] == "" {
			continue
		}
		noWhitespaceNumber := strings.TrimSpace(winners[i])
		winningNumber, _ := strconv.Atoi(noWhitespaceNumber)
		scratchcard.winningNumbers = append(scratchcard.winningNumbers, winningNumber)
	}
	cardNumbers := strings.Split(strings.TrimSpace(numbersByType[1]), " ")
	for i := range cardNumbers {
		if cardNumbers[i] == "" {
			continue
		}
		noWhitespaceNumber := strings.TrimSpace(cardNumbers[i])
		cardNumber, _ := strconv.Atoi(noWhitespaceNumber)
		scratchcard.cardNumbers = append(scratchcard.cardNumbers, cardNumber)
	}

	return scratchcard
}

func Calculate(input string) float64 {
	unparsedCards := strings.Split(input, "\n")
	var cards []Scratchcard

	for i := range unparsedCards {
		cards = append(cards, ParseScratchcard(unparsedCards[i]))
	}

	total := float64(0)
	for j := range cards {
		total = total + cards[j].getPoints()
	}

	return total

}
