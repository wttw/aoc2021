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
	if l.x1 > l.x2 {
		l.x1, l.x2 = l.x2, l.x1
	}
	if l.y1 > l.y2 {
		l.y1, l.y2 = l.y2, l.y1
	}
	return nil
}

func (l Line) Vertical() bool {
	return l.x1 == l.x2
}

func (l Line) Horizontal() bool {
	return l.y1 == l.y2
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
		if line.Vertical() || line.Horizontal() {
			lines = append(lines, line)
		}
	}

	// Sigh. Brute force it.
	intersections := 0
	seabed := map[int]int{}
	for _, line := range lines {
		if line.Vertical() {
			//fmt.Printf("vline %v\n", line)
			for y := line.y1; y <= line.y2; y++ {
				p := line.x1 + 1000*y
				seabed[p]++
			}
			continue
		}
		if line.Horizontal() {
			//fmt.Printf("hline: %v\n", line)
			for x := line.x1; x <= line.x2; x++ {
				p := x + 1000*line.y1

				seabed[p]++
			}
		}
		// Diagonal
	}
	for _, v := range seabed {
		if v > 1 {
			intersections++
		}
	}
	fmt.Println(intersections)
}
