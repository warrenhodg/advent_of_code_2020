package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	hgtRegex = regexp.MustCompile("([0-9]+)(cm|in)")
	hclRegex = regexp.MustCompile("#[0-9a-f]{6}?")
)

func main() {
	all, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(all), "\n")
	completeLine := ""
	validCount := 0
	for row, line := range lines {
		if len(line) != 0 {
			completeLine = fmt.Sprintf("%s %s", completeLine, line)
			continue
		}

		fields := map[string]string{}
		fmt.Printf("%d: %s\n", row, completeLine)
		for _, token := range strings.Split(completeLine, " ") {
			if len(token) == 0 {
				continue
			}

			colon := strings.Index(token, ":")
			field := token[:colon]
			value := token[colon+1:]
			fields[field] = value
		}

		isValid := false
		if valid(fields) {
			validCount += 1
			isValid = true
		}

		fmt.Printf("\t%s : %v\n", completeLine, isValid)

		completeLine = ""
	}

	fmt.Printf("Valid = %d\n", validCount)
}

func valid(m map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, key := range keys {
		if _, found := m[key]; !found {
			fmt.Printf("\t%s missing\n", key)
			return false
		}
	}

	if !validDate("byr", m, 1920, 2002, 4) {
		return false
	}
	if !validDate("iyr", m, 2010, 2020, 4) {
		return false
	}
	if !validDate("eyr", m, 2020, 2030, 4) {
		return false
	}
	if !validDate("eyr", m, 2020, 2030, 4) {
		return false
	}
	hgtParts := hgtRegex.FindStringSubmatch(m["hgt"])
	if hgtParts == nil {
		fmt.Printf("\thgt syntax error\n")
		return false
	}
	if hgtParts[2] == "cm" {
		h, _ := strconv.Atoi(hgtParts[1])
		if h < 150 || h > 193 {
			fmt.Printf("\thgt cm out of range\n")
			return false
		}
	} else if hgtParts[2] == "in" {
		h, _ := strconv.Atoi(hgtParts[1])
		if h < 59 || h > 76 {
			fmt.Printf("\thgt in out of range\n")
			return false
		}
	}
	if !hclRegex.MatchString(m["hcl"]) {
		fmt.Printf("\tHair colour mismatch\n")
		return false
	}
	ecl := m["ecl"]
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		fmt.Printf("\tEye colour wrong\n")
		return false
	}
	if !validDate("pid", m, 0, 999999999, 9) {
		return false
	}

	return true
}

func validDate(fieldName string, m map[string]string, min int, max int, digits int) bool {
	field := m[fieldName]

	if len(field) != digits {
		fmt.Printf("\t%s != %d\n", field, digits)
		return false
	}

	n, err := strconv.Atoi(field)
	if err != nil {
		fmt.Printf("\t%s nan\n", field)
		return false
	}

	if n < min || n > max {
		fmt.Printf("\t%s(%d) out of range\n", fieldName, n)
		return false
	}

	return true
}
