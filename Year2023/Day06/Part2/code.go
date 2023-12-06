package y2023d06p02

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

	timesString = strings.Join(strings.Fields(timesString), "")
	distancesString = strings.Join(strings.Fields(distancesString), "")
	time, _ := strconv.Atoi(timesString)
	distance, _ := strconv.Atoi(distancesString)

	winningOptions := CalcRace(time, distance)

	return len(winningOptions)
}
