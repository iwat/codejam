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
		m := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			fmt.Fscanf(f, "%d", &m[i])
		}
		k, r, c := solve(m, n)
		fmt.Printf("Case #%d: %d %d %d\n", tid+1, k, r, c)
	}
}

func solve(m []int, n int) (int, int, int) {
	k := 0
	for i := 0; i < n; i++ {
		k += m[n*i+i]
	}
	r := 0
	c := 0
	for i := 0; i < n; i++ {
		metR := make(map[int]bool)
		for j := 0; j < n; j++ {
			if metR[m[n*i+j]] {
				r++
				break
			} else {
				metR[m[n*i+j]] = true
			}
		}
		metC := make(map[int]bool)
		for j := 0; j < n; j++ {
			if metC[m[n*j+i]] {
				c++
				break
			} else {
				metC[m[n*j+i]] = true
			}
		}
	}
	return k, r, c
}
