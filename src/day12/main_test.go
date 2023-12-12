package main

import (
	"fmt"
	"testing"
)

var testSolvedInput = []string{
	"#.#.### 1,1,3",
	".#...#....###. 1,1,3",
	".#.###.#.###### 1,3,1,6",
	"####.#...#... 4,1,1",
	"#....######..#####. 1,6,5",
	".###.##....# 3,2,1",
}

var testUnsolvedInput = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestCreateRegex(t *testing.T) {
	for _, tc := range testSolvedInput {
		t.Run(fmt.Sprintf("Can create regex for %s", tc), func(t *testing.T) {
			parsed := ParseInput(tc)

			result := CreateRegexp(parsed.springs)

			if result.Match([]byte(parsed.damaged)) {
				t.Fatalf("%s does not match %s", result.String(), parsed.damaged)
			}
		})
	}
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
			result := len(CreateValidPermutations(ParseInput(tc.input)))

			if result != tc.expected {
				t.Fatalf("%d does not match %d", result, tc.expected)
			}
		})
	}
}
