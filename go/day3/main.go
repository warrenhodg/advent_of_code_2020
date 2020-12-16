package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	x := 0
	d := 1
	trees := 0
	for row, lineStr := range lines {
		if row == 0 || row%2 == 1 {
			continue
		}
		if len(lineStr) == 0 {
			continue
		}
		line := []rune(lineStr)
		x += d
		x = x % len(line)
		tree := false
		if line[x] == '#' {
			tree = true
			trees += 1
		}
		fmt.Printf("row=%d x=%d %v (%s)\n", row, x, tree, lineStr)
	}
	fmt.Printf("Trees = %d\n", trees)
}

//62*184*80*74*36
