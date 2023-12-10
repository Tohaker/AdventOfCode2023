package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testInput1 = []string{
	"-L|F7",
	"7S-7|",
	"L|7||",
	"-L-J|",
	"L|-JF",
}

var testInput2 = []string{
	"7-F7-",
	".FJ|7",
	"SJLL7",
	"|F--J",
	"LJ.LJ",
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{testInput1, 4},
		{testInput2, 8},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("The furthest point is %d away in test %d", tc.expected, i+1), func(t *testing.T) {
			result := Part1(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

var testInput3 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

var testInput4 = []string{
	"FF7FSF7F7F7F7F7F---7",
	"L|LJ||||||||||||F--J",
	"FL-7LJLJ||||||LJL-77",
	"F--JF--7||LJLJ7F7FJ-",
	"L---JF-JLJ.||-FJLJJ7",
	"|F|F-JF---7F7-L7L|7|",
	"|FFJF7L7F-JF7|JL---7",
	"7-L-JL7||F7|L7F-7F7|",
	"L.L7LFJ|||||FJL7||LJ",
	"L7JLJL-JLJLJL--JLJ.L",
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{testInput3, 8},
		{testInput4, 10},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("There are %d enclosed tiles of ground in test %d", tc.expected, i+1), func(t *testing.T) {
			result := Part2(tc.input)

			if result != tc.expected {
				t.Fatalf("%d is not %d", result, tc.expected)
			}
		})
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day10.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 7063

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day10.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 589

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
