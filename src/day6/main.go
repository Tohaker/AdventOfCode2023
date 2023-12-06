package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile(`\d+`)

func BeatRecord(time int, distance int) int {
	result := 0

	for speed := 1; speed < time; speed++ {

		raceDistance := (time - speed) * speed

		if raceDistance > distance {
			result++
		}
	}

	return result
}

func Part1(input []string) int {
	times := numberRegex.FindAllString(input[0], -1)
	distances := numberRegex.FindAllString(input[1], -1)

	result := 1

	for i, t := range times {
		time, _ := strconv.Atoi(t)
		distance, _ := strconv.Atoi(distances[i])

		result *= BeatRecord(time, distance)
	}

	return result
}

func Part2(input []string) int {
	vt := strings.Split(input[0], ":")[1]
	vd := strings.Split(input[1], ":")[1]

	time, _ := strconv.Atoi(strings.ReplaceAll(vt, " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(vd, " ", ""))

	return BeatRecord(time, distance)
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day6.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
