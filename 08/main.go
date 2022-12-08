package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartOne(grid))
	fmt.Println(solvePartTwo(grid))
}

func parse(input string) ([][]int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		grid = append(grid, make([]int, len(line)))
		for j, r := range line {
			grid[i][j] = int(r - '0')
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func solvePartOne(grid [][]int) int {
	var n int = len(grid)

	if n == 0 || n == 1 {
		return n
	}

	visible := 4*n - 4

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {

			var hiddenFromLeft, hiddenFromRight, hiddenFromTop, hiddenFromBot bool

			for k := 0; k < n; k++ {
				if k < i && grid[k][j] >= grid[i][j] {
					hiddenFromTop = true
				}

				if k > i && grid[k][j] >= grid[i][j] {
					hiddenFromBot = true
				}

				if k < j && grid[i][k] >= grid[i][j] {
					hiddenFromLeft = true
				}

				if k > j && grid[i][k] >= grid[i][j] {
					hiddenFromRight = true
				}
			}

			if !(hiddenFromBot && hiddenFromTop && hiddenFromLeft && hiddenFromRight) {
				visible += 1
			}
		}
	}

	return visible
}

func solvePartTwo(grid [][]int) int {
	var n int = len(grid)

	if n == 0 || n == 1 {
		return 0
	}

	maxScenicScore := 0

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {

			var top int
			for k := i - 1; k >= 0; k-- {
				top++
				if grid[k][j] >= grid[i][j] {
					break
				}
			}

			var bot int
			for k := i + 1; k < n; k++ {
				bot++
				if grid[k][j] >= grid[i][j] {
					break
				}
			}

			var left int
			for k := j - 1; k >= 0; k-- {
				left++
				if grid[i][k] >= grid[i][j] {
					break
				}
			}

			var right int
			for k := j + 1; k < n; k++ {
				right++
				if grid[i][k] >= grid[i][j] {
					break
				}
			}

			scenicScore := top * bot * left * right
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}
