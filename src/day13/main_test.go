package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

var testEdgeCase1 = []string{
	"..#.#......#.",
	"###.#.####.#.",
	"##.##.#..#.##",
	"##..#.####.#.",
	"...###.##.###",
	"###.##.##.##.",
	"###..##..#...",
	"###..##..##..",
	"####.#.##.#.#",
}

func TestFindReflection(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{testInput[:7], 0},
		{RotateMirror(testInput[:7]), 5},
		{testInput[8:], 4},
		{RotateMirror(testInput[8:]), 0},
		{RotateMirror(testEdgeCase1), 1},
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

func TestFindSmudgedReflection(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{testInput[:7], 3},
		{testInput[8:], 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("The reflection has value of %d", tc.expected), func(t *testing.T) {
			result := FindSmudgedReflection(tc.input)

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

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 400

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day13.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 28895

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
