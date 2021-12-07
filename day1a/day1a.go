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
	for i, v := range depths[1:] {
		if v > depths[i] {
			deeper++
		}
	}
	fmt.Println(deeper)
}
