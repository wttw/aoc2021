package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Line struct {
	x1, y1, x2, y2 int
}

func (l *Line) Scan(fields []string) error {
	var err error
	l.x1, err = strconv.Atoi(fields[0])
	if err != nil {
		return err
	}
	l.y1, err = strconv.Atoi(fields[1])
	if err != nil {
		return err
	}
	l.x2, err = strconv.Atoi(fields[2])
	if err != nil {
		return err
	}
	l.y2, err = strconv.Atoi(fields[3])
	if err != nil {
		return err
	}
	return nil
}

func (l Line) Vertical() bool {
	return l.x1 == l.x2
}

func (l Line) Horizontal() bool {
	return l.y1 == l.y2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sgn(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func main() {
	var lines []Line
	inputRe := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matches := inputRe.FindStringSubmatch(scanner.Text())
		if matches == nil {
			log.Fatalf("failed to match '%s'", scanner.Text())
		}
		var line Line
		err := line.Scan(matches[1:])
		if err != nil {
			log.Fatalf("failed to scan '%s': %v", scanner.Text(), err)
		}
		lines = append(lines, line)
	}

	// Sigh. Brute force it.
	intersections := 0
	seabed := map[int]int{}
	for _, line := range lines {
		xdir := sgn(line.x2 - line.x1)
		ydir := sgn(line.y2 - line.y1)
		x := line.x1
		y := line.y1
		for {
			//fmt.Println(line, x, y)
			p := x + 1000*y
			seabed[p]++
			if x == line.x2 && y == line.y2 {
				break
			}
			x += xdir
			y += ydir
		}
	}
	for _, v := range seabed {
		if v > 1 {
			intersections++
		}
	}
	fmt.Println(intersections)
}
