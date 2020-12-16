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

		first -= 1
		second -= 1

		c := line[space+1]
		str := line[space+4:]

		fc := str[first]
		sc := str[second]

		isValid := false
		if (fc == c && sc != c) || (fc != c && sc == c) {
			isValid = true
			valid++
		}

		fmt.Printf("%d: %d - %d : %s (%s)[%s, %s] %v\n", i, first, second, string(c), str, string(fc), string(sc), isValid)
	}
	fmt.Printf("Valid=%d\n", valid)
}
