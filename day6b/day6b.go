package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	population := make([]int, 10)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		timer, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		population[timer]++
	}
	for day := 1; day <= 256; day++ {
		population[7] += population[0]
		population[9] = population[0]
		for i := 0; i < 9; i++ {
			population[i] = population[i+1]
		}
		population[9] = 0
		//count := 0
		//for _, v := range population {
		//	count += v
		//}
		//fmt.Printf("%d: %v %d\n", day, population, count)
	}
	count := 0
	for _, v := range population {
		count += v
	}
	fmt.Println(count)
}
