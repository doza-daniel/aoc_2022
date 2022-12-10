package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	isNoop bool
	cycles int
	arg    int
}

func main() {
	instructions, err := parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartOne(instructions))
	solvePartTwo(instructions)
}

func solvePartOne(instructions []Instruction) int {
	breakpoints := []int{20, 60, 100, 140, 180, 220}

	var cycles int = 0
	var sum int = 0

	cyclesCh := runCPU(instructions)
	for x := range cyclesCh {
		cycles++
		if len(breakpoints) > 0 && cycles == breakpoints[0] {
			sum += x * breakpoints[0]
			breakpoints = breakpoints[1:]
		}
	}

	return sum
}

func solvePartTwo(instructions []Instruction) {
	height, width, pos := 6, 40, 0

	cyclesCh := runCPU(instructions)
	for x := range cyclesCh {
		i, j := pos/width%height, pos%width

		if i == 0 && j == 0 {
			fmt.Print("\033[H\033[2J")
		}

		if pos > 0 && j == 0 {
			fmt.Println()
		}

		if j == x-1 || j == x || j == x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		pos++
	}

	fmt.Println()
}

// runCPU emits the value of x regiester at the start of each cycle
func runCPU(instructions []Instruction) <-chan int {
	c := make(chan int)
	x := 1

	go func() {
		defer close(c)

		for _, instruction := range instructions {
			for i := 0; i < instruction.cycles; i++ {
				c <- x
			}

			if !instruction.isNoop {
				x += instruction.arg
			}
		}
	}()

	return c
}

func parse(filename string) ([]Instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	instructions := []Instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var i Instruction
		if line == "noop" {
			i = Instruction{isNoop: true, cycles: 1}
		} else {
			var arg int
			fmt.Sscanf(line, "addx %d", &arg)
			i = Instruction{isNoop: false, cycles: 2, arg: arg}
		}

		instructions = append(instructions, i)
	}

	return instructions, nil
}
