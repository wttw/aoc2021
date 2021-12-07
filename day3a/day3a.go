package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var bitcounts []int
	linecount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if bitcounts == nil {
			bitcounts = make([]int, len(line))
		}
		linecount++
		for i, v := range line {
			switch v {
			case '1':
				bitcounts[i]++
			case '0':
			default:
				log.Fatalf("bad line: '%s'", line)
			}
		}
	}
	epsilon := 0
	gamma := 0
	for _, v := range bitcounts {
		gamma <<= 1
		epsilon <<= 1
		if v*2 > linecount {
			gamma ^= 1
		} else {
			epsilon ^= 1
		}
	}
	fmt.Println(gamma * epsilon)
}
