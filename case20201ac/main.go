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
		for rn := 0; rn < r; rn++ {
			fmt.Println(interest, cells[rn*c:(rn+1)*c])
		}
		var eliminations []int
		for rn := 0; rn < r; rn++ {
			rowSum := int64(0)
			rowN := int64(0)
			for cn := 0; cn < c; cn++ {
				if cells[rn*c+cn] > 0 {
					rowSum += cells[rn*c+cn]
					rowN++
				}
			}
			avg := float64(rowSum) / float64(rowN)
			for cn := 0; cn < c; cn++ {
				if cells[rn*c+cn] > 0 && float64(cells[rn*c+cn]) < avg {
					eliminations = append(eliminations, rn*c+cn)
				}
			}
		}
		for cn := 0; cn < c; cn++ {
			colSum := int64(0)
			colN := int64(0)
			for rn := 0; rn < r; rn++ {
				if cells[rn*c+cn] > 0 {
					colSum += cells[rn*c+cn]
					colN++
				}
			}
			avg := float64(colSum) / float64(colN)
			for rn := 0; rn < r; rn++ {
				if cells[rn*c+cn] > 0 && float64(cells[rn*c+cn]) < avg {
					eliminations = append(eliminations, rn*c+cn)
				}
			}
		}
		if len(eliminations) == 0 {
			break
		}
		for _, e := range eliminations {
			cells[e] = 0
		}
	}
	return interest
}
