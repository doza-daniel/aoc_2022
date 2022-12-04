package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	game, err := parseInput("input")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, r := range game {
		total += r.getScore()
	}

	fmt.Println(total)
}

type symbol int

const (
	rock symbol = iota
	paper
	scissors
)

func fromStr(sym string) symbol {
	switch sym {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		panic("unknown symbol: " + sym)
	}
}

func fromStr2(opponent symbol, sym string) symbol {
	switch {
	case opponent == rock && sym == "X":
		return scissors
	case opponent == rock && sym == "Y":
		return rock
	case opponent == rock && sym == "Z":
		return paper

	case opponent == paper && sym == "X":
		return rock
	case opponent == paper && sym == "Y":
		return paper
	case opponent == paper && sym == "Z":
		return scissors

	case opponent == scissors && sym == "X":
		return paper
	case opponent == scissors && sym == "Y":
		return scissors
	case opponent == scissors && sym == "Z":
		return rock
	default:
		panic("panik :o")
	}
}

type round struct {
	player   symbol
	opponent symbol
}

func (r round) getScore() int {
	loss, draw, win := 0, 3, 6

	if r.player == rock {
		score := 1
		switch r.opponent {
		case rock:
			return score + draw
		case paper:
			return score + loss
		case scissors:
			return score + win
		}
	} else if r.player == paper {
		score := 2
		switch r.opponent {
		case rock:
			return score + win
		case paper:
			return score + draw
		case scissors:
			return score + loss
		}
	} else if r.player == scissors {
		score := 3
		switch r.opponent {
		case rock:
			return score + loss
		case paper:
			return score + win
		case scissors:
			return score + draw
		}
	}

	panic("abudabi")
}

func parseInput(filename string) ([]round, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rounds := []round{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		split := strings.Split(line, " ")

		if len(split) != 2 {
			return nil, fmt.Errorf("malformed input")
		}

		opponent := fromStr(split[0])
		r := round{opponent: opponent, player: fromStr2(opponent, split[1])}
		rounds = append(rounds, r)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rounds, nil
}
