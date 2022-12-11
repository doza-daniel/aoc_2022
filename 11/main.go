package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	throwIfTrue  int
	throwIfFalse int
	arg          int
	div          int
	op           byte
	self         bool
	items        []int
}

func main() {
	monkeys, err := parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartOne(monkeys))

	monkeys, err = parse("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(solvePartTwo(monkeys))
}

func solvePartOne(monkeys []*Monkey) int {
	var nRounds int = 20

	var inspections []int = make([]int, len(monkeys))

	for round := 0; round < nRounds; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[i]++

				var arg int = monkey.arg
				if monkey.self {
					arg = item
				}

				switch monkey.op {
				case '+':
					item += arg
				case '*':
					item *= arg
				}

				item /= 3

				var throwTo *Monkey
				if item%monkey.div == 0 {
					throwTo = monkeys[monkey.throwIfTrue]
				} else {
					throwTo = monkeys[monkey.throwIfFalse]
				}
				throwTo.items = append(throwTo.items, item)
			}

			monkey.items = []int{}
		}
	}

	sort.Sort(sort.IntSlice(inspections))

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func solvePartTwo(monkeys []*Monkey) int {
	var nRounds int = 10000

	var inspections []int = make([]int, len(monkeys))

	var gigaMod int = 1
	for _, monkey := range monkeys {
		gigaMod *= monkey.div
	}

	for round := 0; round < nRounds; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[i]++

				var arg int = monkey.arg
				if monkey.self {
					arg = item
				}

				switch monkey.op {
				case '+':
					item = (item % gigaMod) + (arg % gigaMod)
				case '*':
					item = (item % gigaMod) * (arg % gigaMod)
				}

				var throwTo *Monkey
				if item%monkey.div == 0 {
					throwTo = monkeys[monkey.throwIfTrue]
				} else {
					throwTo = monkeys[monkey.throwIfFalse]
				}
				throwTo.items = append(throwTo.items, item)
			}

			monkey.items = []int{}
		}
	}

	sort.Sort(sort.IntSlice(inspections))

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func parse(filename string) ([]*Monkey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var monkeys []*Monkey = make([]*Monkey, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey") {
			var monkey Monkey

			for i := 0; scanner.Scan() && i < 5; i++ {
				line = scanner.Text()
				if strings.HasPrefix(line, "  Starting") {
					trimmed := strings.TrimPrefix(line, "  Starting items: ")
					split := strings.Split(trimmed, ", ")
					for _, s := range split {
						i, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, err
						}

						monkey.items = append(monkey.items, int(i))
					}
				}

				if strings.HasPrefix(line, "  Operation") {
					trimmed := strings.TrimPrefix(line, "  Operation: new = old ")

					monkey.op = trimmed[0]
					if trimmed[0] == '+' {
						trimmed = trimmed[2:]
						if trimmed == "old" {
							monkey.self = true
						} else {
							x, err := strconv.ParseInt(trimmed, 10, 32)
							if err != nil {
								return nil, err
							}

							monkey.arg = int(x)
						}
					} else if trimmed[0] == '*' {
						trimmed = trimmed[2:]
						if trimmed == "old" {
							monkey.self = true
						} else {
							x, err := strconv.ParseInt(trimmed, 10, 32)
							if err != nil {
								return nil, err
							}

							monkey.arg = int(x)
						}
					}
				}

				if strings.HasPrefix(line, "  Test:") {
					trimmed := strings.TrimPrefix(line, "  Test: divisible by ")
					i, err := strconv.ParseInt(trimmed, 10, 64)
					if err != nil {
						return nil, err
					}

					monkey.div = int(i)
				}

				if strings.HasPrefix(line, "    If true:") {
					trimmed := strings.TrimPrefix(line, "    If true: throw to monkey ")
					i, err := strconv.ParseInt(trimmed, 10, 32)
					if err != nil {
						return nil, err
					}

					monkey.throwIfTrue = int(i)
				}

				if strings.HasPrefix(line, "    If false:") {
					trimmed := strings.TrimPrefix(line, "    If false: throw to monkey ")
					i, err := strconv.ParseInt(trimmed, 10, 32)
					if err != nil {
						return nil, err
					}

					monkey.throwIfFalse = int(i)
				}
			}

			monkeys = append(monkeys, &monkey)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return monkeys, nil
}
