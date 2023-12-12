package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func IsEmptySpace(line []string) bool {
	for _, v := range line {
		if v != "." {
			return false
		}
	}

	return true
}

type ExpandedSpace struct {
	mapLines [][]string
	x        map[int]int
	y        map[int]int
}

// Mark which index has an empty line,
// When it's time to list galaxy points,
// check if those coords are after x number of empty lines
// and increment them accordingly
func GetExpandedSpace(input []string, multiplier int) ExpandedSpace {
	mapLines := make([][]string, 0)
	expandedSpaceX := make(map[int]int, 0)
	expandedSpaceY := make(map[int]int, 0)

	for i, line := range input {
		splitLine := strings.Split(line, "")
		mapLines = append(mapLines, splitLine)

		if IsEmptySpace(splitLine) {
			expandedSpaceY[i] = multiplier
		}
	}

	for i := range mapLines[0] {
		column := make([]string, len(mapLines))

		for j := 0; j < len(mapLines); j++ {
			column[j] = mapLines[j][i]
		}

		if IsEmptySpace(column) {
			expandedSpaceX[i] = multiplier
		}
	}

	return ExpandedSpace{
		mapLines: mapLines,
		x:        expandedSpaceX,
		y:        expandedSpaceY,
	}
}

func ListExpandedGalaxies(expandedSpace ExpandedSpace) []Coordinate {
	coords := make([]Coordinate, 0)

	for y, line := range expandedSpace.mapLines {
		for x, loc := range line {
			if loc == "#" {
				newX := x
				newY := y

				for index, multiplier := range expandedSpace.x {
					if index < x {
						newX += multiplier
					}
				}

				for index, multiplier := range expandedSpace.y {
					if index < y {
						newY += multiplier
					}
				}

				coords = append(coords, Coordinate{newX, newY})
			}
		}
	}

	return coords
}

func MakePairName(c1 Coordinate, c2 Coordinate) string {
	coords := []Coordinate{c1, c2}

	slices.SortFunc(coords, func(a, b Coordinate) int {
		comparison := (a.x + a.y) - (b.x + b.y)

		if comparison == 0 {
			return a.x - b.x
		} else {
			return comparison
		}
	})

	return fmt.Sprintf("%d,%d - %d,%d", coords[0].x, coords[0].y, coords[1].x, coords[1].y)
}

func FindUniquePairsAndDistances(galaxies []Coordinate) map[string]int {
	pairs := make(map[string]int, 0)

	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			// Don't consider the same one
			if galaxy1.x == galaxy2.x && galaxy1.y == galaxy2.y {
				continue
			}

			key := MakePairName(galaxy1, galaxy2)

			_, found := pairs[key]

			if !found {
				pairs[key] = int(math.Abs(float64(galaxy1.x-galaxy2.x)) + math.Abs(float64(galaxy1.y-galaxy2.y)))
			}
		}
	}

	return pairs
}

func Part1(input []string) int {
	pairs := FindUniquePairsAndDistances(ListExpandedGalaxies(GetExpandedSpace(input, 1)))
	result := 0

	for _, v := range pairs {
		result += v
	}

	return result
}

func Part2(input []string, expansionMultiplier int) int {
	pairs := FindUniquePairsAndDistances(ListExpandedGalaxies(GetExpandedSpace(input, expansionMultiplier-1)))
	result := 0

	for _, v := range pairs {
		result += v
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day11.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines, 1000000))
}
