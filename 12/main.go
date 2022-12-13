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
	return ds[end]

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

	min := infinity
	ds := dijkstra(grid, end, newNeighborsFn(downhill))

	for p, d := range ds {
		if grid[p.i][p.j] == 'a' && d < min {
			min = d
		}
	}

	return min
}

const infinity int = math.MaxInt32

func dijkstra(grid [][]byte, init Pos, neighborsFn func([][]byte, Pos) <-chan Pos) map[Pos]int {
	// visited := make(map[Pos]bool)
	distances := make(map[Pos]int)

	unvisited := &MinHeap{
		search: make(map[Pos]int),
		less: func(a, b Pos) bool {
			return distances[a] < distances[b]
		},
	}

	for i := range grid {
		for j := range grid[i] {
			p := Pos{i, j}

			if p == init {
				distances[p] = 0
				unvisited.push(p)
			} else {
				distances[p] = infinity
			}
		}
	}

	for len(unvisited.storage) > 0 {
		c := unvisited.pop()

		for n := range neighborsFn(grid, c) {
			alt := distances[c] + 1
			if alt < distances[n] {
				distances[n] = alt
			}

			_, found := unvisited.search[n]
			if !found {
				unvisited.push(n)
			} else {
				unvisited.recompute(n)
			}

		}
	}

	return distances
}

type MinHeap struct {
	storage []Pos
	search  map[Pos]int
	less    func(i, j Pos) bool
}

func (m *MinHeap) recompute(n Pos) {
	i := m.search[n]
	m.search[n] = i
	for {
		parent := (i - 1) / 2

		if i == 0 || !m.less(m.storage[i], m.storage[parent]) {
			break
		}

		m.storage[i], m.storage[parent] = m.storage[parent], m.storage[i]

		i = parent
		m.search[n] = i
	}
}

func (m *MinHeap) pop() Pos {
	min := m.storage[0]

	size := len(m.storage)
	m.storage[0] = m.storage[size-1]
	m.storage = m.storage[:size-1]

	m.heapify(0)
	delete(m.search, min)

	return min
}

func (m *MinHeap) push(n Pos) {
	m.storage = append(m.storage, n)

	i := len(m.storage) - 1
	m.search[n] = i
	for {
		parent := (i - 1) / 2

		if i == 0 || !m.less(m.storage[i], m.storage[parent]) {
			break
		}

		m.storage[i], m.storage[parent] = m.storage[parent], m.storage[i]

		i = parent
		m.search[n] = i
	}
}

func (m *MinHeap) heapify(i int) {
	for {
		var smallest int = i

		left := 2*i + 1
		if left < len(m.storage) && m.less(m.storage[left], m.storage[smallest]) {
			smallest = left
		}

		right := 2*i + 2
		if right < len(m.storage) && m.less(m.storage[right], m.storage[smallest]) {
			smallest = right
		}

		if smallest == i || m.less(m.storage[i], m.storage[smallest]) {
			break
		}

		m.storage[i], m.storage[smallest] = m.storage[smallest], m.storage[i]
		m.search[m.storage[i]], m.search[m.storage[smallest]] = m.search[m.storage[smallest]], m.search[m.storage[i]]

		i = smallest
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
