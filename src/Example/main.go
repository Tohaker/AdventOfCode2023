package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MoveBetweenFloors(input string) int {
	floor := 0

	movements := strings.Split(input, "")

	for _, s := range movements {
		if s == "(" {
			floor++
		} else if s == ")" {
			floor--
		}
	}

	return floor
}

func Part1(path string) int {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return MoveBetweenFloors(string(content))
}

func FindFirstTimeInBasement(input string) int {
	floor := 0

	movements := strings.Split(input, "")

	for i, s := range movements {
		if s == "(" {
			floor++
		} else if s == ")" {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}

func Part2(path string) int {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return FindFirstTimeInBasement(string(content))
}

func main() {
	absPath, _ := filepath.Abs("./src/Example/input.txt")

	fmt.Printf("Part 1: %d\n", Part1(absPath))
	fmt.Printf("Part 2: %d\n", Part2(absPath))
}
