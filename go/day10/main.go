package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	lines = lines[:len(lines)-1]
	nums := make([]int, len(lines), len(lines)+1)

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		nums[i] = num
	}

	sort.Ints(nums)

	nums = append(nums, nums[len(nums)-1]+3)

	diffs := map[int]int{}
	old := 0
	for _, v := range nums {
		diff := v - old
		diffs[diff]++
		old = v
	}

	counts := make([]int, len(nums)+1)
	for i := range counts {
		counts[i] = -1
	}

	c := count(nums, counts, -1)
	fmt.Printf("%d\n", c)
}

func count(nums []int, counts []int, i int) int {
	sum := 0
	num := 0
	if i >= 0 {
		num = nums[i]
	}

	if counts[i+1] >= 0 {
		return counts[i+1]
	}

	if i == len(nums)-1 {
		return 1
	}

	for j := i + 1; j < len(nums); j++ {
		if nums[j]-num > 3 {
			break
		}

		c := count(nums, counts, j)
		sum += c
	}
	counts[i+1] = sum
	return sum
}
