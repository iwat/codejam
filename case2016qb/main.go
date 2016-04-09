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
	if DEBUG {
		fmt.Printf("DEBUG Case #%d: INPUT %s\n", caseno, input)
	}

	state := []byte(input)
	nflip := 0

	for i := len(state) - 1; i >= 0; i-- {
		if state[i] == '-' {
			flip(state, i)
			nflip++
			if DEBUG {
				fmt.Printf("DEBUG   Flip %s\n", string(state))
			}
		}
	}

	fmt.Printf("Case #%d: %d\n", caseno, nflip)
}

func flip(state []byte, cur int) {
	for i := 0; i <= cur; i++ {
		if state[i] == '+' {
			state[i] = '-'
		} else {
			state[i] = '+'
		}
	}
}
