package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCountMaxInGame(t *testing.T) {
	testCases := []struct {
		input    string
		expected Contents
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Contents{4, 6, 2}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Contents{1, 4, 3}},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Contents{20, 6, 13}},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Contents{14, 15, 3}},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", Contents{6, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.expected), func(t *testing.T) {
			result := CountMaxInGame(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestSumPossibleGames(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	result := SumPossibleGames(input)
	expected := 8

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPowerInGames(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	result := CalcPowerInGames(input)
	expected := 2286

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day2.txt")
	content, _ := os.ReadFile(absPath)

	result := SumPossibleGames(strings.Split(string(content), "\n"))
	expected := 2101

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day2.txt")
	content, _ := os.ReadFile(absPath)

	result := CalcPowerInGames(strings.Split(string(content), "\n"))
	expected := 58269

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
