package y2023d06p01

import (
	"testing"
)

func TestCalcRace(t *testing.T) {
	tests := []struct {
		name     string
		duration int
		distance int
		want     []int
	}{
		{
			name:     "Example Race 1",
			duration: 7,
			distance: 9,
			want:     []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalcRace(tt.duration, tt.distance)

			if len(result) != len(tt.want) {
				t.Fatalf("Incorrect number of wins. Received %+v, expected %+v", result, tt.want)
			}
			for i := range tt.want {
				if tt.want[i] != result[i] {
					t.Fatalf("Receieved %d, expected %d", result, tt.want)
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
			name:  "Example",
			input: "Time:      7  15   30\nDistance:  9  40  200",
			want:  288,
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
			input: "Time:        51     69     98     78\nDistance:   377   1171   1224   1505\n",
			want:  131376,
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
