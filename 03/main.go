package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	part1, err := solve("input")
	if err != nil {
		panic(err)
	}

	part2, err := solvePartTwo("input")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func solvePartTwo(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	groupSize := 3
	group := make([]string, groupSize)

	for lineNo := 0; scanner.Scan(); lineNo++ {
		line := scanner.Text()

		if lineNo > 0 && lineNo%groupSize == 0 {
			r := findCommonRune(group)
			result += calcScore(r)
		}
		group[lineNo%groupSize] = line
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		r := inBothHalves(line)
		result += calcScore(r)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}

func inBothHalves(line string) rune {
	half := len(line) / 2

	found := make(map[rune]bool)
	for i, c := range line {
		if i < half {
			found[c] = true
		} else if found[c] {
			return c
		}
	}
	panic("unreachable")
}

func findCommonRune(lines []string) rune {
	count := make(map[rune]int)
	for _, line := range lines {
		found := make(map[rune]bool)
		for _, r := range line {
			if !found[r] {
				found[r] = true
				count[r]++
			}

			if count[r] == len(lines) {
				return r
			}
		}
	}

	panic("unreachanble")
}

func calcScore(r rune) int {
	if unicode.IsUpper(r) {
		return int(r-'A') + 27
	}

	return int(r-'a') + 1
}
