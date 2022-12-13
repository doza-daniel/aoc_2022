package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	left, right []interface{}
}

type Status byte

const (
	InOrder Status = iota
	NotInOrder
	Equal
)

func main() {
	pairs, err := parse("input")
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePartOne(pairs))
	fmt.Println(solvePartTwo(pairs))
}

func solvePartOne(pairs []Pair) int {
	sum := 0
	for i, p := range pairs {
		x := inOrder(p.left, p.right)
		if x != NotInOrder {
			sum += i + 1
		}
	}
	return sum
}

func solvePartTwo(pairs []Pair) int {
	packets := [][]interface{}{
		{[]interface{}{2}},
		{[]interface{}{6}},
	}

	for _, p := range pairs {
		packets = append(packets, p.left)
		packets = append(packets, p.right)
	}

	sort.Slice(packets, func(i, j int) bool {
		return inOrder(packets[i], packets[j]) == InOrder
	})

	result := 1
	for i, packet := range packets {
		if isDivider(packet) {
			result *= i + 1
		}
	}

	return result
}

func isDivider(packet []interface{}) bool {
	if len(packet) != 1 {
		return false
	}

	list, ok := packet[0].([]interface{})
	if !ok {
		return false
	}

	if len(list) != 1 {
		return false
	}

	n, ok := isNumber(list[0])

	return ok && (n == 6 || n == 2)
}

func inOrder(left, right []interface{}) Status {
	for i := 0; i < len(left); i++ {
		if i == len(right) {
			return NotInOrder
		}

		leftN, leftIsNumber := isNumber(left[i])
		rightN, rightIsNumber := isNumber(right[i])

		if leftIsNumber && rightIsNumber {
			if leftN > rightN {
				return NotInOrder
			}

			if leftN < rightN {
				return InOrder
			}

			continue
		}

		var newLeft, newRight []interface{}

		if !leftIsNumber && !rightIsNumber {
			newLeft = left[i].([]interface{})
			newRight = right[i].([]interface{})
		} else if leftIsNumber {
			newLeft = []interface{}{leftN}
			newRight = right[i].([]interface{})
		} else if rightIsNumber {
			newRight = []interface{}{rightN}
			newLeft = left[i].([]interface{})
		}

		status := inOrder(newLeft, newRight)
		if status != Equal {
			return status
		}
	}

	if len(left) == len(right) {
		return Equal
	}

	return InOrder
}

func isNumber(i interface{}) (int, bool) {
	n, ok := i.(float64)

	if !ok {
		m, ok := i.(int)
		return m, ok
	}

	return int(n), ok
}

func parseLine(line string) []interface{} {
	var data []interface{}

	if err := json.NewDecoder(strings.NewReader(line)).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func parse(filename string) ([]Pair, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := []Pair{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		left := scanner.Text()
		if len(left) == 0 {
			continue
		}

		scanner.Scan()
		right := scanner.Text()

		p := Pair{left: parseLine(left), right: parseLine(right)}

		result = append(result, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
