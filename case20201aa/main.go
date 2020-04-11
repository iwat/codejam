package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	//f := os.Stdin
	f, _ := os.Open("input.txt")

	var t int
	fmt.Fscanf(f, "%d", &t)
	for tid := 0; tid < t; tid++ {
		var n int
		fmt.Fscanf(f, "%d", &n)
		p := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscanln(f, &p[i])
		}
		y := solve(p)
		fmt.Printf("Case #%d: %s\n", tid+1, y)
	}
}

func solve(p []string) string {
	var commonHeads []string
	var commonTails []string
	others := make(map[string]bool)
	for _, pi := range p {
		parts := strings.Split(pi, "*")
		if parts[0] != "" {
			commonHeads = append(commonHeads, parts[0])
			parts = parts[1:]
		}
		if parts[len(parts)-1] != "" {
			commonTails = append(commonTails, parts[len(parts)-1])
			parts = parts[:len(parts)-1]
		}
		others[strings.Join(parts, "")] = true
	}

	head := ""
	if len(commonHeads) > 0 {
		head = commonHeads[0]
		commonHeads = commonHeads[1:]
		for _, h := range commonHeads {
			if len(h) == len(head) {
				if h != head {
					return "*"
				}
			} else if len(h) > len(head) {
				if head != h[:len(head)] {
					return "*"
				}
				head = h
			} else if len(h) < len(head) {
				if h != head[:len(h)] {
					return "*"
				}
			}
		}
	}

	tail := ""
	if len(commonTails) > 0 {
		tail = commonTails[0]
		commonTails = commonTails[1:]
		for _, t := range commonTails {
			if len(t) == len(tail) {
				if t != tail {
					return "*"
				}
			} else if len(t) > len(tail) {
				if tail != t[len(t)-len(tail):] {
					return "*"
				}
				tail = t
			} else if len(t) < len(tail) {
				if t != tail[len(tail)-len(t):] {
					return "*"
				}
			}
		}
	}
	sb := bytes.Buffer{}
	sb.WriteString(head)
	for k := range others {
		sb.WriteString(k)
	}
	sb.WriteString(tail)
	return sb.String()
}
