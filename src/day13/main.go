package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func RotateMirror(mirror []string) []string {
	rotatedMirror := make([]string, 0)
	newMirror := make([]string, len(mirror))
	copy(newMirror, mirror)

	slices.Reverse(newMirror)

	for j := 0; j < len(newMirror[0]); j++ {
		line := ""

		for _, v := range newMirror {
			line += string(v[j])
		}

		rotatedMirror = append(rotatedMirror, line)
	}

	return rotatedMirror
}

func FindReflection(mirror []string) int {
	reflectionLine := 0

	// Find any horizontal reflection first
	rowsAbove := make([]string, 0)

	for i := 0; i < len(mirror)-1; i++ {
		if reflectionLine > 0 {
			break
		}

		rowsAbove = append(rowsAbove, mirror[i])

		if mirror[i+1] == mirror[i] {
			// Check as far as the smallest value
			smallestValue := len(rowsAbove)

			if len(mirror[i:len(mirror)-1]) < len(rowsAbove) {
				smallestValue = len(mirror[i : len(mirror)-1])
			}

			// Start checking reflection
			for j := 0; j < smallestValue; j++ {
				if mirror[j+i+1] != rowsAbove[len(rowsAbove)-j-1] {
					reflectionLine = 0
					break
				}
				reflectionLine = len(rowsAbove)
			}
		}
	}

	return reflectionLine
}

func Part1(input []string) int {
	result := 0

	currentMirror := make([]string, 0)

	for i, v := range input {
		if v != "" {
			currentMirror = append(currentMirror, v)
		}

		if v == "" || i == len(input)-1 {
			result += 100 * FindReflection(currentMirror)
			result += FindReflection(RotateMirror(currentMirror))
			currentMirror = make([]string, 0)
		}
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day13.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	// fmt.Printf("Part 2: %d\n", Part2(lines))
}
