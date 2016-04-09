package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const DEBUG = true

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
	s := strings.SplitN(input, " ", 3)
	K, _ := strconv.Atoi(s[0])
	C, _ := strconv.Atoi(s[1])
	S, _ := strconv.Atoi(s[2])
	if DEBUG {
		fmt.Printf("DEBUG Case #%d: INPUT %d %d %d\n", caseno, K, C, S)
	}

	ccomb := make(chan string)
	go combine(K, C, ccomb)

	for comb := range ccomb {
		if DEBUG {
			fmt.Printf("DEBUG   %s\n", comb)
		}
	}
}

func combine(K, C int, ccomb chan string) {
	state := new(big.Int)
	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(K)), nil)

	for {
		comb := strings.Replace(fmt.Sprintf("%"+strconv.Itoa(K)+"s", state.Text(2)), " ", "0", -1)
		transform(comb, comb, C, ccomb)
		state.Add(state, big.NewInt(1))
		if state.Cmp(max) >= 0 {
			break
		}
	}

	close(ccomb)
}

func transform(orig, state string, C int, ccomb chan string) {
	if C == 1 {
		ccomb <- state
		return
	}

	buf := make([]byte, len(state)*len(orig))
	for i := 0; i < len(state); i++ {
		for b := i * len(orig); b < (i+1)*len(orig); b++ {
			if state[i] == '0' {
				buf[b] = '0'
			} else {
				buf[b] = orig[b%len(orig)]
			}
		}
	}
	transform(orig, string(buf), C-1, ccomb)
}
