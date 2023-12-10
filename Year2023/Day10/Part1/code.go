package y2023d10p01

import (
	"regexp"
	"slices"
	"strings"
)

type Cell struct {
	isPipe bool
	north  bool
	south  bool
	east   bool
	west   bool
}

func ParseGrid(input string) ([][]Cell, []int) {
	var startingCoordinate []int

	var grid [][]Cell
	rows := strings.Split(input, "\n")

	for x := range rows {
		var rowCells []Cell

		columns := strings.Split(rows[x], "")
		for y := range columns {
			if columns[y] == "." {
				rowCells = append(rowCells, Cell{isPipe: false})
				continue
			}
			if columns[y] == "-" {
				rowCells = append(rowCells, Cell{isPipe: true, east: true, west: true})
				continue
			}
			if columns[y] == "|" {
				rowCells = append(rowCells, Cell{isPipe: true, north: true, south: true})
				continue
			}
			if columns[y] == "L" {
				rowCells = append(rowCells, Cell{isPipe: true, north: true, east: true})
				continue
			}
			if columns[y] == "7" {
				rowCells = append(rowCells, Cell{isPipe: true, west: true, south: true})
				continue
			}
			if columns[y] == "J" {
				rowCells = append(rowCells, Cell{isPipe: true, north: true, west: true})
				continue
			}
			if columns[y] == "F" {
				rowCells = append(rowCells, Cell{isPipe: true, east: true, south: true})
				continue
			}
			if columns[y] == "S" {
				startingCoordinate = []int{x, y}
				cell := Cell{isPipe: true}
				if x-1 >= 0 {
					northRegex := regexp.MustCompile("[|7F]")
					isNorth := northRegex.Match([]byte(rows[x-1][y : y+1]))
					cell.north = isNorth
				}
				if x+1 < len(rows) {
					southRegex := regexp.MustCompile("[|LJ]")
					isSouth := southRegex.Match([]byte(rows[x+1][y : y+1]))
					cell.south = isSouth
				}
				if y+1 < len(rows[x]) {
					eastRegex := regexp.MustCompile("[-7J]")
					isEast := eastRegex.Match([]byte{rows[x : x+1][0][y+1]})
					cell.east = isEast
				}
				if y-1 >= 0 {
					westRegex := regexp.MustCompile("[-LF]")
					isWest := westRegex.Match([]byte{rows[x : x+1][0][y-1]})
					cell.west = isWest
				}
				rowCells = append(rowCells, cell)
				continue
			}
		}
		grid = append(grid, rowCells)
	}

	return grid, startingCoordinate
}

func Calculate(input string) int {

	grid, startingCoord := ParseGrid(input)

	currentCoordinate := startingCoord
	var arrivalDirection string
	loopLength := 0
	for loopFound := false; !loopFound; loopLength++ {
		cell := grid[currentCoordinate[0]][currentCoordinate[1]]

		// work out where to go next
		if cell.north && arrivalDirection != "north" {
			arrivalDirection = "south"
			currentCoordinate = []int{currentCoordinate[0] - 1, currentCoordinate[1]}
		} else if cell.east && arrivalDirection != "east" {
			arrivalDirection = "west"
			currentCoordinate = []int{currentCoordinate[0], currentCoordinate[1] + 1}
		} else if cell.south && arrivalDirection != "south" {
			arrivalDirection = "north"
			currentCoordinate = []int{currentCoordinate[0] + 1, currentCoordinate[1]}
		} else if cell.west && arrivalDirection != "west" {
			arrivalDirection = "east"
			currentCoordinate = []int{currentCoordinate[0], currentCoordinate[1] - 1}
		}

		if slices.Equal(currentCoordinate, startingCoord) {
			loopFound = true
		}

	}

	if loopLength%2 == 0 {
		return loopLength / 2
	} else {
		return loopLength + 1/2
	}
}
