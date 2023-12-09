package y2023d09p02

import (
	_ "embed"
	"slices"
	"testing"
)

//go:embed input.txt
var input string

func TestParseRow(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "Example Row 1",
			input: "0 3 6 9 12 15",
			want:  []int{0, 3, 6, 9, 12, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseRow(tt.input)

			if !slices.Equal(result, tt.want) {
				t.Fatalf("Receieved %+v, expected %+v", result, tt.want)
			}
		})
	}

}

func TestFindNextValue(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1 Intermediate Row 3",
			input: []int{0, 0, 0, 0},
			want:  0,
		},
		{
			name:  "Example 1 Intermediate Row 2",
			input: []int{3, 3, 3, 3, 3},
			want:  3,
		},
		{
			name:  "Example 1 Row 1",
			input: []int{0, 3, 6, 9, 12, 15},
			want:  -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindPreviousValue(tt.input)

			if result != tt.want {
				t.Fatalf("Receieved %d, expected %d", result, tt.want)
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
			input: "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45",
			want:  2,
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
			want:  803,
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
