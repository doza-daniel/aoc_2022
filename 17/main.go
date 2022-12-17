package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	moves, err := parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartOne(moves))
}

type Point struct {
	x, y int
}

func solvePartOne(moves string) int {
	grid := make(map[Point]bool)
	height := 0

	var shapes []byte = []byte{'-', '+', 'L', 'I', 'o'}

	var j int = 0
	for i := 0; i < 2022; i++ {
		shape := shapes[i%len(shapes)]
		pos := Point{x: height + 3, y: 2}

		for {
			direction := moves[j%len(moves)]
			j++

			pprint(shape, pos, grid, height)
			move(shape, direction, &pos, grid)
			pprint(shape, pos, grid, height)
			if !move(shape, 'v', &pos, grid) {
				x := halt(shape, pos, grid)
				if x+1 > height {
					height = x + 1
				}
				// pprint(shape, pos, grid, height)
				break
			}
		}
	}

	return height
}

func halt(shape byte, pos Point, grid map[Point]bool) int {
	if shape == '-' {
		for i := 0; i < 4; i++ {
			grid[Point{pos.x, pos.y + i}] = true
		}

		return pos.x
	}

	if shape == '+' {
		for i := 0; i < 3; i++ {
			grid[Point{pos.x + 1, pos.y + i}] = true
			grid[Point{pos.x + i, pos.y + 1}] = true
		}

		return pos.x + 2
	}

	if shape == 'L' {
		for i := 0; i < 3; i++ {
			grid[Point{pos.x, pos.y + i}] = true
			grid[Point{pos.x + i, pos.y + 2}] = true
		}

		return pos.x + 2
	}

	if shape == 'I' {
		for i := 0; i < 4; i++ {
			grid[Point{pos.x + i, pos.y}] = true
		}

		return pos.x + 3
	}

	if shape == 'o' {
		for i := 0; i < 4; i++ {
			grid[Point{pos.x + (i / 2), pos.y + (i % 2)}] = true
		}

		return pos.x + 1
	}

	panic("unreachable")
}

func move(shape, direction byte, pos *Point, grid map[Point]bool) bool {
	fmt.Fprintf(os.Stderr, "%c, %c, %+v\n", shape, direction, *pos)
	if shape == '-' {
		if direction == '<' {
			if pos.y <= 0 || grid[Point{x: pos.x, y: pos.y - 1}] {
				return false
			}

			pos.y--
			return true
		}

		if direction == '>' {
			if pos.y+4 >= 7 || grid[Point{pos.x, pos.y + 4}] {
				return false
			}

			pos.y++
			return true
		}

		if direction == 'v' {
			if pos.x <= 0 {
				return false
			}

			for i := 0; i < 4; i++ {
				p := Point{x: pos.x - 1, y: pos.y + i}
				if grid[p] {
					return false
				}
			}

			pos.x--
			return true
		}
	}

	if shape == '+' {
		left := Point{x: pos.x + 1, y: pos.y}
		bot := Point{x: pos.x, y: pos.y + 1}
		top := Point{x: pos.x + 2, y: pos.y + 1}
		right := Point{x: pos.x + 1, y: pos.y + 2}

		if direction == '<' {
			if left.y <= 0 {
				return false
			}
			for _, p := range []Point{left, right, bot, top} {
				if grid[Point{x: p.x, y: p.y - 1}] {
					return false
				}
			}

			pos.y--
			return true
		}

		if direction == '>' {
			if right.y+1 >= 7 {
				return false
			}

			for _, p := range []Point{left, right, bot, top} {
				if grid[Point{x: p.x, y: p.y + 1}] {
					return false
				}
			}

			pos.y++
			return true
		}

		if direction == 'v' {
			if bot.x <= 0 {
				return false
			}

			for _, p := range []Point{left, right, bot, top} {
				if grid[Point{x: p.x - 1, y: p.y}] {
					return false
				}
			}

			pos.x--
			return true
		}
	}

	if shape == 'L' {
		left := Point{x: pos.x, y: pos.y}
		right := Point{x: pos.x, y: pos.y + 2}
		bot := Point{x: pos.x, y: pos.y + 1}

		if direction == '<' {
			if left.y <= 0 {
				return false
			}

			for i := 0; i < 3; i++ {
				p := Point{x: pos.x + i, y: pos.y + 1}
				if grid[p] {
					return false
				}
				p = left
				p.y = p.y - 1
				if grid[p] {
					return false
				}
			}

			pos.y--
			return true
		}

		if direction == '>' {
			if right.y+1 >= 7 {
				return false
			}

			for i := 0; i < 3; i++ {
				p := Point{x: pos.x + i, y: pos.y + 3}
				if grid[p] {
					return false
				}
			}

			pos.y++
			return true
		}

		if direction == 'v' {
			if bot.x <= 0 {
				return false
			}

			for i := 0; i < 3; i++ {
				if grid[Point{x: pos.x - 1, y: pos.y + i}] {
					return false
				}
			}

			pos.x--
			return true
		}
	}

	if shape == 'I' {
		if direction == '<' {
			if pos.y <= 0 {
				return false
			}

			for i := 0; i < 4; i++ {
				if grid[Point{x: pos.x + i, y: pos.y - 1}] {
					return false
				}
			}

			pos.y--
			return true
		}

		if direction == '>' {
			if pos.y+1 >= 7 {
				return false
			}

			for i := 0; i < 4; i++ {
				if grid[Point{x: pos.x + i, y: pos.y + 1}] {
					return false
				}
			}

			pos.y++
			return true
		}

		if direction == 'v' {
			if pos.x <= 0 || grid[Point{x: pos.x - 1, y: pos.y}] {
				return false
			}

			pos.x--
			return true
		}
	}

	if shape == 'o' {
		left := Point{x: pos.x, y: pos.y}
		right := Point{x: pos.x, y: pos.y + 1}
		bot := Point{x: pos.x, y: pos.y}

		if direction == '<' {
			if left.y <= 0 {
				return false
			}

			for i := 0; i < 2; i++ {
				if grid[Point{x: pos.x + i, y: pos.y - 1}] {
					return false
				}
			}

			pos.y--
			return true
		}

		if direction == '>' {
			if right.y+1 >= 7 {
				return false
			}

			for i := 0; i < 2; i++ {
				if grid[Point{x: pos.x + i, y: pos.y + 2}] {
					return false
				}
			}

			pos.y++
			return true
		}

		if direction == 'v' {
			if bot.x <= 0 {
				return false
			}

			for i := 0; i < 2; i++ {
				if grid[Point{x: pos.x - 1, y: pos.y + i}] {
					return false
				}
			}

			pos.x--
			return true
		}
	}

	fmt.Printf("%c %c %+v %v\n", shape, direction, *pos, grid[Point{}])
	panic("Unreachable")
}

func parse(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(b)), nil
}

func pprint(shape byte, pos Point, grid map[Point]bool, height int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
	bot := height + 5 - 20
	if bot < 0 {
		bot = 0
	}
	for i := height + 5; i >= bot; i-- {
		fmt.Printf("%5d ", i)
		for j := 0; j < 7; j++ {
			p := Point{x: i, y: j}
			if in(shape, pos, i, j) {
				fmt.Print("@")
			} else if grid[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func in(shape byte, pos Point, i, j int) bool {
	switch shape {
	case '-':
		return i == pos.x && j >= pos.y && j < pos.y+4
	case '+':
		if i == pos.x || i == pos.x+2 {
			return j == pos.y+1
		}
		if i == pos.x+1 {
			return j >= pos.y && j < pos.y+3
		}

	case 'L':
		if i == pos.x {
			return j >= pos.y && j < pos.y+3
		}
		if j == pos.y+2 {
			return i >= pos.x && i < pos.x+3
		}
	case 'I':
		return j == pos.y && i >= pos.x && i < pos.x+4

	case 'o':
		return i >= pos.x && i < pos.x+2 && j >= pos.y && j < pos.y+2
	}

	return false
}
