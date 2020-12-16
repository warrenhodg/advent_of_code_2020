package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	name    string
	arg     int
	visited bool
}

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	lines = lines[:len(lines)-1]
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		line = strings.Replace(line, "+", "", 1)
		arg, _ := strconv.Atoi(line[4:])
		instructions[i] = instruction{
			line[:3],
			arg,
			false,
		}
	}

l:
	for i := range instructions {
		was := instructions[i].name

		switch instructions[i].name {
		case "nop":
			instructions[i].name = "jmp"

		case "jmp":
			instructions[i].name = "nop"

		default:
			continue l
		}

		accum, infinite := run(instructions)
		if !infinite {
			fmt.Printf("Finished with accum=%d, by swapping line %d from %s to %s\n", accum, i, instructions[i].name, was)
			return
		}

		instructions[i].name = was
	}
	fmt.Printf("No solution")
}

func show(instructions []instruction) {
	fmt.Printf("-------\n")
	for i := range instructions {
		fmt.Printf("%d) %s %d\n", i, instructions[i].name, instructions[i].arg)
	}
}

func run(instructions []instruction) (accum int, infinite bool) {
	line := 0

	show(instructions)

	//reset
	for i := range instructions {
		instructions[i].visited = false
	}

	for {
		if line == len(instructions) {
			return accum, false
		}

		fmt.Printf("\t%d) %s %d (%d)\n", line, instructions[line].name, instructions[line].arg, accum)

		if instructions[line].visited {
			return accum, true
		}

		instructions[line].visited = true
		instruction := instructions[line]
		switch instruction.name {
		case "nop":
			line++

		case "acc":
			accum += instruction.arg
			line++

		case "jmp":
			line += instruction.arg

		case "": //EOF
			return accum, false
		}
	}
}
