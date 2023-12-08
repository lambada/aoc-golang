package y2023d08p01

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestParseNode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantName    string
		wantOptions Options
	}{
		{
			name:        "Example 1",
			input:       "AAA = (BBB, CCC)",
			wantName:    "AAA",
			wantOptions: Options{L: "BBB", R: "CCC"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeName, nodeOptions := ParseNode(tt.input)

			if nodeName != tt.wantName {
				t.Fatalf("Receieved %s, expected %s", nodeName, tt.wantName)
			}
			if nodeOptions != tt.wantOptions {
				t.Fatalf("Receieved %+v, expected %+v", nodeOptions, tt.wantOptions)
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
			input: "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)",
			want:  2,
		},
		{
			name:  "Example 2",
			input: "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)",
			want:  6,
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
			want:  14257,
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
