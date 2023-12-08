package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type NodeMap struct {
	instructions []string
	nodes        map[string][]string
}

func ParseInput(input string) NodeMap {
	blocks := strings.Split(input, "\n\n")

	instructions := strings.Split(blocks[0], "")
	nodes := make(map[string][]string, 0)

	for _, v := range strings.Split(blocks[1], "\n") {
		b2 := strings.Split(v, " = ")

		key := b2[0]
		values := strings.Split(strings.TrimSuffix(strings.TrimPrefix(b2[1], "("), ")"), ", ")

		nodes[key] = values
	}

	return NodeMap{
		instructions,
		nodes,
	}
}

func Part1(input string) int {
	nodeMap := ParseInput(input)

	node := nodeMap.nodes["AAA"]
	stepCount := 0

	for i := 0; i < len(nodeMap.instructions); i++ {
		instruction := nodeMap.instructions[i]
		next := ""

		if instruction == "L" {
			next = node[0]
		} else {
			next = node[1]
		}

		stepCount++

		if next == "ZZZ" {
			break
		}

		node = nodeMap.nodes[next]

		if i == len(nodeMap.instructions)-1 {
			i = -1 // One before so we count from 0
		}
	}

	return stepCount
}

// Recursive function to return gcd of a and b
func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func findlcm(arr []int) int {
	// Initialize result
	ans := arr[0]

	// ans contains LCM of arr[0], ..arr[i]
	// after i'th iteration,
	for i := 1; i < len(arr); i++ {
		ans = ((arr[i] * ans) /
			(gcd(arr[i], ans)))
	}

	return ans
}

func Part2(input string) int {
	nodeMap := ParseInput(input)

	keys := make([]string, 0)

	for k := range nodeMap.nodes {
		if strings.HasSuffix(k, "A") {
			keys = append(keys, k)
		}
	}

	stepCount := make([]int, 0, len(keys))

	for j := 0; j < len(keys); j++ {
		node := nodeMap.nodes[keys[j]]
		currentCount := 0

		for i := 0; i < len(nodeMap.instructions); i++ {
			instruction := nodeMap.instructions[i]
			next := ""

			if instruction == "L" {
				next = node[0]
			} else {
				next = node[1]
			}

			currentCount++

			if strings.HasSuffix(next, "Z") {
				break
			}

			node = nodeMap.nodes[next]

			if i == len(nodeMap.instructions)-1 {
				i = -1 // One before so we count from 0
			}
		}

		stepCount = append(stepCount, currentCount)
	}

	return findlcm(stepCount)
}

func main() {
	absPath, _ := filepath.Abs("../../inputs/day8.txt")
	content, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", Part1(string(content)))
	fmt.Printf("Part 2: %d\n", Part2(string(content)))
}
