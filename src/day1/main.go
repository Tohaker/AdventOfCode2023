package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ExtractNumber(input string) int {
	re := regexp.MustCompile(`\d`)
	numbers := re.FindAllString(input, -1)

	strVal := numbers[0] + numbers[len(numbers)-1]
	intVal, err := strconv.Atoi(strVal)

	if err != nil {
		panic(err)
	}

	return intVal
}

func ReplaceStringNumber(input string) string {
	// TODO: Make this nicer
	newString := input
	newString = strings.ReplaceAll(newString, "one", "o1e")
	newString = strings.ReplaceAll(newString, "two", "t2o")
	newString = strings.ReplaceAll(newString, "three", "t3e")
	newString = strings.ReplaceAll(newString, "four", "f4r")
	newString = strings.ReplaceAll(newString, "five", "f5e")
	newString = strings.ReplaceAll(newString, "six", "s6x")
	newString = strings.ReplaceAll(newString, "seven", "s7n")
	newString = strings.ReplaceAll(newString, "eight", "e8t")
	newString = strings.ReplaceAll(newString, "nine", "n9e")

	return newString
}

func Part1(input []string) int {
	result := 0

	for _, s := range input {
		result += ExtractNumber(s)
	}

	return result
}

func Part2(input []string) int {
	result := 0

	for _, s := range input {
		result += ExtractNumber(ReplaceStringNumber(s))
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day1.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
