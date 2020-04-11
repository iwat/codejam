package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f := os.Stdin

	var t, b int
	fmt.Fscanf(f, "%d %d", &t, &b)
	for tid := 0; tid < t; tid++ {
		if !solve(f, b) {
			break
		}
	}
}

func solve(f io.Reader, b int) bool {
	sent := 0
	sb := bytes.Buffer{}
	for i := 0; i < b; i++ {
		if sent%10 == 0 {
			var s string
			fmt.Println(i + 1)
			fmt.Fscanf(f, "%s", &s)
			sent++
		}
		// send dummy query every 1st, 11st, ...
		var s string
		fmt.Println(i + 1)
		fmt.Fscanf(f, "%s", &s)
		sent++
		sb.WriteString(s)
	}
	log.Println("Solving", sb.String())
	fmt.Println(sb.String())
	var res string
	fmt.Fscanf(f, "%s", &res)
	return res == "Y"
}
