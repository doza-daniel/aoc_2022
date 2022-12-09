package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction uint

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Move struct {
	direction Direction
	steps     int
}

type Pos struct {
	x, y int
}

func main() {
	moves, err := parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartOne(moves))
	fmt.Println(solvePartTwo(moves))
}

func solvePartOne(moves []Move) int {
	return solve(moves, 2)
}

func solvePartTwo(moves []Move) int {
	return solve(moves, 10)
}

func solve(moves []Move, snekLen int) int {
	snek := make([]Pos, snekLen)

	total := 0
	visited := make(map[Pos]bool)

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			tail := snek[snekLen-1]
			if !visited[tail] {
				visited[tail] = true
				total++
			}

			moveHead(&snek[0], m.direction)
			for i := 0; i < snekLen-1; i++ {
				moveTail(&snek[i+1], &snek[i])
			}
		}
	}

	return total
}

func parse(input string) ([]Move, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	moves := make([]Move, 0)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		var direction rune
		var steps int
		fmt.Fscanf(strings.NewReader(line), "%c %d", &direction, &steps)

		m := Move{parseDirection(direction), steps}
		moves = append(moves, m)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}

func parseDirection(direction rune) Direction {
	switch direction {
	case 'U':
		return Up
	case 'D':
		return Down
	case 'L':
		return Left
	case 'R':
		return Right
	default:
		panic(fmt.Sprintf("unknown direction %c", direction))
	}
}

func moveHead(node *Pos, where Direction) {
	switch where {
	case Up:
		node.y++
	case Down:
		node.y--
	case Left:
		node.x--
	case Right:
		node.x++
	}
}

func moveTail(tail, head *Pos) {
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return
	}

	if tail.x == head.x {
		if tail.y < head.y {
			tail.y++
		} else {
			tail.y--
		}
		return
	} else if tail.y == head.y {
		if tail.x < head.x {
			tail.x++
		} else {
			tail.x--
		}
		return
	}

	if tail.x < head.x {
		tail.x++
	} else {
		tail.x--
	}

	if tail.y < head.y {
		tail.y++
	} else {
		tail.y--
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
