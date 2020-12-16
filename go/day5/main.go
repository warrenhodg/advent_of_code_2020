package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	max := 0
	m := map[int]bool{}
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		fmt.Printf("%s\n", line)
		row := n(line[:7], 0, 128, 'F', 'B')
		col := n(line[7:], 0, 8, 'L', 'R')
		id := row*8 + col
		fmt.Printf("%d) %s - %d*8+%d=%d\n", i, line, row, col, id)
		if id > max {
			max = id
			fmt.Printf("\tNew max: %d\n", max)
		}
		m[id] = true
	}
	fmt.Printf("Max is %d\n", max)

	for i := 0; i < max; i++ {
		if m[i-1] && m[i+1] && !m[i] {
			fmt.Printf("Your seat is %d\n", i)
		}
	}
}

func n(s string, min int, max int, minc rune, maxc rune) int {
	fmt.Printf("\t\t%s (%d - %d)\n", s, min, max)
	if len(s) == 0 {
		return min
	}
	if rune(s[0]) == minc {
		return n(s[1:], min, (min+max)/2, minc, maxc)
	}
	return n(s[1:], (min+max)/2, max, minc, maxc)
}
