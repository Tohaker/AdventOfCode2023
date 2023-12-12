package main

import (
	"fmt"
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

func CreateRegexp(springs []int) regexp.Regexp {
	regex := `^\\.*`

	for i, v := range springs {
		regex += fmt.Sprintf(`#{%d}`, v)

		if i != len(springs)-1 {
			regex += `\\.+`
		}
	}

	regex += `\\.*$`

	return *regexp.MustCompile(regex)
}

func CreateValidPermutations(record SpringsRecord) []string {
	re := CreateRegexp(record.springs)

	// Count number of ?
	unknownCount := strings.Count(record.damaged, "?")
	splitString := strings.Split(record.damaged, "")

	validPermuations := make([]string, 0)

	return validPermuations
}
