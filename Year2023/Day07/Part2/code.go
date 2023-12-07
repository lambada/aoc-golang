package y2023d07p02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type HandType int32

type HandBid struct {
	Hand string
	Bid  int
}

const (
	HIGH_CARD  HandType = iota
	ONE_PAIR   HandType = iota
	TWO_PAIR   HandType = iota
	THREE_KIND HandType = iota
	FULL_HOUSE HandType = iota
	FOUR_KIND  HandType = iota
	FIVE_KIND  HandType = iota
)

func HandTypeCalculator(hand string) HandType {

	cards := strings.Split(hand, "")

	cardQtys := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	jokerCount := 0
	for _, card := range cards {
		switch card {
		case "A":
			cardQtys[11] = cardQtys[11] + 1
			break
		case "K":
			cardQtys[10] = cardQtys[10] + 1
			break
		case "Q":
			cardQtys[9] = cardQtys[9] + 1
			break
		case "J":
			jokerCount = jokerCount + 1
			break
		case "T":
			cardQtys[8] = cardQtys[8] + 1
			break
		default:
			cardVal, _ := strconv.Atoi(card)
			cardQtys[cardVal-2] = cardQtys[cardVal-2] + 1
		}
	}

	orderedQtys := cardQtys
	slices.Sort(orderedQtys)
	slices.Reverse(orderedQtys)
	orderedQtys[0] = orderedQtys[0] + jokerCount

	if orderedQtys[0] == 5 {
		return FIVE_KIND
	}

	if orderedQtys[0] == 4 {
		return FOUR_KIND
	}

	if orderedQtys[0] == 3 {
		if orderedQtys[1] == 2 {
			return FULL_HOUSE
		}
		return THREE_KIND
	}

	if orderedQtys[0] == 2 {
		if orderedQtys[1] == 2 {
			return TWO_PAIR
		}

		return ONE_PAIR
	}

	return HIGH_CARD

}

func HandSorterFunc(hand1 HandBid, hand2 HandBid) int {
	if HandTypeCalculator(hand1.Hand) > HandTypeCalculator(hand2.Hand) {
		return -1
	}
	if HandTypeCalculator(hand1.Hand) < HandTypeCalculator(hand2.Hand) {
		return 1
	}
	// Else we need to compare cards

	const cardValues = "J23456789TQKA"

	for i := 0; i < len(hand1.Hand); i++ {

		if strings.Index(cardValues, hand1.Hand[i:i+1]) > strings.Index(cardValues, hand2.Hand[i:i+1]) {
			return -1
		}
		if strings.Index(cardValues, hand1.Hand[i:i+1]) < strings.Index(cardValues, hand2.Hand[i:i+1]) {
			return 1
		}

	}
	fmt.Println("We hit a totally equal case. That shouldn't be the case? ")
	return 0
}

func Calculate(input string) int {

	rows := strings.Split(input, "\n")

	var processedHands []HandBid

	for _, row := range rows {
		info := strings.Fields(row)
		bid, _ := strconv.Atoi(info[1])
		processedHands = append(processedHands, HandBid{Bid: bid, Hand: info[0]})
	}

	slices.SortFunc(processedHands, HandSorterFunc)

	slices.Reverse(processedHands)
	total := 0
	for i := 0; i < len(processedHands); i++ {
		rank := i + 1
		total = total + (processedHands[i].Bid * rank)
	}

	return total
}
