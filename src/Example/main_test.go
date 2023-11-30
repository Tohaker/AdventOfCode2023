package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFloorMovement(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.expected), func(t *testing.T) {
			result := MoveBetweenFloors(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	absPath, _ := filepath.Abs("./input.txt")
	result := Part1(absPath)

	if result != 74 {
		t.Fatalf("%d is not 74", result)
	}
}

func TestFirstBasement(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.expected), func(t *testing.T) {
			result := FindFirstTimeInBasement(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	absPath, _ := filepath.Abs("./input.txt")
	result := Part2(absPath)

	if result != 1795 {
		t.Fatalf("%d is not 1795", result)
	}
}
