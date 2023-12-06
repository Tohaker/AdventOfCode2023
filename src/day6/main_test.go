package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestRecordBeating(t *testing.T) {
	testCases := []struct {
		time     int
		distance int
		expected int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%dms and %dmm returns %d ways to beat the record", tc.time, tc.distance, tc.expected), func(t *testing.T) {
			result := BeatRecord(tc.time, tc.distance)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 288

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 71503

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day6.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 449820

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day6.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 42250895

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
