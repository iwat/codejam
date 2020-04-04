package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//f := os.Stdin
	f, _ := os.Open("input.txt")

	var t int
	fmt.Fscanf(f, "%d", &t)
	for tid := 0; tid < t; tid++ {
		var s string
		fmt.Fscanln(f, &s)
		y := solve(s)
		fmt.Printf("Case #%d: %s\n", tid+1, y)
	}
}

func solve(s string) string {
	level := uint8(0)
	sb := bytes.Buffer{} // codejam use Go 1.7, strings.Builder{} was not available back then
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		for n > level {
			sb.WriteRune('(')
			level++
		}
		for n < level {
			sb.WriteRune(')')
			level--
		}
		sb.WriteByte(uint8(s[i]))
	}
	for level > 0 {
		sb.WriteRune(')')
		level--
	}
	return sb.String()
}
