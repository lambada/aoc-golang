package y2023d05p01

import (
	"strconv"
	"strings"
)

type MapEntry struct {
	sourceStart int
	sourceEnd   int
	destStart   int
	entryRange  int
}

type AlmanacMap struct {
	source      string
	destination string
	maps        []MapEntry
}

func ParseMapEntry(input string) MapEntry {
	// 50 98 2
	inputs := strings.Split(input, " ")
	var data []int
	for _, input := range inputs {
		converted, _ := strconv.Atoi(input)
		data = append(data, converted)
	}
	return MapEntry{
		sourceStart: data[1],
		sourceEnd:   data[1] + data[2] - 1,
		destStart:   data[0],
		entryRange:  data[2],
	}
}

func ParseAlmanacMaps(input []string) map[string]AlmanacMap {
	var results = make(map[string]AlmanacMap)

	var currentMap AlmanacMap
	for _, row := range input {
		if row == "" {
			results[currentMap.source] = currentMap
			currentMap = AlmanacMap{}
			continue
		}

		if strings.Contains(row, " map:") {
			row, _ = strings.CutSuffix(row, " map:")
			stages := strings.Split(row, "-to-")
			currentMap.source = stages[0]
			currentMap.destination = stages[1]
			continue
		}

		currentMap.maps = append(currentMap.maps, ParseMapEntry(row))
	}

	// Always append the final map
	results[currentMap.source] = currentMap

	return results
}

func followTrail(almanacMaps map[string]AlmanacMap, sourceName string, sourceValue int) int {
	theMap := almanacMaps[sourceName]

	if sourceName == "location" {
		return sourceValue
	}

	for _, mapentry := range theMap.maps {
		if mapentry.sourceStart <= sourceValue && mapentry.sourceEnd >= sourceValue {
			offset := sourceValue - mapentry.sourceStart
			destVal := mapentry.destStart + offset
			return followTrail(almanacMaps, theMap.destination, destVal)
		}
	}

	return followTrail(almanacMaps, theMap.destination, sourceValue)
}

func Calculate(input string) int {
	rows := strings.Split(input, "\n")

	seedsLine := rows[0]
	seedsLine, _ = strings.CutPrefix(seedsLine, "seeds: ")

	seedStrings := strings.Split(seedsLine, " ")

	var seeds []int
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	// 0 is the seeds row. 1 is the separator line
	rows = rows[2:]

	almanacMaps := ParseAlmanacMaps(rows)

	// int cannot be nil. The closest they have is 0... so we instead use references so we can check if it's had a value assigned yet
	var lowestLocationNumber int
	for i, startingSeed := range seeds {
		destinationLocation := followTrail(almanacMaps, "seed", startingSeed)

		if i == 0 || (destinationLocation < lowestLocationNumber) {
			lowestLocationNumber = destinationLocation
		}
	}

	return lowestLocationNumber
}
