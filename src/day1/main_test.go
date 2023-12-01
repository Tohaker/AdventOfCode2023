package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.expected), func(t *testing.T) {
			result := ExtractNumber(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestReplaceStringNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"two1nine", "t2o1n9e"},
		{"eightwothree", "e8t2ot3e"},
		{"xtwone3four", "xt2o1e3f4r"},
		{"4nineeightseven2", "4n9ee8ts7n2"},
		{"zoneight234", "zo1e8t234"},
		{"7pqrstsixteen", "7pqrsts6xteen"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %s", tc.input, tc.expected), func(t *testing.T) {
			result := ReplaceStringNumber(tc.input)

			if result != tc.expected {
				t.Fatalf("%s is not %s", result, tc.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := Part1(input)
	expected := 142

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	result := Part2(input)
	expected := 281

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day1.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 53194

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day1.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 54249

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
