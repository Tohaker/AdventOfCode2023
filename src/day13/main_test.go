package main

import (
	"fmt"
	"testing"
)

var testInput = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestFindReflection(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{testInput[:7], 5},
		{testInput[8:], 400},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("The reflection has value of %d", tc.expected), func(t *testing.T) {
			result := FindReflection(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 405

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
