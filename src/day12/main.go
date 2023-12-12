package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type SpringsRecord struct {
	springs []int
	damaged string
}

func ParseInput(input string) SpringsRecord {

	parts := strings.Split(input, " ")
	springsStr := strings.Split(parts[1], ",")

	springs := make([]int, len(springsStr))
	for i, s := range springsStr {
		val, _ := strconv.Atoi(s)

		springs[i] = val
	}

	return SpringsRecord{
		springs: springs,
		damaged: parts[0],
	}
}

func UnfoldSpring(record SpringsRecord) SpringsRecord {
	newDamaged := strings.Repeat(record.damaged, 5)
	newSprings := make([]int, 0)

	for i := 0; i < 5; i++ {
		newSprings = append(newSprings, record.springs...)
	}

	return SpringsRecord{
		newSprings,
		newDamaged,
	}
}

func CreateRegexp(springs []int) regexp.Regexp {
	regex := `^\.*`

	for i, v := range springs {
		regex += fmt.Sprintf(`#{%d}`, v)

		if i != len(springs)-1 {
			regex += `\.+`
		}
	}

	regex += `\.*$`

	return *regexp.MustCompile(regex)
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func CreatePermutations(s string, i int) []string {
	permutations := make([]string, 0)

	// If at the end of the string, return the string
	if i == len(s) {
		permutations = append(permutations, s)
	} else {
		// If s[i] == ?, recurse with . and # on different branches
		if s[i] == '?' {
			permutations = append(permutations, CreatePermutations(ReplaceAtIndex(s, '.', i), i+1)...)
			permutations = append(permutations, CreatePermutations(ReplaceAtIndex(s, '#', i), i+1)...)
		} else {
			permutations = append(permutations, CreatePermutations(s, i+1)...)
		}
	}

	// compile back into single list
	return permutations
}

func FindValidPermutations(record SpringsRecord) []string {
	re := CreateRegexp(record.springs)

	validPermuations := make([]string, 0)

	for _, perm := range CreatePermutations(record.damaged, 0) {
		if re.MatchString(perm) {
			validPermuations = append(validPermuations, perm)
		}
	}

	return validPermuations
}

func Part1(input []string) int {
	result := 0

	for _, line := range input {
		result += len(FindValidPermutations(ParseInput(line)))
	}

	return result
}

func Part2(input []string) int {
	result := 0

	for _, line := range input {
		result += len(FindValidPermutations(UnfoldSpring(ParseInput(line))))
	}

	return result
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day12.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}
