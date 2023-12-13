package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindReflection(mirror []string) int {
	reflectionLine := 0

	// Find any horizontal reflection first
	rowsAbove := make([]string, 0)

	for i := 0; i < len(mirror)-1; i++ {
		rowsAbove = append(rowsAbove, mirror[i])

		if mirror[i+1] == mirror[i] {
			// Start checking reflection
			for j := 0; j < len(rowsAbove); j++ {
				if j+i+1 == len(mirror) {
					break
				}

				if mirror[j+i+1] != rowsAbove[len(rowsAbove)-j-1] {
					reflectionLine = 0
					break
				}
				reflectionLine = 100 * len(rowsAbove)
			}
		}
	}

	if reflectionLine > 0 {
		return reflectionLine
	}

	// Find any vertical reflection next
	colsLeft := make([]string, 0)

	rotatedMirror := make([]string, 0)

	for j := 0; j < len(mirror[0]); j++ {
		line := ""

		for _, v := range mirror {
			line += string(v[j])
		}

		rotatedMirror = append(rotatedMirror, line)
	}

	for i := 0; i < len(rotatedMirror)-1; i++ {
		colsLeft = append(colsLeft, rotatedMirror[i])

		if rotatedMirror[i+1] == rotatedMirror[i] {
			// Start checking reflection
			for j := 0; j < len(colsLeft); j++ {
				if j+i+1 == len(rotatedMirror) {
					break
				}

				if rotatedMirror[j+i+1] != colsLeft[len(colsLeft)-j-1] {
					reflectionLine = 0
					break
				}
				reflectionLine = len(colsLeft)
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
			result += FindReflection(currentMirror)
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
