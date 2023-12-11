package y2023d11p02

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

func GetEmptyVectors(input [][]rune) [][]int {
	//printGrid(input)
	var rows []int
	var columns []int

	var expandedGrid [][]rune
	for x, row := range input {
		expandedGrid = append(expandedGrid, row)
		if !slices.Contains(row, '#') {
			rows = append(rows, x)
		}
	}

	for y := 0; y < len(input[0]); y++ {
		columnContainsStuff := false

		for x := 0; x < len(input); x++ {
			if input[x][y] == '#' {
				columnContainsStuff = true
				break
			}
		}

		if !columnContainsStuff {
			columns = append(columns, y)
		}

	}

	return [][]int{rows, columns}
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

func GetDistance(a []int, b []int, emptyVectors [][]int) int {
	emptyRows := emptyVectors[0]
	emptyColumns := emptyVectors[1]

	// Calculate the number of extra rows to go through
	minX := int(math.Min(float64(a[0]), float64(b[0])))
	maxX := int(math.Max(float64(a[0]), float64(b[0])))
	extraRows := 0
	for _, emptyRow := range emptyRows {
		if minX < emptyRow && emptyRow < maxX {
			extraRows++
		}
	}

	// Calculate the number of extra columns to go through
	minY := int(math.Min(float64(a[1]), float64(b[1])))
	maxY := int(math.Max(float64(a[1]), float64(b[1])))
	extraCols := 0
	for _, emptyCol := range emptyColumns {
		if minY < emptyCol && emptyCol < maxY {
			extraCols++
		}
	}

	return int(math.Abs(float64(a[0]-b[0]))+math.Abs(float64(a[1]-b[1]))) + extraRows*(1000000-1) + extraCols*(1000000-1)
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
	emptyVectors := GetEmptyVectors(grid)
	coords := GetGalaxies(grid)
	fmt.Println(len(coords))
	total := 0
	distancesComputed := 0
	for a := 0; a < len(coords)-1; a++ {
		for b := a + 1; b < len(coords); b++ {
			total = total + GetDistance(coords[a], coords[b], emptyVectors)
			distancesComputed++
		}
	}

	fmt.Println(distancesComputed)
	return total
}
