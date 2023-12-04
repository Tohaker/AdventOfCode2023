package main

import (
	"fmt"
	"math/bits"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CalculateMatches(card string) int {
	cols := strings.Split(strings.Split(card, ":")[1], "|")

	re := regexp.MustCompile(`\d+`)

	winning := re.FindAllString(cols[0], -1)
	yours := re.FindAllString(cols[1], -1)

	winningMap := make(map[string]struct{}, 0)
	common := make([]string, 0)

	for _, w := range winning {
		winningMap[w] = struct{}{}
	}

	for _, y := range yours {
		if _, found := winningMap[y]; found {
			common = append(common, y)
		}
	}

	return len(common)

}

func Part1(input []string) int {
	result := 0

	for _, card := range input {
		length := CalculateMatches(card)

		if length != 0 {
			result += int(bits.RotateLeft32(1, length-1))
		}
	}

	return result
}

// TODO: Make this waaay more efficient
func Part2(input []string) int {
	cards := map[int][]string{}

	// fill card list
	for i, card := range input {
		cards[i+1] = []string{card}
	}

	for i := 1; i <= len(cards); i++ {
		for _, card := range cards[i] {
			winning := CalculateMatches(card)

			for j := 1; j <= winning; j++ {
				cards[i+j] = append(cards[i+j], cards[i+j][0])
			}
		}
	}

	result := 0

	for _, list := range cards {
		result += len(list)
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day4.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
