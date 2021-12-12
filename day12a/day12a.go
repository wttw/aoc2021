package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Exits []string
}

type Path struct {
	Nodes []string
}

func main() {
	nodes := map[string]Node{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n := strings.Split(scanner.Text(), "-")
		addArc(nodes, n[0], n[1])
		addArc(nodes, n[1], n[0])
	}

	count := 0
	queue := []Path{{[]string{"start"}}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
	Path:
		for _, exit := range nodes[path.Nodes[len(path.Nodes)-1]].Exits {
			if exit == "start" {
				continue
			}
			if exit == "end" {
				count++
				// fmt.Printf("%d: %s-end\n", count, strings.Join(path.Nodes, "-"))
				// fmt.Printf("   %#v\n", path.Nodes)
				continue
			}
			if exit[0] >= 'a' {
				// Lower case == small cave == can't revisit
				for _, v := range path.Nodes {
					if v == exit {
						// We've been here before
						continue Path
					}
				}
			}
			clone := make([]string, len(path.Nodes)+1)
			copy(clone, path.Nodes)
			clone[len(path.Nodes)] = exit

			queue = append(queue, Path{Nodes: clone})
		}
	}
	fmt.Println(count)
}

func addArc(nodes map[string]Node, a, b string) {
	v := nodes[a]
	v.Exits = append(v.Exits, b)
	nodes[a] = v
}
