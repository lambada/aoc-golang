package y2023d11p01

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestParseGrid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]rune
	}{
		{
			name:  "Example 1",
			input: "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....",
			want: [][]rune{
				{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseGrid(tt.input)

			for x := range tt.want {
				for y := range tt.want[x] {
					if tt.want[x][y] != result[x][y] {
						t.Fatalf("Error in %d,%d. Receieved %+v, expected %+v", x, y, result[x][y], tt.want[x][y])
					}
				}
			}
		})
	}

}
func TestExpandGrid(t *testing.T) {
	tests := []struct {
		name  string
		input [][]rune
		want  [][]rune
	}{
		{
			name: "Simple 1",
			input: [][]rune{
				{'0', '3', '6'},
				{'1', '4', '7'},
				{'2', '5', '8'},
			},
			want: [][]rune{
				{'0', '.', '3', '.', '6', '.'},
				{'0', '.', '3', '.', '6', '.'},
				{'1', '.', '4', '.', '7', '.'},
				{'1', '.', '4', '.', '7', '.'},
				{'2', '.', '5', '.', '8', '.'},
				{'2', '.', '5', '.', '8', '.'},
			},
		},
		{
			name: "Example 1",
			input: [][]rune{
				{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
			},
			want: [][]rune{
				{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
				{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExpandGrid(tt.input)

			for x := range tt.want {
				for y := range tt.want[x] {
					if tt.want[x][y] != result[x][y] {
						t.Fatalf("Error in %d,%d. Receieved %c, expected %c", x, y, result[x][y], tt.want[x][y])
					}
				}
			}
		})
	}

}
func TestGetGalaxies(t *testing.T) {
	tests := []struct {
		name  string
		input [][]rune
		want  [][]int
	}{
		{
			name: "Example 1",
			input: [][]rune{
				{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
				{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
			},
			want: [][]int{
				{0, 4},
				{1, 9},
				{2, 0},
				{5, 8},
				{6, 1},
				{7, 12},
				{10, 9},
				{11, 0}, {11, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGalaxies(tt.input)

			for i := range tt.want {
				if tt.want[i][0] != result[i][0] || tt.want[i][1] != result[i][1] {
					t.Fatalf("Co=ordinate %d incorrect. Receieved %+v expected %+v", i, result[i], tt.want[i])
				}
			}
		})
	}

}
func TestGetDistance(t *testing.T) {
	tests := []struct {
		name   string
		inputA []int
		inputB []int
		want   int
	}{
		{
			name:   "Example 1",
			inputA: []int{6, 1},
			inputB: []int{11, 5},
			want:   9,
		},
		{
			name:   "Example 2",
			inputA: []int{0, 4},
			inputB: []int{10, 9},
			want:   15,
		},
		{
			name:   "Example 3",
			inputA: []int{2, 0},
			inputB: []int{7, 12},
			want:   17,
		},
		{
			name:   "Example 4",
			inputA: []int{11, 0},
			inputB: []int{11, 5},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDistance(tt.inputA, tt.inputB)

			if result != tt.want {
				t.Fatalf("Receieved %d expected %d", result, tt.want)

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
			input: "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....",
			want:  374,
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
			want:  9686930,
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
