package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ParseSequence(input string) []int {
	ints := strings.Split(input, " ")

	seq := make([]int, len(ints))

	for i, s := range ints {
		seq[i], _ = strconv.Atoi(s)
	}

	return seq
}

func AllZero(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}

	return true
}

func DifferenceEngine(sequence []int, last int) int {
	diffs := make([]int, len(sequence))

	for i := 0; i < len(sequence)-1; i++ {
		diffs[i] = sequence[i+1] - sequence[i]
	}

	if !AllZero(diffs) {
		diffs[len(diffs)-1] = DifferenceEngine(diffs[:len(diffs)-1], diffs[len(diffs)-2])
	}

	return last + diffs[len(diffs)-1]
}

func Part1(input []string) int {
	total := 0

	for _, v := range input {
		inputSeq := ParseSequence(v)
		total += inputSeq[len(inputSeq)-1] + DifferenceEngine(inputSeq, 0)
	}

	return total
}

func Part2(input []string) int {
	total := 0

	for _, v := range input {
		inputSeq := ParseSequence(v)
		for i, j := 0, len(inputSeq)-1; i < j; i, j = i+1, j-1 {
			inputSeq[i], inputSeq[j] = inputSeq[j], inputSeq[i]
		}

		total += inputSeq[len(inputSeq)-1] + DifferenceEngine(inputSeq, 0)
	}

	return total
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day9.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
