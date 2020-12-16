package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	empty = byte(0)
	full  = byte(1)
	floor = byte(2)
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	lines = lines[:len(lines)-1]
	h := len(lines)
	w := len(lines[0])
	grid := make([]byte, h*w)
	buffer := make([]byte, h*w)
	for y, line := range lines {
		for x, c := range line {
			v := floor
			if c == 'L' {
				v = empty
			} else if c == '#' {
				v = full
			}

			grid[y*w+x] = v
			buffer[y*w+x] = v
		}
	}

	for {
		fmt.Printf(".")
		changed := iterate(grid, buffer, w, h)
		buffer, grid = grid, buffer
		if !changed {
			break
		}
	}

	fmt.Printf("\n%d\n", count(grid))
}

func count(grid []byte) int {
	sum := 0
	for _, v := range grid {
		if v == full {
			sum++
		}
	}
	return sum
}

func show(grid []byte, w, h int) {
	for yy := 0; yy < h; yy++ {
		for xx := 0; xx < w; xx++ {
			switch grid[yy*w+xx] {
			case empty:
				fmt.Printf("L")
			case full:
				fmt.Printf("#")
			case floor:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func iterate(grid []byte, buffer []byte, w, h int) bool {
	changed := false
	for yy := 0; yy < h; yy++ {
		for xx := 0; xx < w; xx++ {
			v := grid[yy*w+xx]
			if v == floor {
				continue
			}
			s := surround(grid, w, h, xx, yy)
			switch {
			case v == empty && s == 0:
				buffer[yy*w+xx] = full
				changed = true
				//fmt.Printf("\ty=%d, x=%d, s=%d sit\n", yy, xx, s)

			case v == full && s >= 4:
				buffer[yy*w+xx] = empty
				changed = true
				//fmt.Printf("\ty=%d, x=%d, s=%d leave\n", yy, xx, s)

			default:
				buffer[yy*w+xx] = v
			}
		}
	}
	return changed
}

func surround(grid []byte, w int, h int, x int, y int) int {
	sum := 0
	for yy := y - 1; yy <= y+1; yy++ {
		if yy < 0 {
			continue
		}
		if yy >= h {
			break
		}
		for xx := x - 1; xx <= x+1; xx++ {
			if xx < 0 {
				continue
			}
			if xx >= w {
				break
			}
			if xx == x && yy == y {
				continue
			}

			if grid[yy*w+xx] == full {
				sum++
			}
		}
	}

	return sum
}
