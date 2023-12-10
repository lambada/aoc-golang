package y2023d10p02

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type Cell struct {
	isPipe     bool
	isFakeCell bool
	north      bool
	south      bool
	east       bool
	west       bool
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

func getLoopCoordinates(grid [][]Cell, startingCoord []int) [][]int {
	currentCoordinate := startingCoord
	var arrivalDirection string
	loopLength := 0
	var loopCoords [][]int
	loopCoords = append(loopCoords, currentCoordinate)

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
		} else {
			loopCoords = append(loopCoords, currentCoordinate)
		}

	}

	return loopCoords
}

func coordComparator(a []int, b []int) bool {
	if len(a) == len(b) && len(a) == 2 && a[0] == b[0] && a[1] == b[1] {
		return true
	}
	return false
}

func printGrid(grid [][]Cell) {
	for _, row := range grid {
		for _, cell := range row {
			if cell.east && cell.west {
				fmt.Print("-")
				continue
			}
			if cell.north && cell.south {
				fmt.Print("|")
				continue
			}
			if cell.north && cell.east {
				fmt.Print("L")
				continue
			}
			if cell.east && cell.south {
				fmt.Print("F")
				continue
			}
			if cell.south && cell.west {
				fmt.Print("7")
				continue
			}
			if cell.west && cell.north {
				fmt.Print("J")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}

}

func Calculate(input string) int {

	grid, startingCoord := ParseGrid(input)

	//Cleanup the grid
	//1 - Rather than returning the length of the loop. Return every cell that is on the loop.
	loopCoords := getLoopCoordinates(grid, startingCoord)
	fmt.Printf("Found the loop of %d length\n", len(loopCoords))
	//2 - Mark every other cell as dirt, to remove fake pipes.
	grid = TurnFakePipesToDirt(grid, loopCoords)
	fmt.Println("Printing grid post cleanup")
	//printGrid(grid)
	fmt.Println("END OF GRID")

	//Remove the need to squeeze
	//3 - Add in an extra row between every row. Add in an extra column between every column.
	expandedGrid := expandGrid(grid)
	fmt.Println("Printing grid post expanion")
	//printGrid(expandedGrid)
	fmt.Println("END OF GRID")

	//4 - for each added cell, use the neighbouring cells to work out if it should be a |, - or . cell.
	// The only ones that can be non dirt are sandwiched between existing rows and columns, so we can check those easily.

	// Finally do the cells that are in a fake row and a fake column.
	// I think these are by definition dirt?

	//Compute the completedZones
	//5 - Start at 0,0. Follow each NSEW until you hit a non-dirt cell. Record each cells co-ordinates. That is Zone A.
	//6 - For the first dirt cell that is not in Zone A... Repeat 5, but call it zone B.
	//7 - Repeat 6 until no more un-zoned dirt cells.

	var completedZones [][][]int
	for x := 0; x < len(expandedGrid); x++ {
		for y := 0; y < len(expandedGrid[x]); y++ {
			// Ignore actual pipes
			if expandedGrid[x][y].isPipe {
				continue
			}
			// Ignore any already in a completed zone
			if slices.ContainsFunc(completedZones, func(zone [][]int) bool {
				return slices.ContainsFunc(zone, func(a []int) bool {
					if a[0] == x && a[1] == y {
						return true
					}
					return false
				})
			}) {
				continue
			}

			var processingZone [][]int

			var coordsToCheck [][]int
			coordsToCheck = append(coordsToCheck, []int{x, y})
			coordsToCheck = append(coordsToCheck, []int{x - 1, y})
			coordsToCheck = append(coordsToCheck, []int{x + 1, y})
			coordsToCheck = append(coordsToCheck, []int{x, y - 1})
			coordsToCheck = append(coordsToCheck, []int{x, y + 1})

			for i := 0; i < len(coordsToCheck); i++ {
				coord := coordsToCheck[i]

				// Skip any out of bounds
				if coord[0] >= len(expandedGrid) || coord[0] < 0 || coord[1] >= len(expandedGrid[coord[0]]) || coord[1] < 0 {
					continue
				}

				// Skip any pipe
				if expandedGrid[coord[0]][coord[1]].isPipe {
					continue
				}

				// Has this coordinate already been checked in a previous zone?
				if slices.ContainsFunc(completedZones, func(zone [][]int) bool {
					return slices.ContainsFunc(zone, func(a []int) bool {
						if a[0] == coord[0] && a[1] == coord[1] {
							return true
						}
						return false
					})
				}) {
					continue
				}

				// Has this coordinate been checked in this zone?
				if slices.ContainsFunc(processingZone, func(a []int) bool {
					if a[0] == coord[0] && a[1] == coord[1] {
						return true
					}
					return false
				}) {
					continue
				}

				processingZone = append(processingZone, []int{coord[0], coord[1]})
				coordsToCheck = append(coordsToCheck, []int{coord[0] - 1, coord[1]})
				coordsToCheck = append(coordsToCheck, []int{coord[0] + 1, coord[1]})
				coordsToCheck = append(coordsToCheck, []int{coord[0], coord[1] - 1})
				coordsToCheck = append(coordsToCheck, []int{coord[0], coord[1] + 1})
			}

			completedZones = append(completedZones, processingZone)
		}
	}
	//Find the right completedZones
	//8 - completedZones inside the loop will be strictly equivilent to completedZones that do not have any edge-cells in them (ie no [0,y] or [x,0]).
	var interiorZones [][][]int

	for _, testZone := range completedZones {

		isInterior := true
		for _, coord := range testZone {
			if coord[0] == 0 || coord[1] == 0 || coord[0] == len(expandedGrid)-1 || coord[1] == len(expandedGrid[coord[0]])-1 {
				isInterior = false
				break
			}
		}

		if isInterior {
			interiorZones = append(interiorZones, testZone)
		}

	}

	//Remove the extra cells
	//9 - From the correct completedZones, remove any whose coordinates involve a cell that is fake (ie %2 = 1 for x or y.
	var realInteriorCells [][]int

	for _, zone := range interiorZones {
		for _, coord := range zone {
			if !expandedGrid[coord[0]][coord[1]].isFakeCell {
				realInteriorCells = append(realInteriorCells, coord)
			}
		}
	}

	//10 - count the number of cells in those completedZones.

	return len(realInteriorCells)
}

func expandGrid(grid [][]Cell) [][]Cell {
	var expandedGrid [][]Cell
	for x := range grid {
		var expandedRow []Cell
		for y := range grid[x] {
			expandedRow = append(expandedRow, grid[x][y])
			if y < len(grid[x])-1 {
				expandedRow = append(expandedRow, Cell{isFakeCell: true})
			}
		}
		expandedGrid = append(expandedGrid, expandedRow)
		if x < len(grid)-1 {
			var extraRow []Cell
			for i := 0; i < len(expandedRow); i++ {
				extraRow = append(extraRow, Cell{isFakeCell: true})
			}
			expandedGrid = append(expandedGrid, extraRow)
		}
	}

	for x := 0; x < len(expandedGrid); x++ {
		for y := 0; y < len(expandedGrid[x]); y++ {
			fakeCell := expandedGrid[x][y]
			if !fakeCell.isFakeCell {
				continue
			}
			if (y-1 >= 0) && (y+1 < len(expandedGrid[x])) {
				if expandedGrid[x][y-1].east && expandedGrid[x][y+1].west {
					fakeCell.isPipe = true
					fakeCell.east = true
					fakeCell.west = true
					expandedGrid[x][y] = fakeCell
					continue
				}
			}
			if (x-1 >= 0) && (x+1 < len(expandedGrid)) {
				if expandedGrid[x-1][y].south && expandedGrid[x+1][y].north {
					fakeCell.isPipe = true
					fakeCell.north = true
					fakeCell.south = true
					expandedGrid[x][y] = fakeCell
					continue
				}
			}
		}
	}

	return expandedGrid
}

func TurnFakePipesToDirt(grid [][]Cell, loopCoords [][]int) [][]Cell {
	for x := range grid {
		for y := range grid[x] {
			if !slices.ContainsFunc(loopCoords, func(a []int) bool {
				return (a[0] == x && a[1] == y)
			}) {
				grid[x][y].isPipe = false
				grid[x][y].north = false
				grid[x][y].south = false
				grid[x][y].east = false
				grid[x][y].west = false
			}
		}
	}
	return grid
}
