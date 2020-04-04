package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

type event struct {
	ndx      int
	start    int
	end      int
	assignee rune
}

type byNdx []event

func (es byNdx) Len() int {
	return len(es)
}

func (es byNdx) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es byNdx) Less(i, j int) bool {
	return es[i].ndx < es[j].ndx
}

type byStart []event

func (es byStart) Len() int {
	return len(es)
}

func (es byStart) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es byStart) Less(i, j int) bool {
	return es[i].start < es[j].start
}

func main() {
	//f := os.Stdin
	f, _ := os.Open("input.txt")

	var t int
	fmt.Fscanf(f, "%d", &t)
	for tid := 0; tid < t; tid++ {
		var n int
		fmt.Fscanf(f, "%d", &n)
		es := make([]event, n)
		for i := 0; i < len(es); i++ {
			es[i].ndx = i
			fmt.Fscanf(f, "%d", &es[i].start)
			fmt.Fscanf(f, "%d", &es[i].end)
		}
		y := solve(es)
		fmt.Printf("Case #%d: %s\n", tid+1, y)
	}
}

func solve(es []event) string {
	sort.Sort(byStart(es))

	cameron := 0
	jamie := 0
	for i, e := range es {
		if cameron <= e.start {
			cameron = e.end
			es[i].assignee = 'C'
		} else if jamie <= e.start {
			jamie = e.end
			es[i].assignee = 'J'
		} else {
			return "IMPOSSIBLE"
		}
	}

	sort.Sort(byNdx(es))
	sb := bytes.Buffer{}
	for _, e := range es {
		sb.WriteRune(e.assignee)
	}
	return sb.String()
}
