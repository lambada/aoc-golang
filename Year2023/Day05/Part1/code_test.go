package y2023d05p01

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func TestParseMapEntry(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  MapEntry
	}{
		{
			name:  "Example Row 1",
			input: "50 98 2",
			want: MapEntry{
				sourceStart: 98,
				destStart:   50,
				entryRange:  2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseMapEntry(tt.input)

			if result != tt.want {
				t.Fatalf("Receieved %v, expected %v", result, tt.want)
			}

		})
	}

}

func TestParseAlmanacMaps(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string]AlmanacMap
	}{
		{
			name:  "new map definition",
			input: []string{"seed-to-soil map:"},
			want: map[string]AlmanacMap{
				"seed": {
					source:      "seed",
					destination: "soil",
				},
			},
		},
		{
			name:  "entry definition",
			input: []string{"50 98 2"},
			want: map[string]AlmanacMap{
				"": {
					source:      "",
					destination: "",
					maps: []MapEntry{
						{
							sourceStart: 98,
							destStart:   50,
							entryRange:  2,
						},
					},
				},
			},
		},
		{
			name:  "Basic Map",
			input: []string{"seed-to-soil map:", "50 98 2", "52 50 48"},
			want: map[string]AlmanacMap{
				"seed": {
					source:      "seed",
					destination: "soil",
					maps: []MapEntry{
						{
							sourceStart: 98,
							destStart:   50,
							entryRange:  2,
						},
						{
							sourceStart: 50,
							destStart:   52,
							entryRange:  48,
						},
					},
				},
			},
		}, {
			name:  "Two Maps",
			input: []string{"seed-to-soil map:", "50 98 2", "52 50 48", "", "soil-to-fertilizer map:", "0 15 37", "37 52 2", "39 0 15"},
			want: map[string]AlmanacMap{
				"seed": {
					source:      "seed",
					destination: "soil",
					maps: []MapEntry{
						{
							sourceStart: 98,
							destStart:   50,
							entryRange:  2,
						}, {
							sourceStart: 50,
							destStart:   52,
							entryRange:  48,
						},
					},
				},
				"soil": {
					source:      "soil",
					destination: "fertilizer",
					maps: []MapEntry{
						{
							sourceStart: 15,
							destStart:   0,
							entryRange:  37,
						},
						{
							sourceStart: 52,
							destStart:   37,
							entryRange:  2,
						},
						{
							sourceStart: 0,
							destStart:   39,
							entryRange:  15,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseAlmanacMaps(tt.input)

			if len(result) != len(tt.want) {
				t.Fatalf("Incorrect number of AlmanacMaps. Received %+v, expected %+v", result, tt.want)
			}
			for i := range tt.want {
				if result[i].source != tt.want[i].source {
					t.Fatalf("Incorrect Source. Received %s, expected %s", result[i].source, tt.want[i].source)
				}
				if result[i].destination != tt.want[i].destination {
					t.Fatalf("Incorrect Destination. Received %s, expected %s", result[i].destination, tt.want[i].destination)
				}

				if result[i].source != tt.want[i].source {
					t.Fatalf("Incorrect Source. Received %s, expected %s", result[i].source, tt.want[i].source)
				}
				if len(result[i].maps) != len(tt.want[i].maps) {
					t.Fatalf("Incorrect number of MapEntries. Received %+v, expected %+v", result[i].maps, tt.want[i].maps)
				}

				for j := range tt.want[i].maps {
					if tt.want[i].maps[j] != result[i].maps[j] {
						t.Fatalf("Incorrect MapEntry. Received %+v, expected %+v", result[i].maps[j], tt.want[i].maps[j])
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
			name:  "Example",
			input: example,
			want:  35,
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
			name:  "Example",
			input: input,
			want:  51752125,
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
