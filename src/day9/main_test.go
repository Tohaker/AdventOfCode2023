package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testInput = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestDifferenceEnginer(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{testInput[0], 18},
		{testInput[1], 28},
		{testInput[2], 68},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("The next number in sequence %s is %d", tc.input, tc.expected), func(t *testing.T) {
			inputSeq := ParseSequence(tc.input)
			result := inputSeq[len(inputSeq)-1] + DifferenceEngine(inputSeq, 0)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 114

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 2

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day9.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 1696140818

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day9.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 1152

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
