package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	yesses := map[string]int{}
	size := 0
	sum := 0
	for l, line := range lines {
		if len(line) == 0 {
			ssum := 0
			for _, y := range yesses {
				if y == size {
					ssum += 1
				}
			}
			sum += ssum
			fmt.Printf("%d (%d)) %v +%d\n", l, size, yesses, ssum)

			size = 0
			yesses = map[string]int{}
			continue
		}

		size += 1
		for _, c := range line {
			yesses[string(c)] += 1
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}
