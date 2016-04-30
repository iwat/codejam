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
		nTopics, _ := strconv.Atoi(scanner.Text())
		topics := make(map[string]bool, nTopics)
		for j := 0; j < nTopics; j++ {
			scanner.Scan()
			topics[scanner.Text()] = true
		}
		doCase(i+1, topics)
	}
}

func doCase(caseno int, inputs map[string]bool) {
	firsts := make(map[string]bool)
	seconds := make(map[string]bool)

	for input, _ := range inputs {
		words := strings.SplitN(input, " ", 2)
		firsts[words[0]] = true
		seconds[words[1]] = true
	}

	combs := make(map[string]bool)

	for first, _ := range firsts {
		for second, _ := range seconds {
			if first == second {
				continue
			}
			comb := first + " " + second
			if _, ok := inputs[comb]; !ok {
				if DEBUG {
					fmt.Println("DEBUG", comb)
				}
				combs[comb] = true
			}
		}
	}
	fmt.Printf("Case #%d: %d\n", caseno, len(combs))
}
