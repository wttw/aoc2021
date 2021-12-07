package main

import (
	"bufio"
	"os"
	"regexp"
)

type Line struct {
	x1, y1, x2, y2 int
}

func main() {
	inputRe := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		
	}

}
