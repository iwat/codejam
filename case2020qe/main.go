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
		var n, k int
		fmt.Fscanf(f, "%d %d", &n, &k)
		m := solve(n, k)
		fmt.Printf("Case #%d: %s\n", tid+1, m)
	}
}

func solve(n, k int) string {
	m := make([]int, n*n)
	// fill trace maker
	for i := 0; i < n; i++ {
		t := k - (n - i - 1)
		if t > n {
			t = n
		}
		m[n*i+i] = t
		k -= t
	}
	// fill the remaining
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[n*i+j] == 0 {
				metR := make(map[int]bool, n+1)
				for k := 0; k < n; k++ {
					metR[m[n*i+k]] = true
				}
				metC := make(map[int]bool, n)
				for k := 0; k < n; k++ {
					metC[m[n*k+j]] = true
				}
				var valid int
				for l := 1; l <= n; l++ {
					if metR[l] == false && metC[l] == false {
						valid = l
						break
					}
				}
				if valid > 0 {
					m[n*i+j] = valid
				}
			}
		}
	}

	// render output
	sb := bytes.Buffer{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				sb.WriteRune(' ')
			}
			sb.WriteRune(rune(m[n*i+j] + '0'))
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')

	return "POSSIBLE\n" + sb.String()
}
