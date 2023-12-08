package main

import (
	"os"
	"path/filepath"
	"testing"
)

var testInput1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var testInput2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var testInput3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart1(t *testing.T) {
	result := Part1(testInput1)
	expected := 2

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}

	result2 := Part1(testInput2)
	expected2 := 6

	if result2 != expected2 {
		t.Fatalf("%d is not %d", result2, expected2)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput3)
	expected := 6

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day8.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(string(content))
	expected := 16897

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day8.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(string(content))
	expected := 16563603485021

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
