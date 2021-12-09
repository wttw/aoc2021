package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// our floor heights
	floor := [][]int{}

	// our map of basins
	basin := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		floor = append(floor, row)
		basin = append(basin, make([]int, len(line)))
	}

	width := len(floor[0])
	height := len(floor)
	basinCount := 0
	sizes := []int{}
	// Find all low points and flood fill a basin from each
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if isLow(floor, width, height, x, y) {
				basinCount++
				size := flood(floor, basin, width, height, x, y, basinCount)
				sizes = append(sizes, size)
			}
		}
	}

	// dump(floor)
	// fmt.Printf("\n")
	// dump(basin)
	// fmt.Println(sizes)

	// Sort sizes, descending, and return the product of the biggest 3
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

// isLow checks whether a particular point is lower than all adjacent points
func isLow(floor [][]int, width, height, x, y int) bool {
	here := floor[y][x]
	if x > 0 && floor[y][x-1] <= here {
		return false
	}
	if x < width-1 && floor[y][x+1] <= here {
		return false
	}
	if y > 0 && floor[y-1][x] <= here {
		return false
	}
	if y < height-1 && floor[y+1][x] <= here {
		return false
	}
	return true
}

type point struct {
	x int
	y int
}

// dump prints a floor or a basin to stdout
func dump(thing [][]int) {
	for _, row := range thing {
		for _, field := range row {
			fmt.Printf("%d ", field)
		}
		fmt.Printf("\n")
	}
}

// flood does a flood fill from a low point to all higher points < 9
func flood(floor, basin [][]int, width, height, x, y, num int) int {
	// slightly inefficient queue
	q := []point{{x: x, y: y}}
	ret := 0
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if basin[p.y][p.x] != 0 {
			continue
		}
		basin[p.y][p.x] = num
		ret++
		here := floor[p.y][p.x]
		if p.x > 0 {
			h := floor[p.y][p.x-1]
			if h != 9 && h > here && basin[p.y][p.x-1] == 0 {
				q = append(q, point{x: p.x - 1, y: p.y})
			}
		}
		if p.x < width-1 {
			h := floor[p.y][p.x+1]
			if h != 9 && h > here && basin[p.y][p.x+1] == 0 {
				q = append(q, point{x: p.x + 1, y: p.y})
			}
		}
		if p.y > 0 {
			h := floor[p.y-1][p.x]
			if h != 9 && h > here && basin[p.y-1][p.x] == 0 {
				q = append(q, point{x: p.x, y: p.y - 1})
			}
		}
		if p.y < height-1 {
			h := floor[p.y+1][p.x]
			if h != 9 && h > here && basin[p.y+1][p.x] == 0 {
				q = append(q, point{x: p.x, y: p.y + 1})
			}
		}
	}
	return ret
}
