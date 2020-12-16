package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	_ "strconv"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	valid := 0
	for i, lineStr := range lines {
		line := lineStr
		if len(line) == 0 {
			continue
		}

		dash := strings.IndexRune(line, '-')
		firstStr := line[:dash]
		space := strings.IndexRune(line, ' ')
		secondStr := lineStr[dash+1 : space]

		first, _ := strconv.Atoi(firstStr)
		second, _ := strconv.Atoi(secondStr)

		c := rune(line[space+1])
		str := line[space+4:]

		count := 0
		for _, v := range []rune(str) {
			if v == c {
				count++
			}
			if count > second {
				break
			}
		}
		isValid := false
		if count >= first && count <= second {
			isValid = true
			valid++
		}

		fmt.Printf("%d: %d - %d : %s-%s %v\n", i, first, second, string(c), str, isValid)
	}
	fmt.Printf("Valid=%d\n", valid)
}
