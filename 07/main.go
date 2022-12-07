package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

type node struct {
	isDir bool
	size  int

	entries map[string]*node
}

func main() {
	root, err := parseDirStructure("input")
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePartOne(root))
	fmt.Println(solvePartTwo(root))
}

func parseDirStructure(filename string) (*node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var root *node = &node{
		isDir:   true,
		entries: make(map[string]*node),
		size:    0,
	}

	var current *node = root

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		if line[0] == '$' {
			var command, param string
			fmt.Fscanf(strings.NewReader(line), "$ %s %s", &command, &param)
			switch command {
			case "cd":
				if param == "/" {
					current = root
				} else {
					current = current.entries[param]
				}
			case "ls":
			default:
				panic("unreachable")
			}
		} else {
			var nodeName string
			var fsNode *node

			if strings.HasPrefix(line, "dir") {
				fmt.Fscanf(strings.NewReader(line), "dir %s", &nodeName)

				fsNode = &node{isDir: true, entries: make(map[string]*node)}
			} else {
				var size int
				fmt.Fscanf(strings.NewReader(line), "%d %s", &size, &nodeName)

				fsNode = &node{size: size, entries: make(map[string]*node)}
			}

			fsNode.entries[".."] = current
			current.entries[nodeName] = fsNode
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return root, nil
}

func solvePartTwo(root *node) int {
	filesystemSize := 70000000
	unusedSpaceNeeded := 30000000

	dirSizes := make(map[string]int)
	usedSpace := calculateDirSize(root, "/", dirSizes)
	unusedSpace := filesystemSize - usedSpace

	min := filesystemSize
	for _, size := range dirSizes {
		if unusedSpace+size >= unusedSpaceNeeded {
			if size < min {
				min = size
			}
		}
	}
	return min
}

func solvePartOne(root *node) int {
	atMost := 100000

	dirSizes := make(map[string]int)
	calculateDirSize(root, "/", dirSizes)

	sum := 0
	for _, dirSize := range dirSizes {
		if dirSize <= atMost {
			sum += dirSize
		}
	}
	return sum
}

func calculateDirSize(root *node, fullPath string, dirSizes map[string]int) int {
	size := 0
	for name, nd := range root.entries {
		if !nd.isDir {
			size += nd.size
		} else if name != ".." {
			size += calculateDirSize(nd, path.Join(fullPath, name), dirSizes)
		}
	}

	dirSizes[fullPath] = size
	return size
}
