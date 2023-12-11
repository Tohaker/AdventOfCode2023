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

func ExpandMap(input []string, multiplier int) [][]string {
	expandedMapLines := make([][]string, 0)

	for _, line := range input {
		splitLine := strings.Split(line, "")
		expandedMapLines = append(expandedMapLines, splitLine)

		if IsEmptySpace(splitLine) {
			for k := 0; k < multiplier; k++ {
				expandedMapLines = append(expandedMapLines, splitLine)
			}
		}
	}

	expandedMap := make([][]string, len(expandedMapLines))
	for i := range expandedMapLines {
		expandedMap[i] = make([]string, 0)
	}

	for i := range expandedMapLines[0] {
		column := make([]string, len(expandedMapLines))

		for j := 0; j < len(expandedMapLines); j++ {
			column[j] = expandedMapLines[j][i]
			expandedMap[j] = append(expandedMap[j], column[j])
		}

		if IsEmptySpace(column) {
			for j := 0; j < len(expandedMapLines); j++ {
				for k := 0; k < multiplier; k++ {
					expandedMap[j] = append(expandedMap[j], ".")
				}
			}
		}
	}

	return expandedMap
}

func ListGalaxies(expandedMap [][]string) []Coordinate {
	coords := make([]Coordinate, 0)

	for y, line := range expandedMap {
		for x, loc := range line {
			if loc == "#" {
				coords = append(coords, Coordinate{x, y})
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
	pairs := FindUniquePairsAndDistances(ListGalaxies(ExpandMap(input, 1)))
	result := 0

	for _, v := range pairs {
		result += v
	}

	return result
}

func Part2(input []string, expansionMultiplier int) int {
	pairs := FindUniquePairsAndDistances(ListGalaxies(ExpandMap(input, expansionMultiplier-1)))
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
