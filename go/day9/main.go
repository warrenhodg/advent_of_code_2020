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
	lines = lines[:len(lines)-1]
	nums := make([]int, len(lines))

	window := 25
	numMap := make(map[int]bool)
	badNum := 0

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		nums[i] = num
		numMap[num] = true
		if i < window {
			continue
		}

		start := i - window
		if start < 0 {
			start = 0
		}

		found := false
		for j := start; j < i-1; j++ {
			num1 := nums[j]
			if numMap[num-num1] {
				found = true
				break
			}
		}

		if !found {
			badNum = num
			break
		}

		if start >= 0 {
			oldNum := nums[start]
			delete(numMap, oldNum)
		}
	}

	fmt.Printf("Broke at %d\n", badNum)
	start := 0
	stop := 0
	sum := 0
	for i := range nums {
		if sum == badNum {
			stop = i
			break
		}

		sum += nums[i]
		fmt.Printf("\t%d (%d) start=%d, sum=%d\n", i, nums[i], start, sum)

		for sum > badNum {
			oldSum := sum
			oldStart := start
			sum -= nums[start]
			start++
			fmt.Printf("\t\t%d start=%d->%d, sum=%d-%d = %d\n", i, oldStart, start, oldSum, nums[oldStart], sum)
		}
	}

	if sum == 0 {
		fmt.Printf("Did not find sum\n")
		return
	}

	min := 0
	max := 0
	for i := start; i < stop; i++ {
		num := nums[i]
		if num > max {
			max = num
		}
		if num < min || min == 0 {
			min = num
		}
	}

	fmt.Printf("Sum is from %d to %d = %d (min=%d, max=%d, sum=%d)\n", start, stop, sum, min, max, min+max)
	fmt.Printf("%#v\n", nums[start:stop])
}
