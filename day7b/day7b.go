package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lane struct {
	Left int
	Right int
	Here int
	Cost int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	max := 0
	min := 1000000
	total := 0
	lanes := make([]Lane, 10000)
	for _, v := range strings.Split(scanner.Text(), ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
		total++
		lanes[i].Here++
	}


	bestCost := 1000000000
	for target := min; target <= max; target++ {
		cost := 0
		for i := min; i <= max; i++ {
			dist := target - i
			if dist < 0 {
				dist = -dist
			}
			cost += (dist * (dist+1) * lanes[i].Here) / 2
		}
		if cost < bestCost {
			bestCost = cost
		}
	}
	fmt.Println(bestCost)
}
