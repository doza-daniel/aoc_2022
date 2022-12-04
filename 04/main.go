package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res, err := solve("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.NewReader(scanner.Text())
		var a, b, x, y int
		fmt.Fscanf(line, "%d-%d,%d-%d", &a, &b, &x, &y)
		if overlap(a, b, x, y) {
			result += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}

func overlap(a, b, x, y int) bool {
	return (a <= x && b >= x) || (x <= a && y >= a)
}

func oneContainsOther(a, b, x, y int) bool {
	return (a <= x && b >= y) || (x <= a && y >= b)
}
