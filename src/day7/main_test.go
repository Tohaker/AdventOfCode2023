package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testInput = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestParseHand(t *testing.T) {
	testCases := []struct {
		input    string
		expected Hand
	}{
		{testInput[0], Hand{cards: []string{"3", "2", "T", "3", "K"}, _type: "one", bid: 765}},
		{testInput[1], Hand{cards: []string{"T", "5", "5", "J", "5"}, _type: "three", bid: 684}},
		{testInput[2], Hand{cards: []string{"K", "K", "6", "7", "7"}, _type: "two", bid: 28}},
		{testInput[3], Hand{cards: []string{"K", "T", "J", "J", "T"}, _type: "two", bid: 220}},
		{testInput[4], Hand{cards: []string{"Q", "Q", "Q", "J", "A"}, _type: "three", bid: 483}},
		{"AAAAA 123", Hand{cards: []string{"A", "A", "A", "A", "A"}, _type: "five", bid: 123}},
		{"AAAAK 123", Hand{cards: []string{"A", "A", "A", "A", "K"}, _type: "four", bid: 123}},
		{"AKQJT 123", Hand{cards: []string{"A", "K", "Q", "J", "T"}, _type: "high", bid: 123}},
		{"33322 123", Hand{cards: []string{"3", "3", "3", "2", "2"}, _type: "full", bid: 123}},
	}

	for _, tc := range testCases {
		t.Run("Returns correct", func(t *testing.T) {
			result := ParseHand(tc.input)

			if len(result.cards) != len(tc.expected.cards) {
				t.Fatalf("%d is not %d", len(result.cards), len(tc.expected.cards))
			}

			if result._type != tc.expected._type {
				t.Fatalf("%s is not %s", result._type, tc.expected._type)
			}

			if result.bid != tc.expected.bid {
				t.Fatalf("%d is not %d", result.bid, tc.expected.bid)
			}
		})
	}
}

func TestJokerify(t *testing.T) {
	testCases := []struct {
		input    string
		expected Hand
	}{
		{testInput[0], Hand{cards: []string{"3", "2", "T", "3", "K"}, _type: "one", bid: 765}},
		{testInput[1], Hand{cards: []string{"T", "5", "5", "J", "5"}, _type: "four", bid: 684}},
		{testInput[2], Hand{cards: []string{"K", "K", "6", "7", "7"}, _type: "two", bid: 28}},
		{testInput[3], Hand{cards: []string{"K", "T", "J", "J", "T"}, _type: "four", bid: 220}},
		{testInput[4], Hand{cards: []string{"Q", "Q", "Q", "J", "A"}, _type: "four", bid: 483}},
		{"AAAAA 123", Hand{cards: []string{"A", "A", "A", "A", "A"}, _type: "five", bid: 123}},
		{"AAAAK 123", Hand{cards: []string{"A", "A", "A", "A", "K"}, _type: "four", bid: 123}},
		{"AKQJT 123", Hand{cards: []string{"A", "K", "Q", "J", "T"}, _type: "one", bid: 123}},
		{"33322 123", Hand{cards: []string{"3", "3", "3", "2", "2"}, _type: "full", bid: 123}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Returns correct for %s", tc.input), func(t *testing.T) {
			result := Jokerify(tc.input)

			if len(result.cards) != len(tc.expected.cards) {
				t.Fatalf("%d is not %d", len(result.cards), len(tc.expected.cards))
			}

			if result._type != tc.expected._type {
				t.Fatalf("%s is not %s", result._type, tc.expected._type)
			}

			if result.bid != tc.expected.bid {
				t.Fatalf("%d is not %d", result.bid, tc.expected.bid)
			}
		})
	}
}

func TestJokerSort(t *testing.T) {
	useJokers = true

	result := SortHands(ParseHand("KKK2J 123"), ParseHand("KKKJ2 123"))
	expected := 1

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := Part1(testInput)
	expected := 6440

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)
	expected := 5905

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart1Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day7.txt")
	content, _ := os.ReadFile(absPath)

	result := Part1(strings.Split(string(content), "\n"))
	expected := 249638405

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}

func TestPart2Real(t *testing.T) {
	absPath, _ := filepath.Abs("../../inputs/day7.txt")
	content, _ := os.ReadFile(absPath)

	result := Part2(strings.Split(string(content), "\n"))
	expected := 249776650

	if result != expected {
		t.Fatalf("%d is not %d", result, expected)
	}
}
