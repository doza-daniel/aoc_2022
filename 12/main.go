package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	grid, err := parse("input")
	if err != nil {
		panic(err)
	}
	// fmt.Println(solvePartOne(grid))
	fmt.Println(solvePartTwo(grid))
}

type Pos struct {
	i, j int
}

func solvePartOne(grid [][]byte) int {
	var start, end Pos
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				start = Pos{i: i, j: j}
				grid[i][j] = 'a'
			}

			if grid[i][j] == 'E' {
				end = Pos{i: i, j: j}
				grid[i][j] = 'z'
			}
		}
	}

	ds := dijkstra(grid, start, newNeighborFn(uphill))
	for i, row := range ds {
		for j, d := range row {
			p := Pos{i, j}
			if p == end {
				return d
			}
		}
	}

	panic("asdf")
}

func solvePartTwo(grid [][]byte) int {
	var end Pos
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = 'a'
			}

			if grid[i][j] == 'E' {
				end = Pos{i: i, j: j}
				grid[i][j] = 'z'
			}
		}
	}

	min := math.MaxInt32
	ds := dijkstra(grid, end, newNeighborFn(downhill))
	for i, row := range ds {
		for j, d := range row {
			if grid[i][j] == 'a' && d < min {
				min = d
			}
		}
	}

	return min
}

func dijkstra(grid [][]byte, init Pos, neighborFn func([][]byte, Pos) []Pos) [][]int {
	visited := make(map[Pos]bool)
	unvisited := make(map[Pos]struct{})
	distances := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		distances[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distances[i][j] = math.MaxInt32
			unvisited[Pos{i, j}] = struct{}{}
		}
	}
	distances[init.i][init.j] = 0

	currentNode := init
	for {
		newD := distances[currentNode.i][currentNode.j] + 1

		neighbors := neighborFn(grid, currentNode)
		for _, n := range neighbors {
			if visited[n] {
				continue
			}

			oldD := distances[n.i][n.j]
			if newD < oldD {
				distances[n.i][n.j] = newD
			}
		}

		delete(unvisited, currentNode)
		visited[currentNode] = true

		min := math.MaxInt32
		for p := range unvisited {
			d := distances[p.i][p.j]
			if d < min {
				min = d
				currentNode = p
			}
		}

		if min == math.MaxInt32 {
			return distances
		}
	}
}

func newNeighborFn(direction func(byte, byte) bool) func([][]byte, Pos) []Pos {
	return func(grid [][]byte, curr Pos) []Pos {
		return getNeighbors(grid, curr, direction)
	}
}

func getNeighbors(grid [][]byte, curr Pos, canReach func(from, to byte) bool) []Pos {
	var neighbors []Pos = make([]Pos, 0)

	up := Pos{i: curr.i - 1, j: curr.j}
	down := Pos{i: curr.i + 1, j: curr.j}
	right := Pos{i: curr.i, j: curr.j + 1}
	left := Pos{i: curr.i, j: curr.j - 1}

	from := grid[curr.i][curr.j]

	if curr.i < len(grid)-1 {
		to := grid[down.i][down.j]

		if canReach(from, to) {
			neighbors = append(neighbors, down)
		}
	}

	if curr.i > 0 {
		to := grid[up.i][up.j]

		if canReach(from, to) {
			neighbors = append(neighbors, up)
		}
	}

	if curr.j < len(grid[curr.i])-1 {
		to := grid[right.i][right.j]

		if canReach(from, to) {
			neighbors = append(neighbors, right)
		}
	}

	if curr.j > 0 {
		to := grid[left.i][left.j]

		if canReach(from, to) {
			neighbors = append(neighbors, left)
		}
	}

	return neighbors
}

func uphill(from, to byte) bool {
	return int(to)-int(from) <= 1
}

func downhill(from, to byte) bool {
	return int(to)-int(from) >= -1
}

func parse(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
