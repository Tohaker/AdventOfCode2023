package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testUnsolvedInput = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestValidPermutations(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{testUnsolvedInput[0], 1},
		{testUnsolvedInput[1], 4},
		{testUnsolvedInput[2], 1},
		{testUnsolvedInput[3], 1},
		{testUnsolvedInput[4], 4},
		{testUnsolvedInput[5], 10},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Find possible arrangements for %s", tc.input), func(t *testing.T) {
			result := len(FindValidPermutations(ParseInput(tc.input)))

			if result != tc.expected {
				t.Fatalf("%d does not match %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testUnsolvedInput)
	expected := 21

	if result != expected {
		t.Fatalf("%d does not match %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testUnsolvedInput)
	expected := 525152

	if result != expected {
		t.Fatalf("%d does not match %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day12.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 7753

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
