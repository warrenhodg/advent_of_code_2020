package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Bus struct {
	id     int64
	offset int64
}

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	ids := strings.Split(lines[1], ",")
	buses := make([]Bus, 0, len(ids))
	for i, id := range strings.Split(lines[1], ",") {
		if len(id) == 0 {
			continue
		}
		if id == "x" {
			continue
		}

		idn, _ := strconv.Atoi(id)
		offset := (idn - i) % idn
		if offset < 0 {
			offset += idn
		}
		buses = append(buses, Bus{int64(idn), int64(offset)})
	}

	sort.Slice(buses, func(i, j int) bool {
		return buses[i].id > buses[j].id
	})

	fmt.Printf("Buses: %#v\n", buses)

	n := 0
	bus := 1
	m := buses[0].id
	id := buses[bus].id
	offset := buses[bus].offset
	t := int64(buses[0].offset)
	fmt.Printf("Moved on to bus %d(id=%d, offset=%d) at t=%d, m=%d\n", bus, id, offset, t, m)
	for {
		n++
		//fmt.Printf("t=%d\n", t)
		if (t % id) != offset {
			t += m
			continue
		}
		m *= buses[bus].id
		bus++
		if bus == len(buses) {
			break
		}
		id = buses[bus].id
		offset = buses[bus].offset
		fmt.Printf("Moved on to bus %d(id=%d, offset=%d) at t=%d, m=%d\n", bus, id, offset, t, m)
	}

	fmt.Printf("Stopped at t=%d, n=%d\n", t, n)
}
