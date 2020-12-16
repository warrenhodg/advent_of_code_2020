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
	m := map[int64]bool{}
	for _, line := range lines {
		vv1, _ := strconv.ParseInt(line, 10, 64)
		v1 := int64(vv1)

		for v2 := range m {
			if v1+v2 > 2020 {
				continue
			}

			v3 := 2020 - v1 - v2
			_, found := m[v3]
			if found {
				fmt.Printf("%d*%d*%d=%d", v1, v2, v3, v1*v2*v3)
				return
			}
		}

		m[v1] = true
	}
	fmt.Printf("%d\n", len(lines))
}
