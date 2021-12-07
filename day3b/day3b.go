package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		input = append(input, line)
	}

	oxy := findEntry(input, func(zeros, ones int) bool {
		return ones >= zeros
	})
	co2 := findEntry(input, func(zeros, ones int) bool {
		return ones < zeros
	})

	fmt.Println(oxy * co2)
}

func findEntry(input []string, compare func(zeros, ones int) bool) int64 {
	f := input
	for bit := 0; bit < len(f[0]); bit++ {
		zeros, ones := countBits(f, bit)
		want := '0'
		if compare(zeros, ones) {
			want = '1'
		}

		f = filter(f, bit, want)
		if len(f) == 1 {
			ret, err := strconv.ParseInt(f[0], 2, 64)
			if err != nil {
				log.Fatalf("bad binary '%s': %v", f[0], err)
			}
			return ret
		}
		if len(f) == 0 {
			log.Fatalf("ran out of entries at bit %d", bit)
		}
	}
	log.Fatalf("nothing found")
	return 0
}

func countBits(in []string, bit int) (int, int) {
	ones := 0
	zeros := 0
	for _, s := range in {
		switch s[bit] {
		case '1':
			ones++
		case '0':
			zeros++
		default:
			log.Fatalf("bad input '%s'", s)
		}
	}
	return zeros, ones
}

func filter(in []string, bit int, want rune) []string {
	var filtered []string
	for _, s := range in {
		if s[bit] == uint8(want) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
