package y2023d03p01

import (
	strutils "github.com/torden/go-strutil"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func parseGrid(input string) [][]string {
	var grid [][]string

	unparsedRows := strings.Split(input, "\n")

	for i := range unparsedRows {
		unparsedRow := unparsedRows[i]
		entries := strings.Split(unparsedRow, "")
		grid = append(grid, entries)
	}

	return grid
}

func findPartNumbers(grid [][]string) []int {

	symbolRegex := regexp.MustCompile("[^0-9.]")

	var partNumbers []int
	var processedPartLocations []coordinate

	for x := range grid {
		for y := range grid[x] {
			symbol := grid[x][y]
			isSymbol := symbolRegex.Match([]byte(symbol))

			if isSymbol {
				// Check east
				result := checkEast(grid, x, y)

				if result.partFound && !slices.Contains(processedPartLocations, result.startCoordinate) {
					partNumbers = append(partNumbers, result.partNumber)
					processedPartLocations = append(processedPartLocations, result.startCoordinate)
				}

				// Check west
				result = checkWest(grid, x, y)

				if result.partFound && !slices.Contains(processedPartLocations, result.startCoordinate) {
					partNumbers = append(partNumbers, result.partNumber)
					processedPartLocations = append(processedPartLocations, result.startCoordinate)
				}

				results := checkNorth(grid, x, y)

				for i := range results {
					if !slices.Contains(processedPartLocations, results[i].startCoordinate) {
						partNumbers = append(partNumbers, results[i].partNumber)
						processedPartLocations = append(processedPartLocations, results[i].startCoordinate)
					}

				}

				results = checkSouth(grid, x, y)

				for i := range results {
					if !slices.Contains(processedPartLocations, results[i].startCoordinate) {
						partNumbers = append(partNumbers, results[i].partNumber)
						processedPartLocations = append(processedPartLocations, results[i].startCoordinate)
					}

				}

			}

		}
	}
	return partNumbers
}

type checkResult struct {
	partFound        bool
	partNumber       int
	partNumberString string
	startCoordinate  coordinate
}

type coordinate struct {
	x int
	y int
}

func checkEast(grid [][]string, x int, y int) checkResult {
	digitRegex := regexp.MustCompile("[0-9]")

	partNumber := ""

	offset := y + 1

	if offset >= len(grid[x]) {
		return checkResult{partFound: false}
	}

	startCoord := coordinate{x: x, y: offset}

	for isDigit := digitRegex.Match([]byte(grid[x][offset])); isDigit; isDigit = digitRegex.Match([]byte(grid[x][offset])) {
		partNumber = partNumber + grid[x][offset]
		offset = offset + 1
		if offset >= len(grid[x]) {
			break
		}
	}

	if partNumber != "" {
		partNumberInt, _ := strconv.Atoi(partNumber)
		return checkResult{partFound: true, partNumber: partNumberInt, partNumberString: partNumber, startCoordinate: startCoord}
	}
	return checkResult{partFound: false}
}
func checkWest(grid [][]string, x int, y int) checkResult {
	digitRegex := regexp.MustCompile("[0-9]")
	strproc := strutils.NewStringProc()

	partNumber := ""

	offset := y - 1

	if offset < 0 {
		return checkResult{partFound: false}
	}

	startCoord := coordinate{x: x, y: offset}

	for isDigit := digitRegex.Match([]byte(grid[x][offset])); isDigit; isDigit = digitRegex.Match([]byte(grid[x][offset])) {
		partNumber = partNumber + grid[x][offset]
		offset = offset - 1
		if offset < 0 {
			break
		}
		startCoord = coordinate{x: x, y: offset}
	}

	if partNumber != "" {
		partNumber = strproc.ReverseStr(partNumber)
		partNumberInt, _ := strconv.Atoi(partNumber)
		return checkResult{partFound: true, partNumber: partNumberInt, partNumberString: partNumber, startCoordinate: startCoord}
	}
	return checkResult{partFound: false}
}
func checkNorth(grid [][]string, x int, y int) []checkResult {
	digitRegex := regexp.MustCompile("[0-9]")

	// The approach here is to cheat. Treat the cell north of the real symbol as a fake symbol.
	// check west and east of it.
	// Then check the northern cell itself. If it's a digit, stick the results together.
	// then return all the found part numbers.

	if x-1 < 0 {
		return []checkResult{}
	}

	westResult := checkWest(grid, x-1, y)
	eastResult := checkEast(grid, x-1, y)
	isNorthADigit := digitRegex.Match([]byte(grid[x-1][y]))

	if isNorthADigit {
		fullPartNumber := ""
		var startCoord coordinate
		if westResult.partFound {
			fullPartNumber = fullPartNumber + strconv.Itoa(westResult.partNumber)
			startCoord = westResult.startCoordinate
		} else {
			startCoord = coordinate{x: x - 1, y: y}
		}

		fullPartNumber = fullPartNumber + grid[x-1][y]

		if eastResult.partFound {
			// We specifically need to use strings here because East may have a leading 0 as its most significant digit. Which becomes a problem if I'm stitching with north
			fullPartNumber = fullPartNumber + eastResult.partNumberString
		}

		partNumber, _ := strconv.Atoi(fullPartNumber)
		return []checkResult{{partFound: true, partNumber: partNumber, partNumberString: fullPartNumber, startCoordinate: startCoord}}
	}

	var results []checkResult

	if westResult.partFound {
		results = append(results, westResult)
	}
	if eastResult.partFound {
		results = append(results, eastResult)
	}
	return results
}
func checkSouth(grid [][]string, x int, y int) []checkResult {
	digitRegex := regexp.MustCompile("[0-9]")

	// The approach here is to cheat. Treat the cell north of the real symbol as a fake symbol.
	// check west and east of it.
	// Then check the northern cell itself. If it's a digit, stick the results together.
	// then return all the found part numbers.

	if x+1 >= len(grid) {
		return []checkResult{}
	}
	westResult := checkWest(grid, x+1, y)
	eastResult := checkEast(grid, x+1, y)
	isNorthADigit := digitRegex.Match([]byte(grid[x+1][y]))

	if isNorthADigit {
		fullPartNumber := ""
		var startCoord coordinate
		if westResult.partFound {
			fullPartNumber = fullPartNumber + strconv.Itoa(westResult.partNumber)
			startCoord = westResult.startCoordinate
		} else {
			startCoord = coordinate{x: x + 1, y: y}
		}

		fullPartNumber = fullPartNumber + grid[x+1][y]

		if eastResult.partFound {
			// Use strings here because east may have a leading 0 as its most significant digit which is actually important
			fullPartNumber = fullPartNumber + eastResult.partNumberString
		}

		partNumber, _ := strconv.Atoi(fullPartNumber)
		return []checkResult{{partFound: true, partNumber: partNumber, partNumberString: fullPartNumber, startCoordinate: startCoord}}
	}

	var results []checkResult

	if westResult.partFound {
		results = append(results, westResult)
	}
	if eastResult.partFound {
		results = append(results, eastResult)
	}
	return results
}

func Calculate(input string) int {
	grid := parseGrid(input)
	partNumbers := findPartNumbers(grid)
	total := 0
	for i := range partNumbers {
		total = total + partNumbers[i]
	}
	return total
}
