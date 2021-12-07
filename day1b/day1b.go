package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var depths []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			depth, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("'%s': %v\n", line, err)
			}
			depths = append(depths, depth)
		}
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}

	deeper := 0
	sum := depths[0] + depths[1] + depths[2]
	for i, v := range depths[3:] {
		newsum := sum + v - depths[i]
		if newsum > sum {
			deeper++
		}
		sum = newsum
	}
	fmt.Println(deeper)
}
