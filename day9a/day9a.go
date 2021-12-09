package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	floor := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		floor = append(floor, row)
	}

	width := len(floor[0])
	height := len(floor)
	risk := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if isLow(floor, width, height, x, y) {
				risk += floor[y][x]+1
			}
		}
	}
	fmt.Println(risk)
}

func isLow(floor [][]int, width, height, x, y int) bool {
	here := floor[y][x]
	if x > 0 && floor[y][x-1] <= here {
		return false
	}
	if x < width - 1 && floor[y][x+1] <= here {
		return false
	}
	if y > 0 && floor[y-1][x] <= here {
		return false
	}
	if y < height -1 && floor[y+1][x] <= here {
		return false
	}
	return true
}