package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Contents struct {
	red   int
	blue  int
	green int
}

func CountMaxInGame(game string) Contents {
	var contents = Contents{}

	roundsStr := strings.Split(game, `: `)
	rounds := strings.Split(roundsStr[1], `; `)

	for _, v := range rounds {
		draws := strings.Split(v, `, `)

		for _, draw := range draws {
			splitDraw := strings.Split(draw, " ")

			count, _ := strconv.Atoi(splitDraw[0])

			switch splitDraw[1] {
			case `red`:
				if contents.red < count {
					contents.red = count
				}
			case `blue`:
				if contents.blue < count {
					contents.blue = count
				}
			case `green`:
				if contents.green < count {
					contents.green = count
				}
			}
		}
	}

	return contents
}

func SumPossibleGames(games []string) int {
	total := 0

	for i, game := range games {
		contents := CountMaxInGame(game)

		if (contents.red <= 12) && (contents.blue <= 14) && (contents.green <= 13) {
			total += (i + 1)
		}
	}

	return total
}

func CalcPowerInGames(games []string) int {
	total := 0

	for _, game := range games {
		contents := CountMaxInGame(game)

		total += (contents.red * contents.blue * contents.green)

	}

	return total
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day2.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", SumPossibleGames(lines))
	fmt.Printf("Part 2: %d\n", CalcPowerInGames(lines))
}
