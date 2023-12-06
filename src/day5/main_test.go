package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var testInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestParsing(t *testing.T) {
	result := ParseAlmanac(testInput)
	expected := 4

	if len(result.seeds) != expected {
		t.Fatalf("%d is not %d", len(result.seeds), expected)
	}
}

func TestSoilNumber(t *testing.T) {
	almanac := AlmanacMap{
		{
			source:      98,
			destination: 50,
			_range:      2,
		},
		{
			source:      50,
			destination: 52,
			_range:      48,
		},
	}

	testCases := []struct {
		input    int
		expected int
	}{
		{79, 81},
		{14, 14},
		{55, 57},
		{13, 13},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d returns %d", tc.input, tc.expected), func(t *testing.T) {
			result := ProcessSeed(tc.input, almanac)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 35

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 46

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day5.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(string(content))
	expected := 218513636

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

// func TestPart2Real(t *testing.T) {
// 	absPath, _ := filepath.Abs("../../inputs/day5.txt")
// 	content, _ := os.ReadFile(absPath)

// 	result := Part2(string(content))
// 	expected := 81956384

// 	if result != expected {
// 		t.Fatalf("%d is not %d", result, expected)
// 	}
// }
