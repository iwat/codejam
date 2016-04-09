package main

import (
	"bufio"
	"fmt"
	"math"
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
	s := strings.SplitN(input, " ", 2)
	N, _ := strconv.Atoi(s[0])
	J, _ := strconv.Atoi(s[1])
	if DEBUG {
		fmt.Printf("DEBUG Case #%d: INPUT %d %d\n", caseno, N, J)
	}

	state := make([]byte, N)
	for i := 1; i < len(state)-1; i++ {
		state[i] = '0'
	}
	state[0] = '1'
	state[len(state)-1] = '1'

	fmt.Printf("Case #%d:\n", caseno)

	for j := 0; j < J; j++ {
		for {
			if divs, ok := valid(state); ok {
				fmt.Printf("%s %s\n", string(state), strings.Join(divs, " "))
				next(state)
				break
			}
			next(state)
		}
	}
}

func valid(state []byte) ([]string, bool) {
	if DEBUG {
		fmt.Printf("DEBUG   Trying %s\n", string(state))
	}
	divs := make([]string, 9)

	for base := 2; base <= 10; base++ {
		val, _ := strconv.ParseInt(string(state), base, 0)
		if DEBUG {
			fmt.Printf("DEBUG     base %d is %d ... ", base, val)
		}
		div := divisor(val)
		if DEBUG {
			fmt.Printf("%d\n", div)
		}
		if div > 0 {
			divs[base-2] = strconv.Itoa(div)
		} else {
			return divs, false
		}
	}

	for _, div := range divs {
		if div == "" {
			return divs, false
		}
	}

	return divs, true
}

func divisor(val int64) int {
	if val%2 == 0 {
		return 2
	}

	for i := int64(3); i <= int64(math.Sqrt(float64(val))); i += 2 {
		if val%i == 0 {
			return int(i)
		}
	}
	return 0
}

func next(state []byte) {
	now, _ := strconv.ParseInt(string(state), 2, 0)
	next := now + 2
	newstate := strconv.FormatInt(next, 2)

	if len(newstate) != len(state) {
		panic("new state is too large")
	}

	for i := 0; i < len(newstate); i++ {
		state[i] = newstate[i]
	}
}
