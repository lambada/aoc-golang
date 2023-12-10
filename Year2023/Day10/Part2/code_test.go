package y2023d10p02

import (
	_ "embed"
	"slices"
	"testing"
)

//go:embed input.txt
var input string

func TestParseGrid(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantGrid  [][]Cell
		wantCoord []int
	}{
		{
			name:  "Example Row 1",
			input: ".....\n.S-7.\n.|.|.\n.L-J.\n.....",
			wantGrid: [][]Cell{
				{Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}},
				{Cell{isPipe: false}, Cell{isPipe: true, east: true, south: true}, Cell{isPipe: true, east: true, west: true}, Cell{isPipe: true, west: true, south: true}, Cell{isPipe: false}},
				{Cell{isPipe: false}, Cell{isPipe: true, north: true, south: true}, Cell{isPipe: false}, Cell{isPipe: true, north: true, south: true}, Cell{isPipe: false}},
				{Cell{isPipe: false}, Cell{isPipe: true, north: true, east: true}, Cell{isPipe: true, west: true, east: true}, Cell{isPipe: true, west: true, north: true}, Cell{isPipe: false}},
				{Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}, Cell{isPipe: false}},
			},
			wantCoord: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, coord := ParseGrid(tt.input)

			for x := range tt.wantGrid {
				for y := range tt.wantGrid[x] {
					if tt.wantGrid[x][y] != result[x][y] {
						t.Fatalf("Error in %d,%d. Receieved %+v, expected %+v", x, y, result[x][y], tt.wantGrid[x][y])
					}
				}
			}
			if !slices.Equal(tt.wantCoord, coord) {
				t.Fatalf("Receieved %+v, expected %+v", coord, tt.wantCoord)
			}
		})
	}

}

func TestTurnFakePipesToDirt(t *testing.T) {
	tests := []struct {
		name        string
		inputGrid   [][]Cell
		inputCoords [][]int
		want        [][]Cell
	}{
		{
			name: "Example",
			inputGrid: [][]Cell{
				{{isPipe: true, east: true, west: true}, {isPipe: true, east: true, north: true}, {isPipe: true, north: true, south: true}, {isPipe: true, south: true, east: true}, {isPipe: true, west: true, south: true}},
				{{isPipe: true, south: true, west: true}, {isPipe: true, east: true, south: true}, {isPipe: true, east: true, west: true}, {isPipe: true, south: true, west: true}, {isPipe: true, north: true, south: true}},
				{{isPipe: true, east: true, north: true}, {isPipe: true, south: true, north: true}, {isPipe: true, west: true, south: true}, {isPipe: true, south: true, north: true}, {isPipe: true, north: true, south: true}},
				{{isPipe: true, east: true, west: true}, {isPipe: true, east: true, north: true}, {isPipe: true, east: true, west: true}, {isPipe: true, north: true, west: true}, {isPipe: true, north: true, south: true}},
				{{isPipe: true, north: true, east: true}, {isPipe: true, south: true, north: true}, {isPipe: true, east: true, west: true}, {isPipe: true, north: true, west: true}, {isPipe: true, east: true, south: true}},
			},
			inputCoords: [][]int{
				{1, 1}, {1, 2}, {1, 3},
				{2, 1}, {2, 3},
				{3, 1}, {3, 2}, {3, 3},
			},
			want: [][]Cell{
				{{isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, east: true, south: true}, {isPipe: true, east: true, west: true}, {isPipe: true, south: true, west: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, south: true, north: true}, {isPipe: false}, {isPipe: true, south: true, north: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, east: true, north: true}, {isPipe: true, east: true, west: true}, {isPipe: true, north: true, west: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TurnFakePipesToDirt(tt.inputGrid, tt.inputCoords)

			for x := range tt.want {
				for y := range tt.want[x] {
					if result[x][y] != tt.want[x][y] {
						t.Fatalf("Receieved bad cell at %d, %d. Received %+v, expected %+v", x, y, result[x][y], tt.want[x][y])
					}
				}
			}
		})
	}
}

func TestExpandGrid(t *testing.T) {
	tests := []struct {
		name        string
		inputGrid   [][]Cell
		inputCoords [][]int
		want        [][]Cell
	}{
		{
			name: "Example",
			inputGrid: [][]Cell{
				{{isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, east: true, south: true}, {isPipe: true, east: true, west: true}, {isPipe: true, south: true, west: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, south: true, north: true}, {isPipe: false}, {isPipe: true, south: true, north: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: true, east: true, north: true}, {isPipe: true, east: true, west: true}, {isPipe: true, north: true, west: true}, {isPipe: false}},
				{{isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}, {isPipe: false}},
			},
			want: [][]Cell{
				{{isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}},
				{{isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}},
				{{isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: true, east: true, south: true}, {isFakeCell: true, isPipe: true, east: true, west: true}, {isPipe: true, east: true, west: true}, {isFakeCell: true, isPipe: true, east: true, west: true}, {isPipe: true, south: true, west: true}, {isFakeCell: true, isPipe: false}, {isPipe: false}},
				{{isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: true, north: true, south: true}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: true, north: true, south: true}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}},
				{{isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: true, south: true, north: true}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: true, south: true, north: true}, {isFakeCell: true, isPipe: false}, {isPipe: false}},
				{{isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: true, north: true, south: true}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: true, north: true, south: true}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}},
				{{isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: true, east: true, north: true}, {isFakeCell: true, isPipe: true, east: true, west: true}, {isPipe: true, east: true, west: true}, {isFakeCell: true, isPipe: true, east: true, west: true}, {isPipe: true, north: true, west: true}, {isFakeCell: true, isPipe: false}, {isPipe: false}},
				{{isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}, {isFakeCell: true, isPipe: false}},
				{{isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}, {isFakeCell: true, isPipe: false}, {isPipe: false}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := expandGrid(tt.inputGrid)

			for x := range tt.want {
				for y := range tt.want[x] {
					if result[x][y] != tt.want[x][y] {
						t.Fatalf("Receieved bad cell at %d, %d. Received %+v, expected %+v", x, y, result[x][y], tt.want[x][y])
					}
				}
			}
		})
	}
}

func TestCalculateExample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: ".....\n.S-7.\n.|.|.\n.L-J.\n.....",
			want:  1,
		},
		{
			name:  "Example 2",
			input: "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...",
			want:  1,
		},
		{
			name:  "Final example",
			input: "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L",
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Calculate(tt.input)

			if result != tt.want {
				t.Fatalf("Receieved %d, expected %d", result, tt.want)
			}
		})
	}

}
func TestCalculateInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Input",
			input: input,
			want:  287,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Calculate(tt.input)

			if result != tt.want {
				t.Fatalf("Receieved %d, expected %d", result, tt.want)
			}
		})
	}

}
