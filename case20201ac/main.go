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
		var r, c int
		fmt.Fscanf(f, "%d %d", &r, &c)
		cells := make([]int64, r*c)
		for i := 0; i < len(cells); i++ {
			fmt.Fscanf(f, "%d", &cells[i])
		}
		y := solve(cells, r, c)
		fmt.Printf("Case #%d: %d\n", tid+1, y)
	}
}

func solve(cells []int64, r, c int) int64 {
	interest := int64(0)
	for {
		for _, i := range cells {
			interest += i
		}
		//for rn := 0; rn < r; rn++ {
		//	fmt.Println(cells[rn*c : (rn+1)*c])
		//}
		//fmt.Println("Interest", interest)
		var eliminations []int
		for rn := 0; rn < r; rn++ {
			for cn := 0; cn < c; cn++ {
				sum := cells[rn*c+cn]
				n := 1
				if sum == 0 {
					continue
				}
				for rrn := rn - 1; rrn >= 0; rrn-- {
					if cells[rrn*c+cn] > 0 {
						sum += cells[rrn*c+cn]
						n++
						break
					}
				}
				for rrn := rn + 1; rrn < r; rrn++ {
					if cells[rrn*c+cn] > 0 {
						sum += cells[rrn*c+cn]
						n++
						break
					}
				}
				for rcn := cn - 1; rcn >= 0; rcn-- {
					if cells[rn*c+rcn] > 0 {
						sum += cells[rn*c+rcn]
						n++
						break
					}
				}
				for rcn := cn + 1; rcn < c; rcn++ {
					if cells[rn*c+rcn] > 0 {
						sum += cells[rn*c+rcn]
						n++
						break
					}
				}
				if n > 1 && float64(cells[rn*c+cn]) < float64(sum)/float64(n) {
					eliminations = append(eliminations, rn*c+cn)
				}
			}
		}
		if len(eliminations) == 0 {
			break
		}
		//fmt.Println("Elim", eliminations)
		for _, e := range eliminations {
			cells[e] = 0
		}
	}
	return interest
}
