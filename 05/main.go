package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stacks := make([][]rune, 9)
	stacks[0] = []rune{'N', 'B', 'D', 'T', 'V', 'G', 'Z', 'J'}
	stacks[1] = []rune{'S', 'R', 'M', 'D', 'W', 'P', 'F'}
	stacks[2] = []rune{'V', 'C', 'R', 'S', 'Z'}
	stacks[3] = []rune{'R', 'T', 'J', 'Z', 'P', 'H', 'G'}
	stacks[4] = []rune{'T', 'C', 'J', 'N', 'D', 'Z', 'Q', 'F'}
	stacks[5] = []rune{'N', 'V', 'P', 'W', 'G', 'S', 'F', 'M'}
	stacks[6] = []rune{'G', 'C', 'V', 'B', 'P', 'Q'}
	stacks[7] = []rune{'Z', 'B', 'P', 'N'}
	stacks[8] = []rune{'W', 'P', 'J'}

	g := game{stacks: stacks}

	solution, err := solve("input", &g)
	if err != nil {
		panic(err)
	}
	fmt.Println(solution)
}

type game struct {
	stacks [][]rune
}

func (g *game) move(count, from, to int) {
	for i := 0; i < count; i++ {
		k := len(g.stacks[from]) - 1 - i
		g.stacks[to] = append(g.stacks[to], g.stacks[from][k])
	}
	g.stacks[from] = g.stacks[from][:len(g.stacks[from])-count]
}

func (g *game) moveStacked(count, from, to int) {
	k := len(g.stacks[from]) - count
	g.stacks[to] = append(g.stacks[to], g.stacks[from][k:]...)
	g.stacks[from] = g.stacks[from][:len(g.stacks[from])-count]
}

func (g *game) top() string {
	s := make([]rune, len(g.stacks))
	for i, stack := range g.stacks {
		s[i] = stack[len(stack)-1]
	}
	return string(s)
}

func solve(filename string, g *game) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var from, to, count int
		fmt.Fscanf(strings.NewReader(line), "move %d from %d to %d", &count, &from, &to)
		g.moveStacked(count, from-1, to-1)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return g.top(), nil
}
