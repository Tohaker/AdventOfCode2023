package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var handTypes = []string{
	"five",
	"four",
	"full",
	"three",
	"two",
	"one",
	"high",
}

var cardTypes = []string{
	"A",
	"K",
	"Q",
	// no J
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

var useJokers = false

type Hand struct {
	cards []string
	_type string
	bid   int
}

func SortHands(hand1 Hand, hand2 Hand) int {

	// Checks different hand types
	if slices.Index(handTypes, hand1._type) != slices.Index(handTypes, hand2._type) {
		return slices.Index(handTypes, hand2._type) - slices.Index(handTypes, hand1._type)
	}

	// Check the same hand type
	for j, c := range hand1.cards {
		c2 := hand2.cards[j]

		if c == c2 {
			continue
		}

		if useJokers {
			if c == "J" {
				return -1
			} else if c2 == "J" {
				return 1
			}
		}

		numC, err := strconv.Atoi(c)
		numC2, err2 := strconv.Atoi(c2)

		// numC is a string and numC2 is a number
		if err != nil && err2 == nil {
			return 1
		}

		// numC is a number and numC2 is a string
		if err == nil && err2 != nil {
			return -1
		}

		// Both are numbers
		if err == nil && err2 == nil {
			return numC - numC2
		}

		if useJokers {
			if c == "A" {
				return 1
			} else if c == "K" && (c2 != "A" && c2 != "K") {
				return 1
			} else if c == "Q" && (c2 != "A" && c2 != "K" && c2 != "Q") {
				return 1
			} else if c == "T" && (c2 != "A" && c2 != "K" && c2 != "Q" && c2 != "T") {
				return 1
			} else {
				return -1
			}
		}

		if c == "A" {
			return 1
		} else if c == "K" && (c2 != "A" && c2 != "K") {
			return 1
		} else if c == "Q" && (c2 != "A" && c2 != "K" && c2 != "Q") {
			return 1
		} else if c == "J" && (c2 != "A" && c2 != "K" && c2 != "Q" && c2 != "J") {
			return 1
		} else if c == "T" && (c2 != "A" && c2 != "K" && c2 != "Q" && c2 != "J" && c2 != "T") {
			return 1
		} else {
			return -1
		}

	}

	return 0
}

func ParseHand(hand string) Hand {
	split := strings.Split(hand, " ")

	bid, _ := strconv.Atoi(split[1])
	cards := strings.Split(split[0], "")
	_type := ""

	// Figure out what type the hand is
	cardMap := make(map[string]int, 0)
	for _, c := range cards {
		cardMap[c]++
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range cardMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		v := kv.Value

		if v == 5 {
			_type = "five"
		} else if v == 4 {
			_type = "four"
		} else if v == 3 {
			_type = "full"
		} else if v == 2 {
			if _type == "" {
				_type = "one"
				continue
			}

			// Two of a kind
			if _type == "one" {
				_type = "two"
				continue
			}
		} else {
			if _type == "full" {
				_type = "three"
				continue
			}

			if _type == "" {
				_type = "high"
				continue
			}
		}
	}

	return Hand{
		cards,
		_type,
		bid,
	}
}

func Jokerify(hand string) Hand {
	original := ParseHand(hand)
	permutations := make([]Hand, 0)

	for _, c := range cardTypes {
		permutations = append(permutations, ParseHand(strings.ReplaceAll(hand, "J", c)))
	}

	slices.SortFunc(permutations, SortHands)

	return Hand{
		cards: original.cards,
		_type: permutations[len(permutations)-1]._type,
		bid:   original.bid,
	}
}

func Part1(input []string) int {
	useJokers = false

	hands := make([]Hand, len(input))

	for i, h := range input {
		hands[i] = ParseHand(h)
	}

	slices.SortFunc(hands, SortHands)

	result := 0

	for i, h := range hands {
		result += h.bid * (i + 1)
	}

	return result
}

func Part2(input []string) int {
	hands := make([]Hand, len(input))

	for i, h := range input {
		hands[i] = Jokerify(h)
	}

	useJokers = true
	slices.SortFunc(hands, SortHands)

	result := 0

	for i, h := range hands {
		result += h.bid * (i + 1)
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day7.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
