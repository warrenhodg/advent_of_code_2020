package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	fmt.Printf("%#v\n", all)
	earliest, _ := strconv.Atoi(lines[0])
	min := 0
	minBus := 0
	for i, id := range strings.Split(lines[1], ",") {
		if len(id) == 0 {
			continue
		}
		if id == "x" {
			continue
		}

		bus, _ := strconv.Atoi(id)
		d := earliest / bus
		r := earliest % bus
		if r > 0 {
			d++
		}
		time := bus * d
		if min == 0 || time < min {
			min = time
			minBus = bus
		}
	}
	fmt.Printf("Bus %d leaves at %d, wait %d\n", minBus, min, min-earliest)
}
