package y2023d02p01

import (
	"strings"
	"testing"
)

func TestParseDraw(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Draw
	}{
		{
			name:  "Example Game 1 Draw 1",
			input: "3 blue, 4 red",
			want:  Draw{b: 3, r: 4},
		},
		{
			name:  "Example Game 1 Draw 2",
			input: "1 red, 2 green, 6 blue",
			want:  Draw{r: 1, g: 2, b: 6},
		},
		{
			name:  "Example Game 1 Draw 3",
			input: "2 green",
			want:  Draw{g: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDraw(tt.input)
			if result != tt.want {
				t.Fatalf(`received %+v, expected %+v`, result, tt.want)
			}
		})
	}

}

func TestParseGame(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  maxBalls
	}{
		{
			name:  "Example Game 1",
			input: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  maxBalls{r: 4, g: 2, b: 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseGame(tt.input)
			if result.getMaxBalls() != tt.want {
				t.Fatalf(`received %+v, expected %+v`, result.getMaxBalls(), tt.want)
			}
		})
	}
}

func TestTotalExample(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	want := 8
	result := Calculate(input)
	if result != want {
		t.Fatalf(`received %d, expected %d`, result, want)
	}
}

func TestActualPuzzle(t *testing.T) {
	rawInput := "Game 1: 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue\nGame 2: 12 blue, 1 red, 1 green; 1 red, 12 blue, 3 green; 5 green, 1 red, 9 blue; 1 red, 7 blue, 4 green\nGame 3: 1 red; 12 blue, 15 red; 1 green, 10 red, 2 blue; 1 green, 3 red, 9 blue\nGame 4: 6 blue, 5 green; 2 blue, 6 green, 6 red; 11 blue, 5 red; 6 green, 11 red, 7 blue; 4 green, 10 red; 1 green, 7 red, 13 blue\nGame 5: 10 green, 1 red, 2 blue; 3 red, 4 green, 4 blue; 5 green, 5 red\nGame 6: 1 green, 6 blue, 14 red; 9 blue, 5 red; 14 red, 12 blue\nGame 7: 1 green, 9 red, 8 blue; 9 blue, 1 green, 6 red; 1 green, 15 blue, 19 red\nGame 8: 9 red, 7 green, 2 blue; 6 red, 17 green; 18 red, 16 green, 2 blue; 10 red, 14 green\nGame 9: 1 blue, 11 red, 9 green; 8 red, 1 blue, 9 green; 4 blue, 16 red, 9 green; 8 green, 3 blue, 6 red; 8 green, 11 red, 3 blue; 11 red, 2 blue\nGame 10: 8 green, 14 blue; 1 red, 6 blue, 9 green; 6 blue, 4 green, 1 red; 16 green, 9 blue\nGame 11: 6 green, 11 blue, 1 red; 10 green, 1 red; 7 blue, 2 green, 1 red\nGame 12: 3 green, 5 blue, 2 red; 14 blue, 16 green, 4 red; 8 green, 14 blue, 4 red\nGame 13: 5 red, 12 blue, 2 green; 2 green, 1 red, 9 blue; 1 red, 2 blue, 3 green; 3 green, 3 red, 7 blue; 2 red, 13 blue; 1 red, 10 blue, 2 green\nGame 14: 5 blue, 1 red, 2 green; 8 blue, 1 green, 1 red; 1 blue, 2 green\nGame 15: 14 blue, 9 green, 1 red; 2 red, 15 blue, 12 green; 1 blue, 2 green, 1 red; 1 red, 16 green, 15 blue; 1 red, 12 green, 8 blue; 1 red, 17 blue\nGame 16: 7 red, 1 green, 18 blue; 7 blue, 5 green, 17 red; 14 blue, 8 red, 6 green\nGame 17: 4 green, 5 blue; 5 green, 1 red, 7 blue; 3 green, 6 blue, 4 red; 2 green, 5 blue; 9 green, 6 red, 6 blue\nGame 18: 8 red, 6 blue; 4 blue, 19 red; 4 blue, 9 red; 9 blue, 10 red; 2 green, 9 blue, 13 red; 3 blue, 7 red\nGame 19: 8 green, 2 red, 17 blue; 11 blue, 4 red, 5 green; 8 blue, 8 green, 10 red; 9 green, 4 blue, 2 red; 4 green, 10 red, 6 blue\nGame 20: 9 green, 3 blue, 1 red; 5 blue, 16 green, 3 red; 3 green, 3 red; 2 blue, 1 red, 5 green\nGame 21: 7 green, 1 red, 10 blue; 7 green, 5 blue, 7 red; 7 green, 9 blue\nGame 22: 5 red, 2 blue, 9 green; 6 red, 11 green; 6 green, 6 red\nGame 23: 14 red, 2 blue, 9 green; 9 green, 1 blue, 4 red; 9 red, 1 green, 1 blue; 6 green; 3 blue, 1 green, 9 red; 1 blue, 2 red\nGame 24: 3 red, 7 green, 6 blue; 1 green, 5 blue; 6 blue, 1 red, 2 green; 5 red, 1 blue, 4 green; 6 red, 2 blue, 11 green; 2 green, 2 red, 1 blue\nGame 25: 5 green, 1 red, 3 blue; 3 blue, 6 green, 3 red; 3 red, 4 green, 1 blue; 6 green, 1 blue, 9 red; 2 blue, 2 red, 1 green\nGame 26: 3 green, 4 red, 12 blue; 2 red, 1 green, 15 blue; 7 red, 16 green, 4 blue; 11 blue, 11 green, 3 red; 8 green, 15 blue, 10 red\nGame 27: 9 red; 10 red, 2 blue; 3 red; 8 red, 1 green, 2 blue; 1 red, 2 blue; 1 blue, 4 red\nGame 28: 5 blue, 8 red, 5 green; 10 blue, 4 red, 5 green; 8 red, 14 blue, 10 green; 10 blue, 4 red, 1 green; 5 red, 17 green, 4 blue\nGame 29: 16 green, 11 red, 5 blue; 11 red, 14 blue, 13 green; 13 blue, 8 green; 3 red, 18 green, 15 blue\nGame 30: 2 red, 4 blue, 8 green; 6 green, 2 red, 2 blue; 6 green, 6 blue, 2 red\nGame 31: 2 red, 1 blue, 16 green; 10 green, 1 blue, 7 red; 1 blue, 14 green, 7 red; 2 blue, 1 green, 1 red; 6 red, 13 green; 2 blue, 6 red, 10 green\nGame 32: 4 green, 4 blue; 1 green, 5 red; 6 green, 1 red; 3 green, 5 red, 2 blue; 4 red, 1 blue, 4 green; 6 green, 2 blue, 6 red\nGame 33: 5 blue, 2 red, 1 green; 5 blue; 1 blue, 1 green, 10 red; 8 red, 3 blue, 1 green\nGame 34: 15 blue, 7 green; 12 green, 17 blue; 10 blue, 11 green; 1 red, 5 blue, 9 green; 2 red, 10 blue, 11 green\nGame 35: 2 red, 6 blue; 2 red, 5 blue, 4 green; 2 red, 8 green, 10 blue\nGame 36: 4 red, 9 green, 3 blue; 4 red, 6 green; 6 red; 11 red, 4 green\nGame 37: 3 blue, 12 green, 14 red; 3 red, 5 green, 7 blue; 2 blue, 2 green, 16 red\nGame 38: 17 blue, 16 red, 8 green; 4 green, 17 blue, 4 red; 8 red, 7 blue, 6 green; 2 blue, 9 green, 17 red; 10 blue, 8 green, 11 red\nGame 39: 10 blue, 1 red, 4 green; 4 green, 4 red, 6 blue; 11 blue\nGame 40: 5 green, 17 blue; 11 blue, 4 green, 7 red; 2 green, 6 red, 13 blue; 7 blue, 12 green, 16 red; 15 red, 3 green, 8 blue; 12 green, 3 blue, 12 red\nGame 41: 13 blue, 3 red, 1 green; 2 green, 10 red; 1 blue, 5 red, 3 green; 5 green, 16 blue; 9 blue, 2 green; 14 blue, 4 green, 5 red\nGame 42: 2 blue, 15 green, 3 red; 3 red, 17 green; 6 red, 1 blue, 8 green\nGame 43: 8 green, 9 red, 3 blue; 1 blue, 13 red; 5 red, 1 blue, 6 green; 2 red, 2 blue; 17 red, 2 blue, 6 green\nGame 44: 10 red, 3 blue; 10 blue, 5 green; 4 red, 4 blue, 1 green; 16 blue, 6 red, 7 green; 3 green, 12 blue\nGame 45: 12 blue, 2 red; 2 blue, 3 red, 2 green; 8 blue, 3 green; 4 green, 8 blue, 5 red; 3 red, 2 blue, 1 green; 1 red, 2 blue, 7 green\nGame 46: 1 blue, 11 red, 6 green; 2 blue, 11 red, 6 green; 8 red, 5 green\nGame 47: 2 blue, 9 red; 1 green, 5 blue; 10 red, 2 blue, 2 green; 10 red, 3 green, 3 blue; 3 red, 6 blue, 2 green; 1 red, 1 green, 5 blue\nGame 48: 1 red, 7 green; 1 blue, 10 green, 5 red; 4 red, 8 green; 10 red, 10 green; 2 red, 16 green; 11 red, 14 green, 1 blue\nGame 49: 1 red, 1 blue, 5 green; 6 green, 5 red; 3 blue, 4 red, 3 green; 3 red, 5 green, 2 blue; 3 blue, 3 red\nGame 50: 17 red, 1 green, 7 blue; 4 blue, 1 red, 5 green; 10 red, 13 blue; 17 red\nGame 51: 2 red, 1 green; 1 green, 10 blue, 2 red; 5 red, 1 green, 7 blue; 7 blue, 1 red; 9 blue, 5 red, 2 green\nGame 52: 8 green, 1 blue; 14 green, 1 red; 10 green, 1 red\nGame 53: 17 green, 6 blue; 17 blue, 9 green; 1 red, 12 blue\nGame 54: 4 blue, 7 red, 9 green; 7 red, 2 green; 14 green, 10 red, 3 blue; 9 green, 6 blue, 5 red; 2 blue, 3 green, 11 red\nGame 55: 11 green, 4 red; 14 green; 3 red, 3 green; 3 green, 4 red, 1 blue; 15 green, 6 red, 2 blue; 4 red, 3 blue, 15 green\nGame 56: 8 blue, 5 red, 9 green; 11 green, 5 blue, 6 red; 1 green, 1 blue, 7 red; 7 green, 4 red, 1 blue; 9 blue, 5 red, 1 green; 5 red, 2 blue\nGame 57: 11 green, 19 blue, 5 red; 15 green, 5 red, 18 blue; 16 green, 5 red, 10 blue; 19 blue, 3 red; 9 green, 8 blue\nGame 58: 4 blue, 12 green; 11 green, 4 blue; 6 blue, 6 green; 1 red, 2 green; 11 green, 3 blue; 13 blue, 6 green\nGame 59: 10 blue, 1 red; 1 green, 4 blue; 4 blue\nGame 60: 7 red, 2 green, 6 blue; 1 green, 13 red, 12 blue; 9 blue, 9 green, 8 red\nGame 61: 7 green, 3 red, 2 blue; 1 red, 1 blue; 5 green, 3 blue; 4 blue, 1 red, 4 green\nGame 62: 1 green, 8 blue, 6 red; 7 blue, 3 red, 12 green; 2 blue, 7 red, 6 green\nGame 63: 3 red, 2 green; 3 green, 4 blue, 9 red; 3 blue, 3 green, 16 red; 4 red, 1 blue\nGame 64: 10 red, 2 green, 15 blue; 4 red, 14 green; 6 red, 14 green, 2 blue\nGame 65: 11 red, 14 green, 5 blue; 7 blue, 14 green, 15 red; 1 blue, 14 green; 4 green, 4 blue, 7 red\nGame 66: 6 blue, 9 green, 6 red; 6 blue, 2 red, 4 green; 3 blue; 8 blue, 5 green, 8 red; 17 blue, 11 green; 12 green, 11 blue, 4 red\nGame 67: 8 red, 4 blue, 6 green; 4 blue, 8 red, 2 green; 1 green, 6 red, 2 blue; 10 red, 1 green, 2 blue; 1 blue, 5 red; 2 red, 1 green, 2 blue\nGame 68: 10 green, 9 red, 13 blue; 2 blue, 2 green, 4 red; 11 red, 13 blue; 4 green, 2 red, 8 blue\nGame 69: 16 red; 12 red, 1 green, 4 blue; 1 green, 14 red, 9 blue; 12 blue, 2 green, 13 red; 14 red, 2 green, 10 blue; 11 blue\nGame 70: 1 red, 19 green; 4 blue, 6 green; 12 green, 2 red\nGame 71: 9 green, 2 blue, 3 red; 5 red; 1 red, 1 blue, 5 green\nGame 72: 1 green, 19 red; 12 red, 1 green, 1 blue; 16 red, 6 blue; 14 red, 7 blue; 11 blue, 1 green, 13 red; 16 blue, 4 red\nGame 73: 1 green, 1 red, 2 blue; 8 green, 2 red, 4 blue; 7 blue, 7 green, 7 red\nGame 74: 1 blue, 1 green; 1 red; 1 red, 5 blue, 4 green; 2 blue, 1 red; 1 blue\nGame 75: 8 green, 1 red; 3 blue, 1 red, 5 green; 12 green\nGame 76: 8 green, 6 red, 2 blue; 7 red, 1 blue; 2 blue, 9 green, 1 red; 8 red, 13 green; 12 green, 2 red; 7 green, 5 red\nGame 77: 3 blue, 1 green, 10 red; 13 red; 1 blue, 13 red, 1 green\nGame 78: 2 red, 3 green, 14 blue; 3 red, 16 blue, 6 green; 3 blue, 3 red, 9 green; 4 blue, 11 green; 6 green, 2 blue; 2 red, 10 green, 11 blue\nGame 79: 5 green, 10 blue, 2 red; 16 green, 15 blue, 1 red; 1 red, 11 green; 1 red, 11 blue, 16 green; 7 blue, 18 green\nGame 80: 9 blue, 10 green; 13 green, 9 blue; 1 red, 5 green, 5 blue; 13 green, 5 blue, 1 red\nGame 81: 9 green, 15 red, 11 blue; 11 blue, 9 red, 5 green; 2 green, 11 blue, 19 red; 14 green, 15 red, 5 blue\nGame 82: 4 green, 6 red, 13 blue; 5 blue, 5 red, 4 green; 2 green, 7 blue\nGame 83: 12 blue, 8 red; 6 red, 1 blue, 1 green; 7 red, 1 green, 1 blue; 6 red, 1 green, 9 blue; 10 blue, 3 red; 1 red, 5 blue\nGame 84: 9 blue, 13 red; 8 blue, 1 green; 9 blue, 1 green; 3 blue, 5 red, 1 green\nGame 85: 14 green, 5 blue, 8 red; 1 blue, 5 green, 1 red; 10 red, 7 blue, 17 green; 3 blue, 6 green; 6 red, 5 blue, 4 green; 5 blue, 4 green, 6 red\nGame 86: 14 red; 20 red, 3 blue; 1 green, 12 blue, 15 red; 16 red, 13 blue; 13 red, 12 blue; 2 blue, 20 red\nGame 87: 2 blue, 2 red, 10 green; 8 green, 9 red, 1 blue; 11 red, 1 green, 4 blue; 13 red, 1 blue; 11 green, 16 red, 3 blue\nGame 88: 5 green, 4 red, 1 blue; 3 blue, 8 red, 10 green; 11 green, 7 red, 4 blue; 11 green, 5 blue, 4 red; 9 red, 9 green; 4 blue, 6 green, 9 red\nGame 89: 2 blue, 2 red, 5 green; 2 red, 2 blue, 3 green; 2 red, 1 blue, 7 green; 6 green, 1 red, 2 blue\nGame 90: 4 green, 1 blue, 5 red; 2 blue, 2 red, 10 green; 2 green, 8 red, 1 blue; 10 green, 5 red; 2 red, 3 green, 2 blue\nGame 91: 16 blue, 5 red, 15 green; 4 green, 7 red, 3 blue; 4 red, 8 green, 12 blue; 4 green, 8 red, 17 blue\nGame 92: 13 red, 2 blue, 12 green; 19 green, 7 red; 17 green, 2 blue, 3 red; 6 blue, 11 red, 10 green; 6 red, 15 green, 3 blue; 6 blue, 20 green, 11 red\nGame 93: 2 blue, 3 green; 1 blue, 4 red; 1 red\nGame 94: 3 red, 5 green, 6 blue; 7 blue, 5 green, 6 red; 9 blue, 1 green, 2 red; 4 blue, 1 green, 4 red\nGame 95: 8 green, 9 red, 2 blue; 7 green, 7 red; 2 green, 4 blue, 6 red; 6 blue, 2 red, 2 green\nGame 96: 11 blue, 4 red; 1 green, 3 red, 14 blue; 2 green, 3 red, 8 blue; 7 red, 1 green, 3 blue; 8 blue, 6 red, 2 green; 9 blue, 3 red, 3 green\nGame 97: 5 green, 13 red, 7 blue; 2 blue, 12 red, 6 green; 10 blue, 11 red, 3 green; 4 green, 11 blue, 15 red; 8 green, 16 blue, 1 red; 15 blue, 4 red, 5 green\nGame 98: 3 blue, 1 red; 4 blue; 2 green, 1 blue; 2 green, 1 red, 5 blue\nGame 99: 4 green, 3 blue, 9 red; 6 blue, 5 red, 3 green; 2 green, 4 blue, 7 red; 8 red, 4 blue; 2 green, 15 red; 4 red, 5 blue, 3 green\nGame 100: 8 red, 4 blue, 4 green; 10 blue, 3 red, 4 green; 10 green, 4 red; 18 red, 9 blue, 2 green; 12 red, 4 green, 2 blue"
	input := strings.Split(rawInput, "\n")
	want := 2416
	result := Calculate(input)
	if result != want {
		t.Fatalf(`received %d, expected %d`, result, want)
	}
}
