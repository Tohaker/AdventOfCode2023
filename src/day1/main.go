package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ConvertStringsToNumber(a string, b string) int {
	strVal := a + b
	intVal, err := strconv.Atoi(strVal)

	if err != nil {
		panic(err)
	}

	return intVal
}

func ExtractNumber(input string) int {
	re := regexp.MustCompile(`\d`)
	numbers := re.FindAllString(input, -1)

	return ConvertStringsToNumber(numbers[0], numbers[len(numbers)-1])
}

var stringToNumberLookup = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func ExtractStringNumbers(input string) int {
	forwardRE := regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|\d`)
	reverseRE := regexp.MustCompile(`(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)|\d`)

	forwardMatch := (forwardRE.FindAllString(input, 1))[0]
	reverseMatch := (reverseRE.FindAllString(reverse(input), 1))[0]

	first, ok := stringToNumberLookup[forwardMatch]

	if !ok {
		first = forwardMatch
	}

	reversedReverseMatch := reverse(reverseMatch)
	second, ok := stringToNumberLookup[reversedReverseMatch]

	if !ok {
		second = reversedReverseMatch
	}

	return ConvertStringsToNumber(first, second)
}

func Part1(input []string) int {
	result := 0

	for _, s := range input {
		result += ExtractNumber(s)
	}

	return result
}

func Part2(input []string) int {
	result := 0

	for _, s := range input {
		result += ExtractStringNumbers(s)
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day1.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
