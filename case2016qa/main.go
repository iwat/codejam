package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(input)
	if n == 0 {
		fmt.Printf("Case #%d: INSOMNIA\n", caseno)
		return
	}

	if DEBUG {
		fmt.Printf("DEBUG Case #%d: INPUT %d\n", caseno, n)
	}

	matched := 0
	curr := uint(n)

	for {
		m := curr
		if DEBUG {
			fmt.Printf("DEBUG   Checking %d\n", m)
		}

		for {
			d := m % 10
			m = m / 10
			matched |= 1 << d
			if DEBUG {
				fmt.Printf("DEBUG     Digit %d Remain %d Matched %b\n", d, m, matched)
			}

			if m == 0 {
				break
			}
		}
		if matched == 0x03ff {
			fmt.Printf("Case #%d: %d\n", caseno, curr)
			return
		}

		curr += uint(n)
	}
}
