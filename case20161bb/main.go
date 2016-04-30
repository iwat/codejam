package main

import (
	"bufio"
	"bytes"
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

	splitted := strings.SplitN(input, " ", 2)
	c := []byte(splitted[0])
	j := []byte(splitted[1])

	for i := 0; i < len(c); i++ {
		cmp := bytes.Compare(c[0:i], j[0:i])
		if DEBUG {
			fmt.Print(string(c), " vs ", string(j), " : ", string(c[i]), " ", string(j[i]), " : ", cmp, " > ")
		}
		if c[i] == '?' && j[i] == '?' {
			if cmp == 0 {
				c[i] = '0'
				j[i] = '0'
			} else if cmp == -1 {
				c[i] = '9'
				j[i] = '0'
			} else if cmp == 1 {
				c[i] = '0'
				j[i] = '9'
			}
		} else if c[i] != '?' && j[i] == '?' {
			if cmp == 0 {
				j[i] = c[i]
			} else if cmp == -1 {
				j[i] = '0'
			} else if cmp == 1 {
				j[i] = '9'
			}
		} else if c[i] == '?' && j[i] != '?' {
			if cmp == 0 {
				c[i] = j[i]
			} else if cmp == -1 {
				c[i] = '0'
			} else if cmp == 1 {
				c[i] = '9'
			}
		}
		if DEBUG {
			fmt.Println(string(c), string(j))
		}
	}

	fmt.Printf("Case #%d: %v %v\n", caseno, string(c), string(j))
}
