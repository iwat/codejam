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
	result := input[0:1]

	for i := 1; i < len(input); i++ {
		if input[i] >= result[0] {
			result = input[i:i+1] + result
		} else {
			result = result + input[i:i+1]
		}
	}

	fmt.Printf("Case #%d: %s\n", caseno, result)
}
