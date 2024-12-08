package main

import (
	"fmt"
	"regexp"
	"strings"
)

func calculate3v1(in string) int {
	matchRegex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)$")
	matched := matchRegex.MatchString(in)
	if !matched {
		return 0
	}
	extractRegex := regexp.MustCompile("\\((\\d{1,3},\\d{1,3})\\)$")
	extracts := extractRegex.FindStringSubmatch(in)
	nums := splitStringToIntDelimiter(extracts[1], ",")
	result := 1

	for _, v := range nums {
		result *= v
	}
	return result
}

func solve03v1() {
	result := 0
	for _, val := range getInputByDelimiter(byte(')')) {
		result += calculate3v1(val)
	}

	fmt.Println("Result:", result)
}

func solve03v2() {
	result := 0
	do := true

	for _, val := range getInputByDelimiter(byte(')')) {

		if strings.HasSuffix(val, "do()") {
			do = true
		} else if strings.HasSuffix(val, "don't()") {
			do = false
		}

		if do {
			result += calculate3v1(val)
		}
	}

	fmt.Println("Result:", result)
}
