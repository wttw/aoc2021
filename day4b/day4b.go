package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Values []int
	Set    []bool
	Won    bool
}

func main() {
	var boards []*Board
	var values []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	draws := strings.Split(scanner.Text(), ",")

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		for _, v := range strings.Fields(line) {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			values = append(values, val)
		}
		if len(values) == 25 {
			boards = append(boards, &Board{Values: values, Set: make([]bool, len(values))})
			values = nil
		}
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}

	boardsWon := 0

	for _, d := range draws {
		draw, err := strconv.Atoi(d)
		if err != nil {
			log.Fatalln(err)
		}
		for _, board := range boards {
			if board.Won {
				continue
			}
			board.Play(draw)
			if board.Bingo() {
				board.Won = true
				boardsWon++
				if boardsWon == len(boards) {
					sum := 0
					for i, v := range board.Values {
						if !board.Set[i] {
							sum += v
						}
					}
					fmt.Println(sum * draw)
					os.Exit(0)
				}
			}
		}
	}
}

func (b *Board) Play(val int) {
	for i, v := range b.Values {
		if v == val {
			b.Set[i] = true
		}
	}
}

func (b *Board) Bingo() bool {
Row:
	for i := 0; i < 5; i++ {
		for j := 0; j < 25; j += 5 {
			if !b.Set[i+j] {
				continue Row
			}
		}
		return true
	}
Column:
	for i := 0; i < 25; i += 5 {
		for j := 0; j < 5; j++ {
			if !b.Set[i+j] {
				continue Column
			}
		}
		return true
	}
	return false
}
