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
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fields := strings.Fields(line)
			if len(fields) != 2 {
				log.Fatalf("Wrong number of fields: '%s'", line)
			}
			delta, err := strconv.Atoi(fields[1])
			if err != nil {
				log.Fatalf("'%s': %v", line, err)
			}
			switch fields[0] {
			case "forward":
				x += delta
				depth += delta * aim
			case "down":
				aim += delta
			case "up":
				aim -= delta
			default:
				log.Fatalf("Bad line: '%s'", line)
			}
		}
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
	fmt.Println(x * depth)
}
