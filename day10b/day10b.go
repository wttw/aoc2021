package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var scores []int
	for scanner.Scan() {
		score := handle(scanner.Text())
		if score > 0 {
			//fmt.Printf("%s %d\n", scanner.Text(), score)
			scores = append(scores, score)
		}
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
	sort.Ints(scores)
	fmt.Println(scores[(len(scores)-1)/2])
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
		return 0 // invalid line
	}
	// Incomplete line
	ret := 0
	// Walk the stack backwards to get the missing characters
	for i := len(stack) - 1; i >= 0; i-- {
		switch stack[i] {
		case ')':
			ret = ret*5 + 1
		case ']':
			ret = ret*5 + 2
		case '}':
			ret = ret*5 + 3
		case '>':
			ret = ret*5 + 4
		}
	}
	return ret
}
