package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		total += handle(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
	fmt.Println(total)
}

func handle(in string) int {
	flip := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
		'<': '>',
	}
	stack := []rune{}
	for _, c := range in {
		expecting, ok := flip[c]
		if ok {
			// it's an open symbol
			stack = append(stack, expecting)
			continue
		}
		// Must be a closing symbol, pop what we're expecting
		want := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if c == want {
			continue
		}
		switch c {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		default:
			log.Fatalf("unexpected char '%c' in %s\n", c, in)
		}
	}
	return 0
}
