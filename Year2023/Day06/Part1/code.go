package y2023d06p01

import (
	"strconv"
	"strings"
)

func CalcRace(time int, currentRecord int) []int {
	var winningHoldTimes []int
	for holdTime := 0; holdTime <= time; holdTime++ {
		racingTime := time - holdTime
		distance := racingTime * holdTime * 1 // you gain 1 mm/ms for each ms held
		if distance > currentRecord {
			winningHoldTimes = append(winningHoldTimes, holdTime)
		}
	}

	return winningHoldTimes
}

func Calculate(input string) int {
	rows := strings.Split(input, "\n")

	timesString, _ := strings.CutPrefix(rows[0], "Time:")
	distancesString, _ := strings.CutPrefix(rows[1], "Distance:")

	timesString = strings.Join(strings.Fields(timesString), " ")
	distancesString = strings.Join(strings.Fields(distancesString), " ")

	timeStrings := strings.Split(timesString, " ")
	distanceStrings := strings.Split(distancesString, " ")

	var numberOfWinsEachRace []int
	for i := 0; i < len(timeStrings); i++ {

		time, _ := strconv.Atoi(timeStrings[i])
		distance, _ := strconv.Atoi(distanceStrings[i])

		winningOptions := CalcRace(time, distance)
		numberOfWinsEachRace = append(numberOfWinsEachRace, len(winningOptions))
	}

	result := 1

	for _, winNumber := range numberOfWinsEachRace {
		result = result * winNumber
	}

	return result
}
