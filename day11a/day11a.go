package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var klowder [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		klowder = append(klowder, row)
	}
	flashes := 0
	for i := 1; i <= 100; i++ {
		flashes += step(klowder)
	}
	fmt.Println(flashes)
}

func dump(k [][]int) {
	for _, row := range k {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}

type loc struct {
	x, y int
}

func step(k [][]int) int {
	width := len(k[0])
	height := len(k)
	queue := []loc{}
	flashes := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := k[y][x]
			v++
			k[y][x] = v
			if v == 10 {
				queue = append(queue, loc{x, y})
			}
		}
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		flashes++
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				if x >= 0 && x < width && y >= 0 && y < height {
					v := k[y][x]
					v++
					k[y][x] = v
					if v == 10 {
						queue = append(queue, loc{x, y})
					}
				}
			}
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if k[y][x] >= 10 {
				k[y][x] = 0
			}
		}
	}
	return flashes
}
