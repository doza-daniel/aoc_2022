package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	sensor, beacon Point
}

type Point struct {
	x, y int
}

func main() {
	pairs, err := parse("input")
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePartOne(pairs))
	// fmt.Println(solvePartTwo(pairs))
}

func solvePartOne(pairs []Pair) int {
	return solve(pairs, 2000000)
}

func solvePartTwo(pairs []Pair) int {
	for i := 0; i <= 20; i++ {
		solve(pairs, i)
	}

	return 0
}

func bounds(a Point) bool {
	return a.x >= 0 && a.x <= 4000000 && a.y >= 0 && a.y <= 4000000
}

func solve(pairs []Pair, target int) int {
	result := 0

	visited := make(map[Point]bool)

	for _, pair := range pairs {
		s := pair.sensor
		b := pair.beacon

		visited[b] = true

		l := Point{x: s.x, y: target}

		d := manhattan(s, b)
		f := manhattan(s, l)

		if d < f {
			continue
		}

		n := d - f

		for i := 0; i <= n; i++ {
			left := Point{x: s.x - i, y: target}
			right := Point{x: s.x + i, y: target}

			if !visited[left] {
				result++
				visited[left] = true
			}

			if !visited[right] {
				result++
				visited[right] = true
			}
		}
	}

	return result
}

func manhattan(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func parse(filename string) ([]Pair, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := []Pair{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var sensor, beacon Point

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		result = append(result, Pair{sensor: sensor, beacon: beacon})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
