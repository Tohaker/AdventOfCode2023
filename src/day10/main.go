package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func ParseTiles(tiles []string) [][]string {
	grid := make([][]string, len(tiles))

	for i, line := range tiles {
		grid[i] = strings.Split(line, "")
	}

	return grid
}

func FindNextTile(currentTile string, currentCoord Coordinate, previousCoord Coordinate) Coordinate {
	nextCoord := currentCoord

	switch currentTile {
	case "|":
		if previousCoord.y < currentCoord.y {
			nextCoord.y++
		} else {
			nextCoord.y--
		}

	case "-":
		if previousCoord.x < currentCoord.x {
			nextCoord.x++
		} else {
			nextCoord.x--
		}

	case "L":
		if previousCoord.y < currentCoord.y {
			nextCoord.x++
		} else {
			nextCoord.y--
		}

	case "J":
		if previousCoord.y < currentCoord.y {
			nextCoord.x--
		} else {
			nextCoord.y--
		}

	case "7":
		if previousCoord.y > currentCoord.y {
			nextCoord.x--
		} else {
			nextCoord.y++
		}

	case "F":
		if previousCoord.y > currentCoord.y {
			nextCoord.x++
		} else {
			nextCoord.y++
		}
	}

	return nextCoord
}

func FindLoop(grid [][]string) []Coordinate {
	/// Find the start tile
	startCoord := Coordinate{-1, -1}

	for i, v := range grid {
		for j, tile := range v {
			if tile == "S" {
				startCoord.x = j
				startCoord.y = i
				break
			}
		}
	}

	/// Find any connecting pipe to the start
	startingTile := startCoord

	left := ""
	right := ""
	up := ""
	down := ""

	if startCoord.x > 0 {
		left = grid[startCoord.y][startCoord.x-1]
	}

	if startCoord.x < len(grid[0])-1 {
		right = grid[startCoord.y][startCoord.x+1]
	}

	if startCoord.y > 0 {
		up = grid[startCoord.y-1][startCoord.x]
	}

	if startCoord.y < len(grid)-1 {
		down = grid[startCoord.y+1][startCoord.x]
	}

	if left == "-" || left == "L" || left == "F" { // Left (-, L, F)
		startingTile.x = startingTile.x - 1
	} else if right == "-" || right == "J" || right == "7" { // Right (-, J, 7)
		startingTile.x = startingTile.x + 1
	} else if up == "|" || up == "7" || up == "F" { // Up (|, 7, F)
		startingTile.y = startingTile.y - 1
	} else if down == "|" || down == "L" || down == "J" { // Down (|, L, J)
		startingTile.y = startingTile.y + 1
	}

	if startingTile.x == startCoord.x && startingTile.y == startCoord.y {
		panic("No pipes connected to S")
	}

	/// From that pipe, start finding the next connected pipe

	previousTile := startCoord
	currentTile := startingTile

	tilesInPipe := make([]Coordinate, 0)

	for {
		if currentTile.x == startCoord.x && currentTile.y == startCoord.y {
			break
		}

		tilesInPipe = append(tilesInPipe, currentTile)

		_tmp := currentTile
		currentTile = FindNextTile(grid[currentTile.y][currentTile.x], currentTile, previousTile)
		previousTile = _tmp
	}

	return tilesInPipe
}

func Part1(tiles []string) int {
	grid := ParseTiles(tiles)

	tilesInPipe := FindLoop(grid)

	// Get median value for longest distance from start
	if len(tilesInPipe)%2 == 0 {
		return len(tilesInPipe) / 2
	} else {
		return (len(tilesInPipe) + 1) / 2
	}
}

func Part2(tiles []string) int {
	grid := ParseTiles(tiles)

	tilesInPipe := FindLoop(grid)

	// Find a square a quarter of the size of the input
	// https://www.reddit.com/r/adventofcode/comments/18ez5jb/2023_day_10_part_2_shortcut_solution_using_shape/

	squareX := len(grid[0]) / 4
	squareY := len(grid) / 4

	count := 0

	for i, row := range grid[squareY:(len(grid) - squareY)] {
		for j, tile := range row[squareX:(len(row) - squareX)] {
			x := j + squareX
			y := i + squareY

			if !slices.Contains(tilesInPipe, Coordinate{x, y}) && tile != "S" {
				count++
			}
		}
	}

	return count
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day10.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
