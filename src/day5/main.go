package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type AlmanacMap = []struct {
	source      int
	destination int
	_range      int
}

var numbersRegex = regexp.MustCompile(`\d+`)

func ParseMaps(input string) []AlmanacMap {
	blocks := strings.Split(input, "\n\n")

	result := make([]AlmanacMap, 0)

	for _, block := range blocks {
		almanacMap := make(AlmanacMap, 0)

		lines := strings.Split(block, "\n")

		for _, line := range lines[1:] {
			values := numbersRegex.FindAllString(line, -1)

			if len(values) == 0 {
				break
			}

			source, _ := strconv.Atoi(values[1])
			destination, _ := strconv.Atoi(values[0])
			_range, _ := strconv.Atoi(values[2])

			almanacMap = append(almanacMap, struct {
				source      int
				destination int
				_range      int
			}{
				source,
				destination,
				_range,
			})
		}

		result = append(result, almanacMap)
	}

	return result
}

func ParseAlmanac(input string) struct {
	seeds []int
	maps  []AlmanacMap
} {
	sections := strings.SplitN(input, "\n\n", 2)

	seedsStr := numbersRegex.FindAllString(sections[0], -1)
	seeds := make([]int, len(seedsStr))

	for i, v := range seedsStr {
		value, _ := strconv.Atoi(v)
		seeds[i] = value
	}

	maps := ParseMaps(sections[1])

	return struct {
		seeds []int
		maps  []AlmanacMap
	}{
		seeds,
		maps,
	}
}

func ProcessSeed(seed int, almanac AlmanacMap) int {
	for _, block := range almanac {
		if block.source <= seed && (block.source+block._range-1) >= seed {
			return block.destination + (seed - block.source)
		}
	}

	return seed
}

func Part1(input string) int {
	result := -1

	almanac := ParseAlmanac(input)

	for _, seed := range almanac.seeds {
		value := seed

		for _, alm := range almanac.maps {
			value = ProcessSeed(value, alm)
		}

		if result < 0 || value < result {
			result = value
		}
	}

	return result
}

func Part2(input string) int {
	result := -1

	almanac := ParseAlmanac(input)

	for i := 0; i < len(almanac.seeds); i += 2 {
		expandedSeeds := make([]int, almanac.seeds[i+1])

		for j := 0; j < almanac.seeds[i+1]; j++ {
			expandedSeeds[j] = almanac.seeds[i] + j
		}

		for _, seed := range expandedSeeds {
			value := seed

			for _, alm := range almanac.maps {
				value = ProcessSeed(value, alm)
			}

			if result < 0 || value < result {
				result = value

				fmt.Println(value)
			}
		}

		expandedSeeds = nil
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day5.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", Part1(string(content)))
	fmt.Printf("Part 2: %d\n", Part2(string(content)))
}
