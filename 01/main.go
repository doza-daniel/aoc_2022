package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(content))

	max := make([]int, 3)
	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			insMax(max, current)
			current = 0
		} else {
			i, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			current += int(i)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner error:", err)
	}

	fmt.Println(sum(max))
}

func sum(list []int) int {
	sum := 0
	for _, n := range list {
		sum += n
	}
	return sum
}

func insMax(list []int, num int) {
	for i := 0; i < len(list); i++ {
		if list[i] < num {
			for j := len(list) - 1; j > i; j-- {
				list[j] = list[j-1]
			}
			list[i] = num
			return
		}
	}
}
