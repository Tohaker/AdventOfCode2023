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

func IsSmudged(m1 string, m2 string) struct {
	smudged bool
	str     string
} {
	diffCount := 0
	str := ""

	for i := 0; i < len(m1); i++ {
		if m1[i] != m2[i] {
			diffCount++
			if m2[i] == '.' {
				str += "#"
			} else {
				str += "."
			}
		} else {
			str += string(m1[i])
		}
	}

	return struct {
		smudged bool
		str     string
	}{smudged: diffCount == 1, str: str}
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

func FindSmudgedReflection(mirror []string) int {
	reflectionLine := 0
	smudgeCount := 0

	// Find any horizontal reflection first
	rowsAbove := make([]string, 0)

	for i := 0; i < len(mirror)-1; i++ {
		if reflectionLine > 0 {
			break
		}

		rowsAbove = append(rowsAbove, mirror[i])

		isSmudged := IsSmudged(mirror[i], mirror[i+1])

		if smudgeCount == 0 && isSmudged.smudged {
			mirror[i] = isSmudged.str
			mirror[i+1] = isSmudged.str
			smudgeCount++
		}

		if mirror[i+1] == mirror[i] {
			// Check as far as the smallest value
			smallestValue := len(rowsAbove)

			if len(mirror[i:len(mirror)-1]) < len(rowsAbove) {
				smallestValue = len(mirror[i : len(mirror)-1])
			}

			// Start checking reflection
			for j := 0; j < smallestValue; j++ {
				lineSmudged := IsSmudged(mirror[j+i+1], rowsAbove[len(rowsAbove)-j-1])

				if smudgeCount == 0 && lineSmudged.smudged {
					mirror[j+i+1] = lineSmudged.str
					mirror[slices.Index(mirror, rowsAbove[len(rowsAbove)-j-1])] = lineSmudged.str
					rowsAbove[len(rowsAbove)-j-1] = lineSmudged.str
				}

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

func Part2(input []string) int {
	result := 0

	currentMirror := make([]string, 0)

	for i, v := range input {
		if v != "" {
			currentMirror = append(currentMirror, v)
		}

		if v == "" || i == len(input)-1 {
			h := 100 * FindSmudgedReflection(currentMirror)
			result += h
			if h == 0 {
				result += FindSmudgedReflection(RotateMirror(currentMirror))
			}
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
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
