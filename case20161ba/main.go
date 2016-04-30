package main

import (
	"bufio"
	"fmt"
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
	numbers := new([10]int)

	input = process(input, "Z", "ZERO", 0, numbers)
	input = process(input, "W", "TWO", 2, numbers)
	input = process(input, "U", "FOUR", 4, numbers)
	input = process(input, "X", "SIX", 6, numbers)
	input = process(input, "G", "EIGHT", 8, numbers)
	input = process(input, "O", "ONE", 1, numbers)
	input = process(input, "R", "THREE", 3, numbers)
	input = process(input, "F", "FIVE", 5, numbers)
	input = process(input, "V", "SEVEN", 7, numbers)
	input = process(input, "N", "NINE", 9, numbers)

	if len(input) != 0 {
		panic(input)
	}

	result := ""
	for i := 0; i < 10; i++ {
		result = result + strings.Repeat(strconv.Itoa(i), numbers[i])
	}

	fmt.Printf("Case #%d: %v\n", caseno, result)
}

func process(input, mark, full string, ndx int, numbers *[10]int) string {
	numbers[ndx] = strings.Count(input, mark)
	if DEBUG {
		fmt.Printf("Found %s %d times\n", full, numbers[ndx])
	}
	for i := 0; i < len(full); i++ {
		input = strings.Replace(input, string(full[i]), "", numbers[ndx])
	}

	return input
}
