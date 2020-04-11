package main

import (
	"fmt"
	"os"
)

func main() {
	//f := os.Stdin
	f, _ := os.Open("input.txt")

	var t int
	fmt.Fscanf(f, "%d", &t)
	for tid := 0; tid < t; tid++ {
		var n int
		fmt.Fscanf(f, "%d", &n)
		y := solve(n)
		fmt.Printf("Case #%d:\n", tid+1)
		for _, rk := range y {
			fmt.Println(rk)
		}
	}
}

type rk struct {
	r, k int
}

func solve(n int) []string {
	mem := make(map[rk]int)
	value := pascal(5, 3, mem)
	fmt.Println(value)
	return []string{"1 1", "2 2"}
}

func pascal(r, k int, mem map[rk]int) int {
	if v, ok := mem[rk{r, k}]; ok {
		return v
	}
	if r == 1 || k == 1 || r == k {
		return 1
	}
	return pascal(r-1, k-1, mem) + pascal(r-1, k, mem)
}
