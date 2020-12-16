package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")

	r := regexp.MustCompile("^ ?([0-9]+) (.*) bags?$")

	outerBags := map[string]map[string]int{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		bci := strings.Index(line, " bags contain ")
		outerBag := line[:bci]

		innerBags := map[string]int{}

		rest := line[bci+len(" bags contain "):]
		rest = strings.Replace(rest, ".", ",", 1)

		bagCounts := strings.Split(rest, ",")
		for _, bagCount := range bagCounts {
			matches := r.FindStringSubmatch(bagCount)
			if matches == nil {
				continue
			}

			count, _ := strconv.Atoi(matches[1])
			innerBag := matches[2]
			innerBags[innerBag] = count
		}

		outerBags[outerBag] = innerBags
	}

	fmt.Printf("Total bags : %d\n", contains(outerBags, "shiny gold"))
}

func contains(outerBags map[string]map[string]int, outerBag string) int {
	sum := 0
	innerBags := outerBags[outerBag]
	for innerBag, count := range innerBags {
		sum += count * (1 + contains(outerBags, innerBag))
	}
	return sum
}
