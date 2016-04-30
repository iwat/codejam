package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {
	fin, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	ncases, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < ncases; i++ {
		scanner.Scan()
		doCase(i+1, scanner.Text())
	}
}

func doCase(caseno int, input string) {
	if DEBUG {
		fmt.Printf("Case #%d: %s\n", caseno, input)
	}

	splitted := strings.SplitN(input, " ", 2)
	c := []byte(splitted[0])
	j := []byte(splitted[1])

	for i := 0; i < len(c); i++ {
		bestN := -1
		bestM := -1
		bestDiff := new(big.Int)
		bestDiff.SetString("999999999999999999", 10)
		for n := 0; n <= 9; n++ {
			for m := 0; m <= 9; m++ {
				bigC := new(big.Int)
				bigJ := new(big.Int)
				bigC.SetString(strings.Replace(string(c), "?", strconv.Itoa(n), -1), 10)
				bigJ.SetString(strings.Replace(string(j), "?", strconv.Itoa(m), -1), 10)

				newDiff := new(big.Int)
				newDiff.Abs(newDiff.Sub(bigC, bigJ))
				if newDiff.Cmp(bestDiff) < 0 {
					bestDiff = newDiff
					bestN = n
					bestM = m
				}
			}
		}

		if c[i] == '?' {
			c[i] = '0' + byte(bestN)
		}
		if j[i] == '?' {
			j[i] = '0' + byte(bestM)
		}
	}

	fmt.Printf("Case #%d: %v %v\n", caseno, string(c), string(j))
}
