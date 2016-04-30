package main

import (
	"bufio"
	"fmt"
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
		nTopics, _ := strconv.Atoi(scanner.Text())
		topics := make(map[string]int, nTopics)
		for j := 0; j < nTopics; j++ {
			scanner.Scan()
			topics[scanner.Text()]++
		}
		doCase(i+1, topics)
	}
}

func doCase(caseno int, inputs map[string]int) {
	firsts := make(map[string]int)
	seconds := make(map[string]int)

	for input, _ := range inputs {
		words := strings.SplitN(input, " ", 2)
		firsts[words[0]]++
		seconds[words[1]]++
	}

	fakes := 0
	for input, _ := range inputs {
		words := strings.SplitN(input, " ", 2)
		if firsts[words[0]] == 1 || seconds[words[1]] == 1 {
			continue
		}
		fake := false
		for first, _ := range firsts {
			comb := first + " " + words[1]
			if comb == input {
				continue
			}
			if _, ok := inputs[comb]; ok {
				if DEBUG {
					fmt.Println("FAKE", input, "by", comb)
				}
				fake = true
				break
			}
		}
		if fake {
			fakes++
		}
		continue

		for second, _ := range seconds {
			comb := words[0] + " " + second
			if comb == input {
				continue
			}
			if _, ok := inputs[comb]; ok {
				if DEBUG {
					fmt.Println("FAKE", input, "by", comb)
				}
				fake = true
				break
			}
		}
		if fake {
			fakes++
		}
	}

	fmt.Printf("Case #%d: %d\n", caseno, fakes)
}
