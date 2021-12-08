package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// 0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....
//
// 5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

const all = "abcdefg"

var digits = []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}


func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 15 || fields[10] != "|" {
			log.Fatalf("Bad input: '%s'", scanner.Text())
		}
		sum += handle(fields[:10], fields[11:])
	}
	fmt.Println(sum)
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig string, p []int) string {
	result := []uint8(orig)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return string(result)
}

// handle generates all possible permutations and sees which work
func handle(inputs []string, code []string) int {
	digitCheck := map[string]int{}
	for i, d := range digits {
		digitCheck[d] = i
	}

	var mappings []string
	Perm:
	for p := make([]int, len(all)); p[0] < len(p); nextPerm(p) {
		mapping := getPerm(all, p)
		for _, v := range inputs {
			//fmt.Printf("%s: %s -> %s\n", mapping, v, translate(v, mapping))
			_, ok := digitCheck[translate(v, mapping)]
			if !ok {
				continue Perm
			}
		}
		// We have a possible mapping
		mappings = append(mappings, mapping)
	}
	if len(mappings) != 1 {
		log.Fatalf("found %d mappings for %v", len(mappings), inputs)
	}
	ret := 0
	for _, c := range code {
		i, ok := digitCheck[translate(c, mappings[0])]
		if !ok {
			log.Fatalf("Failed to map to digit %v | %v", inputs, code)
		}
		ret = ret * 10 + i
	}
	return ret
}

func translate(in, mapping string) string {
	ret := make([]uint8, len(in))
	for i, v := range in {
		ret[i] = mapping[v - 'a']
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return string(ret)
}