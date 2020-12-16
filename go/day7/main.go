package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")

	r := regexp.MustCompile("^ ?[0-9]+ (.*) bags?$")

	outerBags := map[string][]string{}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		bci := strings.Index(line, " bags contain ")
		outerBag := line[:bci]

		innerBags := make([]string, 0, 5)

		rest := line[bci+len(" bags contain "):]
		rest = strings.Replace(rest, ".", ",", 1)
		fmt.Printf("%d) >%s< >%s<\n", i, outerBag, rest)

		bagCounts := strings.Split(rest, ",")
		for j, bagCount := range bagCounts {
			matches := r.FindStringSubmatch(bagCount)
			if matches == nil {
				fmt.Printf("\t%d) nothing (%s)\n", j, bagCount)
				continue
			}

			innerBags = append(innerBags, matches[1])
			fmt.Printf("\t%d : >%s<\n", j, matches[1])
		}

		outerBags[outerBag] = innerBags
	}

	fmt.Printf("\n\n----------\n\n")

	num := 0
	goldContainers := map[string]bool{}
	for outerBag := range outerBags {
		fmt.Printf("%s\n", outerBag)
		if contains(outerBags, goldContainers, outerBag, "\t") {
			//fmt.Printf("%s - yes\n", outerBag)
			num++
		} else {
			//fmt.Printf("%s - no\n", outerBag)
		}
	}

	fmt.Printf("%#v\n", num)
}

func contains(outerBags map[string][]string, goldContainers map[string]bool, outerBag string, indent string) bool {
	//fmt.Printf("%s- %s\n", indent, outerBag)
	if goldContainers[outerBag] {
		return true
	}

	found := false
	for _, innerBag := range outerBags[outerBag] {
		if innerBag == "shiny gold" {
			found = true
			goldContainers[outerBag] = true
			continue
		}

		if contains(outerBags, goldContainers, innerBag, fmt.Sprintf("%s\t", indent)) {
			return true
		}
	}

	return found
}
