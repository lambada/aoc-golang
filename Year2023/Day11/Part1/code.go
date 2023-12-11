package y2023d11p01

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func ParseGrid(input string) [][]rune {

	rows := strings.Split(input, "\n")

	var parsedGrid [][]rune
	for _, row := range rows {
		parsedGrid = append(parsedGrid, []rune(row))
	}

	return parsedGrid
}

func ExpandGrid(input [][]rune) [][]rune {
	//printGrid(input)
	var expandedGrid [][]rune
	for _, row := range input {
		expandedGrid = append(expandedGrid, row)
		if !slices.Contains(row, '#') {
			expandedGrid = append(expandedGrid, row)
		}
	}
	//printGrid(expandedGrid)

	extraColumnsAdded := 0
	for y := 0; y < len(input[0]); y++ {
		columnContainsStuff := false

		for x := 0; x < len(input); x++ {
			if input[x][y] == '#' {
				columnContainsStuff = true
				break
			}
		}

		if !columnContainsStuff {
			columnToExpand := y + extraColumnsAdded
			for x := 0; x < len(expandedGrid); x++ {
				var updatedRow []rune
				for i := 0; i < len(expandedGrid[x]); i++ {
					updatedRow = append(updatedRow, expandedGrid[x][i])
					if i == columnToExpand {
						updatedRow = append(updatedRow, '.')

					}
				}
				expandedGrid[x] = updatedRow
			}
			extraColumnsAdded++
		}

	}
	//printGrid(expandedGrid)

	return expandedGrid
}

func GetGalaxies(input [][]rune) [][]int {

	var coords [][]int
	for x := range input {
		for y := range input[x] {
			if input[x][y] == '#' {
				coords = append(coords, []int{x, y})
			}
		}
	}

	return coords

}

func GetDistance(a []int, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func printGrid(grid [][]rune) {
	fmt.Println("---------")
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Print("\n")
	}
	fmt.Println("---------")
}

func Calculate(input string) int {
	grid := ParseGrid(input)
	grid = ExpandGrid(grid)
	coords := GetGalaxies(grid)
	fmt.Println(len(coords))
	total := 0
	distancesComputed := 0
	for a := 0; a < len(coords)-1; a++ {
		for b := a + 1; b < len(coords); b++ {
			total = total + GetDistance(coords[a], coords[b])
			distancesComputed++
		}
	}

	fmt.Println(distancesComputed)
	return total
}
