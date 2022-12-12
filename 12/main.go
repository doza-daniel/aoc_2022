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
	fmt.Println(solvePartOne(copyGrid(grid)))
	fmt.Println(solvePartTwo(copyGrid(grid)))
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

	ds := dijkstra(grid, start, newNeighborsFn(uphill))
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
	ds := dijkstra(grid, end, newNeighborsFn(downhill))
	for i, row := range ds {
		for j, d := range row {
			if grid[i][j] == 'a' && d < min {
				min = d
			}
		}
	}

	return min
}

func dijkstra(grid [][]byte, init Pos, neighborsFn func([][]byte, Pos) <-chan Pos) [][]int {
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

		neighbors := neighborsFn(grid, currentNode)
		for n := range neighbors {
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

func newNeighborsFn(direction func(byte, byte) bool) func([][]byte, Pos) <-chan Pos {
	return func(grid [][]byte, curr Pos) <-chan Pos {
		return neighbors(grid, curr, direction)
	}
}

func uphill(from, to byte) bool {
	return int(to)-int(from) <= 1
}

func downhill(from, to byte) bool {
	return int(to)-int(from) >= -1
}

func neighbors(grid [][]byte, c Pos, canReach func(byte, byte) bool) <-chan Pos {
	ch := make(chan Pos)

	inBounds := func(p Pos) bool {
		return p.i >= 0 && p.i <= len(grid)-1 && p.j >= 0 && p.j <= len(grid[p.i])-1
	}

	go func() {
		defer close(ch)

		directions := []Pos{
			{i: c.i - 1, j: c.j},
			{i: c.i + 1, j: c.j},
			{i: c.i, j: c.j + 1},
			{i: c.i, j: c.j - 1},
		}

		from := grid[c.i][c.j]
		for _, d := range directions {
			if inBounds(d) {
				to := grid[d.i][d.j]
				if canReach(from, to) {
					ch <- d
				}
			}
		}
	}()

	return ch
}

func copyGrid(grid [][]byte) [][]byte {
	c := make([][]byte, len(grid))
	for i, row := range grid {
		c[i] = make([]byte, len(row))
		for j, col := range grid[i] {
			c[i][j] = col
		}
	}
	return c
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
