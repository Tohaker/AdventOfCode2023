package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

var testInput = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestExtractPartNumbers(t *testing.T) {
	result := ExtractPartNumbers(testInput)

	if len(result) == 0 || slices.Contains(result, 114) || slices.Contains(result, 58) {
		t.Fatalf("Result contains parts it should not")
	}
}

func TestExtractGearRatios(t *testing.T) {
	result := ExtractGearRatios(testInput)

	res13 := result["13"]
	res43 := result["43"]
	res85 := result["85"]

	if !slices.Contains(res13, 467) && !slices.Contains(res13, 35) && !slices.Contains(res43, 617) && !slices.Contains(res85, 755) && !slices.Contains(res85, 598) {
		t.Fatalf("Result doesn't contain all its parts")
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 4361

	if result != expected {
		t.Fatalf("Result %d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 467835

	if result != expected {
		t.Fatalf("Result %d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day3.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 528819

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day3.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 80403602

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
