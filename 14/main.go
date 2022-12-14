package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	walls, err := parse("input")
	if err != nil {
		panic(err)
	}

	_ = walls

	// fmt.Println(solvePartOne(walls))
	fmt.Println(solvePartTwo(walls))
}

func solvePartOne(walls [][]Point) int {
	result := 0

	lowestPoint := findLowestHorizontalWall(walls)
	startingPoint := Point{y: 0, x: 500}
	landed := make(map[Point]bool)

	for {
		var stalled bool

		for grain := startingPoint; grain.y < lowestPoint && !stalled; {

			bottom := Point{x: grain.x, y: grain.y + 1}
			if !isWall(walls, bottom) && !landed[bottom] {
				grain = bottom
				continue
			}

			bottomLeft := Point{x: grain.x - 1, y: grain.y + 1}
			if !isWall(walls, bottomLeft) && !landed[bottomLeft] {
				grain = bottomLeft
				continue
			}

			bottomRight := Point{x: grain.x + 1, y: grain.y + 1}
			if !isWall(walls, bottomRight) && !landed[bottomRight] {
				grain = bottomRight
				continue
			}

			stalled = true
			landed[grain] = true
		}

		if stalled {
			result++
		} else {
			break
		}
	}

	return result
}

func solvePartTwo(walls [][]Point) int {
	result := 0

	floor := findLowestHorizontalWall(walls) + 2
	startingPoint := Point{y: 0, x: 500}
	landed := make(map[Point]bool)
	walls = append(walls, []Point{{x: -100000, y: floor}, {x: math.MaxInt32, y: floor}})

	for {
		var stalled bool

		for grain := startingPoint; !stalled; {

			bottom := Point{x: grain.x, y: grain.y + 1}
			if !isWall(walls, bottom) && !landed[bottom] {
				grain = bottom
				continue
			}

			bottomLeft := Point{x: grain.x - 1, y: grain.y + 1}
			if !isWall(walls, bottomLeft) && !landed[bottomLeft] {
				grain = bottomLeft
				continue
			}

			bottomRight := Point{x: grain.x + 1, y: grain.y + 1}
			if !isWall(walls, bottomRight) && !landed[bottomRight] {
				grain = bottomRight
				continue
			}

			stalled = true
			landed[grain] = true
			if grain == startingPoint {
				result++
				return result
			}
		}

		if stalled {
			result++
		}
	}
}

func findLowestHorizontalWall(walls [][]Point) int {
	max := 0

	for _, wall := range walls {
		for i := 0; i < len(wall)-1; i++ {
			if isHorizontalLine(wall[i], wall[i+1]) && wall[i].y > max {
				max = wall[i].y
			}
		}
	}

	return max
}

func isHorizontalLine(start, end Point) bool {
	return start.y == end.y
}

func isWall(walls [][]Point, p Point) bool {
	for _, wall := range walls {
		for i := 0; i < len(wall)-1; i++ {
			if onLine(wall[i], wall[i+1], p) {
				return true
			}
		}
	}
	return false
}

func onLine(start, end, p Point) bool {
	return dist(start, p)+dist(p, end) == dist(start, end)
}

func dist(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func parse(filename string) ([][]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := [][]Point{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		points := []Point{}
		for _, coord := range coords {
			i := strings.Index(coord, ",")
			x, err := strconv.ParseInt(coord[:i], 10, 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseInt(coord[i+1:], 10, 64)
			if err != nil {
				return nil, err
			}
			points = append(points, Point{int(x), int(y)})
		}
		result = append(result, points)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
