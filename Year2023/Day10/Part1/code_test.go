package y2023d10p01

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

func TestCalculateExample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: ".....\n.S-7.\n.|.|.\n.L-J.\n.....",
			want:  4,
		},
		{
			name:  "Example 2",
			input: "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...",
			want:  8,
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
			want:  6870,
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
