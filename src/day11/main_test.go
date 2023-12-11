package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testInput = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 374

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{10, 1030},
		{100, 8410},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Expanding the universe by %d gives length %d", tc.input, tc.expected), func(t *testing.T) {

			result := Part2(testInput, tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day11.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 9974721

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day11.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"), 1000000)
	expected := 9974721

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
