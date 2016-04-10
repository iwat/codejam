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
	s := strings.SplitN(input, " ", 2)
	N, _ := strconv.Atoi(s[0])
	J, _ := strconv.Atoi(s[1])
	if DEBUG {
		fmt.Printf("DEBUG Case #%d: INPUT %d %d\n", caseno, N, J)
	}

	fmt.Printf("Case #%d:\n", caseno)
	state, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", N-2)+"1", 2)
	for {
		divs, ok := valid(state)
		if ok {
			fmt.Printf("%s %s\n", state.Text(2), strings.Join(divs, " "))
			J--

			if J == 0 {
				break
			}
		}
		state.Add(state, BIG_2)
	}
}

func valid(state *big.Int) ([]string, bool) {
	if DEBUG {
		fmt.Printf("DEBUG   Trying %s\n", state)
	}
	divs := make([]string, 9)

	for base := 2; base <= 10; base++ {
		val, _ := big.NewInt(0).SetString(state.Text(2), base)
		if DEBUG {
			fmt.Printf("DEBUG     base %d is %d ... ", base, val)
		}
		div := divisor(val)
		if DEBUG {
			fmt.Printf("%d\n", div)
		}
		if div > 0 {
			divs[base-2] = strconv.Itoa(int(div))
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

var BIG_0 = big.NewInt(0)
var BIG_1_2 = big.NewInt(1)
var BIG_1 = big.NewInt(1)
var BIG_2 = big.NewInt(2)
var BIG_3 = big.NewInt(3)
var BIG_6 = big.NewInt(6)
var BIG_1000 = big.NewInt(1000)

func divisor(val *big.Int) uint64 {
	test := big.NewInt(0)
	if test.Mod(val, BIG_2).Cmp(BIG_0) == 0 {
		return 2
	}
	if test.Mod(val, BIG_3).Cmp(BIG_0) == 0 {
		return 3
	}

	div := big.NewInt(0)
	for k := big.NewInt(6); k.Cmp(BIG_1000) < 0 && k.Cmp(val) < 0; k.Add(k, BIG_6) {
		if test.Mod(val, div.Sub(k, BIG_1)).Cmp(BIG_0) == 0 {
			return div.Uint64()
		}
		if test.Mod(val, div.Add(k, BIG_1)).Cmp(BIG_0) == 0 {
			return div.Uint64()
		}
	}
	return 0
}
