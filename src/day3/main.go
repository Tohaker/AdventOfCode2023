package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ExtractPartNumbers(input []string) []int {
	var partNumbers []int

	var tempStr = ""
	var isValidPart = false

	matchNumbers := regexp.MustCompile(`\d`)
	matchSymbol := regexp.MustCompile(`[^.|\d]`)

	for i, row := range input {
		for j, char := range row {
			character := string(char)

			if matchNumbers.MatchString(character) {
				tempStr += character

				// Find any symbol adjacent to make it a valid part

				// Top
				if i != 0 {
					// N
					if matchSymbol.MatchString(string(input[i-1][j])) {
						isValidPart = true
					}

					// NW
					if j != 0 {
						if matchSymbol.MatchString(string(input[i-1][j-1])) {
							isValidPart = true
						}
					}

					// NE
					if j != len(row)-1 {
						if matchSymbol.MatchString(string(input[i-1][j+1])) {
							isValidPart = true
						}
					}
				}

				// Bottom
				if i != len(input)-1 {
					// S
					if matchSymbol.MatchString(string(input[i+1][j])) {
						isValidPart = true
					}

					// SW
					if j != 0 {
						if matchSymbol.MatchString(string(input[i+1][j-1])) {
							isValidPart = true
						}
					}

					// SE
					if j != len(row)-1 {
						if matchSymbol.MatchString(string(input[i+1][j+1])) {
							isValidPart = true
						}
					}
				}

				// Left
				if j != 0 {
					if matchSymbol.MatchString(string(input[i][j-1])) {
						isValidPart = true
					}
				}

				// Right
				if j != len(row)-1 {
					if matchSymbol.MatchString(string(input[i][j+1])) {
						isValidPart = true
					}
				}
			} else {
				i, err := strconv.Atoi(tempStr)

				tempStr = ""
				if isValidPart && err == nil {
					isValidPart = false
					partNumbers = append(partNumbers, i)
				}
			}
		}
	}

	return partNumbers
}

func ExtractGearRatios(input []string) map[string][]int {
	var gearMap = map[string][]int{}

	tempStr := ""
	gearIndex := ""
	isValidPart := false

	matchNumbers := regexp.MustCompile(`\d`)
	matchGear := regexp.MustCompile(`\*`)

	for i, row := range input {
		for j, char := range row {
			character := string(char)

			if matchNumbers.MatchString(character) {
				tempStr += character

				// Find any symbol adjacent to make it a valid part

				// Top
				if i != 0 {
					// N
					if matchGear.MatchString(string(input[i-1][j])) {
						isValidPart = true
						gearIndex = strconv.Itoa(i-1) + strconv.Itoa(j)
					}

					// NW
					if j != 0 {
						if matchGear.MatchString(string(input[i-1][j-1])) {
							isValidPart = true
							gearIndex = strconv.Itoa(i-1) + strconv.Itoa(j-1)
						}
					}

					// NE
					if j != len(row)-1 {
						if matchGear.MatchString(string(input[i-1][j+1])) {
							isValidPart = true
							gearIndex = strconv.Itoa(i-1) + strconv.Itoa(j+1)
						}
					}
				}

				// Bottom
				if i != len(input)-1 {
					// S
					if matchGear.MatchString(string(input[i+1][j])) {
						isValidPart = true
						gearIndex = strconv.Itoa(i+1) + strconv.Itoa(j)
					}

					// SW
					if j != 0 {
						if matchGear.MatchString(string(input[i+1][j-1])) {
							isValidPart = true
							gearIndex = strconv.Itoa(i+1) + strconv.Itoa(j-1)
						}
					}

					// SE
					if j != len(row)-1 {
						if matchGear.MatchString(string(input[i+1][j+1])) {
							isValidPart = true
							gearIndex = strconv.Itoa(i+1) + strconv.Itoa(j+1)
						}
					}
				}

				// Left
				if j != 0 {
					if matchGear.MatchString(string(input[i][j-1])) {
						isValidPart = true
						gearIndex = strconv.Itoa(i) + strconv.Itoa(j-1)
					}
				}

				// Right
				if j != len(row)-1 {
					if matchGear.MatchString(string(input[i][j+1])) {
						isValidPart = true
						gearIndex = strconv.Itoa(i) + strconv.Itoa(j+1)
					}
				}
			} else {
				i, err := strconv.Atoi(tempStr)

				tempStr = ""
				if isValidPart && err == nil {
					isValidPart = false
					list, ok := gearMap[gearIndex]

					if !ok {
						gearMap[gearIndex] = []int{i}
					} else {
						gearMap[gearIndex] = append(list, i)
					}
				}
			}
		}
	}

	return gearMap
}

func Part1(input []string) int {
	result := 0

	for _, part := range ExtractPartNumbers(input) {
		result += part
	}

	return result
}

func Part2(input []string) int {
	result := 0

	for _, part := range ExtractGearRatios(input) {
		if len(part) == 2 {
			result += part[0] * part[1]
		}
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day3.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
