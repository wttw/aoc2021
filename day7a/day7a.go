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

	// Record how many are to the left and right of here
	count := 0
	leftmostCost := 0 // cost to move to lane 0
	for i := 0; i <= max; i++ {
		leftmostCost += i * lanes[i].Here
		lanes[i].Left = count
		count += lanes[i].Here
		lanes[i].Right = total - count
	}

	costToHere := leftmostCost
	lanes[0].Cost = costToHere
	bestCost := 1000000
	for i := 1; i <= max; i++ {
		costToHere += lanes[i].Left
		costToHere -= lanes[i].Right
		costToHere -= lanes[i].Here
		lanes[i].Cost = costToHere
		if costToHere < bestCost {
			bestCost = costToHere
		}
	}
	//f
	fmt.Println(bestCost)
}
