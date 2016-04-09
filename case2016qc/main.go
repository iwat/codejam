package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
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

	cin := make(chan string)
	cout := make(chan string)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(cin chan string, cout chan string) {
			for in := range cin {
				divs, ok := valid(in)
				if ok {
					cout <- fmt.Sprintf("%s %s", in, strings.Join(divs, " "))
				}
			}

		}(cin, cout)
	}

	go func() {
		for {
			cin <- string(state)
			next(state)
		}
	}()

	for j := 0; j < J; j++ {
		fmt.Println(<-cout)
	}
}

func valid(state string) ([]string, bool) {
	if DEBUG {
		fmt.Printf("DEBUG   Trying %s\n", state)
	}
	divs := make([]string, 9)

	for base := 2; base <= 10; base++ {
		val, _ := strconv.ParseInt(state, base, 0)
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
	if val%3 == 0 {
		return 3
	}

	for k := int64(1); k <= int64(math.Sqrt(float64(val))); k++ {
		check := 6*k - 1
		if val%check == 0 {
			return int(check)
		}
		check = 6*k + 1
		if val%check == 0 {
			return int(check)
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
