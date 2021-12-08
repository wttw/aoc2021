package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 15 || fields[10] != "|" {
			log.Fatalf("Bad input: '%s'", scanner.Text())
		}
		//handle(fields[:10], fields[11:])
		for _, v := range fields[11:] {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println(count)
}

/* Lots of junk I was playing with follows

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

type StateType int

const (
	Maybe StateType = iota
	Yes
	No
)

func (s StateType) String() string {
	switch s {
	case Maybe:
		return "?"
	case Yes:
		return "Y"
	case No:
		return "N"
	}
	log.Fatalln("Invalid state")
	return ""
}

type Mapping [][]StateType

func NewMapping() Mapping {
	ret := make([][]StateType, 7)
	for i := range ret {
		ret[i] = make([]StateType, 7)
	}
	return ret
}

func (m Mapping) String() string {
	var builder strings.Builder
	builder.WriteString(" abcdefg\n")
	for i, v := range m {
		builder.WriteString("abcdefg"[i : i+1])
		for _, c := range v {
			builder.WriteString(c.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

// Set a single in/out mapping
func (m Mapping) Set(input, output rune, val StateType) {
	m[input-'a'][output-'a'] = val
}

// Get a single in/out mapping
func (m Mapping) Get(input, output rune) StateType {
	return m[input-'a'][output-'a']
}

// SetPossibles - we know that input maps to output, but not the order
func (m Mapping) SetPossibles(input, output string) {
	notoutput := Invert(output)
	notinput := Invert(input)

	for _, i := range input {
		for _, j := range notoutput {
			m.Set(i, j, No)
		}
	}
	for _, i := range notinput {
		for _, j := range input {
			m.Set(i, j, No)
		}
	}
}

// TidyYes removes Maybes excluded by a conflicting Yes
func (m Mapping) TidyYes() bool {
	changed := false
	for _, i := range all {
		for _, j := range all {
			if m.Get(i, j) == Yes {
				fmt.Printf("Yes at %c %c\n", i, j)
				for _, ii := range all {
					if i == ii {
						continue
					}
					switch m.Get(ii, j) {
					case Yes:
						log.Fatalf("Contradicting Yes:\n%v", m)
					case No:
					case Maybe:
						changed = true
						m.Set(ii, j, No)
					}
				}
				for _, jj := range all {
					if j == jj {
						continue
					}
					switch m.Get(i, jj) {
					case Yes:
						log.Fatalf("Contradicting Yes:\n%v", m)
					case No:
					case Maybe:
						changed = true
						m.Set(i, jj, No)
					}
				}
			}
		}
	}
	return changed
}

func (m Mapping) TidyMaybes() bool {
	// Look for single Maybes
	for _, i := range all {
		maybeCount := 0
		var maybeAt rune
		for _, j := range all {
			if m.Get(i, j) == Maybe {
				maybeCount++
				maybeAt = j
			}
		}
		if maybeCount == 1 {
			fmt.Printf("Single maybe at %c %c\n%v\n\n", i, maybeAt, m)
			m.Set(i, maybeAt, Yes)
			return true
		}
	}

	for _, j := range all {
		maybeCount := 0
		var maybeAt rune
		for _, i := range all {
			if m.Get(i, j) == Maybe {
				maybeCount++
				maybeAt = i
			}
		}
		if maybeCount == 1 {
			fmt.Printf("Single maybe at %c %c\n%v\n\n", maybeAt, j, m)
			m.Set(maybeAt, j, Yes)
			return true
		}
	}
	return false
}

// Tidy eliminates states excluded
func (m Mapping) Tidy() {
	for {
		for m.TidyYes() {
		}
		if !m.TidyMaybes() {
			return
		}
	}
}

func (m Mapping) PossibleMatch(input, output string) {

}

// Invert returns all the characters a-g not in the input
func Invert(s string) string {
	ret := ""
	for _, c := range all {
		if !strings.ContainsRune(s, c) {
			ret += string(c)
		}
	}
	return ret
}

func handle(inputs []string, _ []string) {
	m := NewMapping()
	for _, v := range inputs {
		switch len(v) {
		case 2:
			m.SetPossibles(v, digits[1])
			fmt.Printf("%s -> %s:\n%v\n", v, digits[1], m)
		case 3:
			m.SetPossibles(v, digits[7])
			fmt.Printf("%s -> %s:\n%v\n", v, digits[7], m)
		case 4:
			m.SetPossibles(v, digits[4])
			fmt.Printf("%s -> %s:\n%v\n", v, digits[4], m)
		case 7:
			m.SetPossibles(v, digits[8])
			fmt.Printf("%s -> %s:\n%v\n", v, digits[8], m)
		}
	}
	fmt.Println(m)
	m.Tidy()
	fmt.Println(m)
	os.Exit(0)
}

func parse(s string) []int {
	ret := make([]int, len(s))
	for i, c := range s {
		ret[i] = int(c - 'a')
	}
	sort.Ints(ret)
	return ret
}

*/
