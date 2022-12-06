package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	index, err := solve("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(index)
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return findFirstRun(line, 14), nil
}

func findFirstRun(line string, runLen int) int {
	for i := range line {
		var j int = i
		found := make(map[rune]bool)
		for j < len(line) && j < i+runLen {
			var r rune = rune(line[j])
			if found[r] {
				break
			}
			found[r] = true
			j++
		}
		if j == i+runLen {
			return i + runLen
		}
	}

	panic("unreachable")
}
