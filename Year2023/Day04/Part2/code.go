package y2023d04p02

import (
	"slices"
	"strconv"
	"strings"
)

type Scratchcard struct {
	winningNumbers []int
	cardNumbers    []int
	qty            int
}

func (s Scratchcard) getWinningQty() int {
	matchQty := 0
	for i := range s.cardNumbers {
		if slices.Contains(s.winningNumbers, s.cardNumbers[i]) {
			matchQty++
		}
	}
	return matchQty
}

func ParseScratchcard(input string) Scratchcard {
	var scratchcard Scratchcard
	scratchcard.qty = 1
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

func Calculate(input string) int {
	unparsedCards := strings.Split(input, "\n")
	var cards []Scratchcard

	for i := range unparsedCards {
		cards = append(cards, ParseScratchcard(unparsedCards[i]))
	}

	for j := range cards {

		// We need to distribute to this number of cards
		cardsToUpdate := cards[j].getWinningQty()

		for k := 1; k <= cardsToUpdate; k++ {
			cards[j+k].qty = cards[j+k].qty + cards[j].qty
		}

	}

	totalCards := 0
	for m := range cards {
		totalCards = totalCards + cards[m].qty
	}

	return totalCards

}
