package main

import (
	"fmt"
	"math"
	"strings"
)

func Solve04v1() {
	input := extractInput("input.txt")
	// splitter is same for each line
	const INITIAL = 9
	splitter := strings.Index(input[0], "|")
	answer := 0
	for _, line := range input {
		numbersMatched := getNumbersMatchedForLine(line, INITIAL, splitter)
		if numbersMatched > 0 {
			answer += int(math.Pow(2, float64(numbersMatched-1)))
		}
	}
	println(answer)
}

func Solve04v2() {
	input := extractInput("input.txt")
	const INITIAL = 9
	splitter := strings.Index(input[0], "|")
	answer := 0
	// taken sliding frame approach
	copy_frame := make([]int, 10)
	for x, line := range input {
		count := 1 + copy_frame[x%10]
		copy_frame[x%10] = 0
		answer += count
		numberMatched := getNumbersMatchedForLine(line, INITIAL, splitter)
		for i := x + 1; i <= x+numberMatched; i++ {
			copy_frame[i%10] += count
		}
	}
	fmt.Println(answer)
}

func getNumbersMatchedForLine(line string, INITIAL, SPLITTER int) int {
	// in attempt to write efficient code, I wrote this mess
	// better version would be to just use regex and split with \s+ and
	// check if first element in both arrays is empty string then slice it
	splitter := strings.Index(line, "|")
	winning_numbers := getWinningNumbers(line[INITIAL : splitter-1])
	got_numbers := getGotNumbers(line[splitter+2:])
	return getIntersectionCount(got_numbers, winning_numbers)
}

func getIntersectionCount(a []string, b []string) int {
	count := 0
	// it is O(n^2) algorithm for getting intersection count but it
	// is efficient compared to other methods for small size
	for _, i := range a {
		for _, j := range b {
			if i == j {
				count++
			}
		}
	}
	return count
}

func getWinningNumbers(input string) []string {
	const WINNING_NUMBER_COUNT = 10
	winning_numbers := make([]string, WINNING_NUMBER_COUNT)
	i := 0
	for _, num := range strings.Split(input, " ") {
		if num != "" {
			winning_numbers[i] = num
			i++
		}
	}
	return winning_numbers
}

func getGotNumbers(input string) []string {
	const GOT_NUMBER_COUNT = 25
	got_numbers := make([]string, GOT_NUMBER_COUNT)
	i := 0
	for _, num := range strings.Split(input, " ") {
		if num != "" {
			got_numbers[i] = num
			i++
		}
	}
	return got_numbers
}
