package main

import (
	"fmt"
)

func isValidDifference(prevNumber, num int, prevIncreasing bool) bool {
	diff := num - prevNumber
	return (prevIncreasing && diff >= 1 && diff <= 3) || (!prevIncreasing && diff <= -1 && diff >= -3)
}

type ValidateFnType02 func([]int) bool

func validReportCount(validateFn ValidateFnType02) int {
	result := 0
	for _, line := range extractInputByLine("input.txt") {

		report := splitStringToInt(line)
		if validateFn(report) {
			result++
		}
	}
	return result
}
func isReportSafe2v1(in []int) bool {
	prevNumber := in[0]
	prevIncreasing := true
	if in[1] < in[0] {
		prevIncreasing = false
	}

	for _, num := range in[1:] {
		if isValidDifference(prevNumber, num, prevIncreasing) {
			prevNumber = num
		} else {
			return false
		}
	}
	return true
}

func solve02v1() {
	result := validReportCount(isReportSafe2v1)
	fmt.Println("Result:", result)
}

func solve02v2() {
	isReportIncreasing := func(in []int) bool {
		r := 0
		p := in[0]

		for i := 1; i < min(4, len(in)); i++ {
			if in[i] < p {
				r--
			} else {
				r++
			}
			p = in[i]
		}
		return r > 0
	}

	isReportSafe := func(in []int) bool {
		prevNumber := in[0]
		errorOccured := false
		prevIncreasing := isReportIncreasing(in)
		for _, num := range in[1:] {
			if !isValidDifference(prevNumber, num, prevIncreasing) {
				if errorOccured {
					fmt.Println("Unsafe:", in, "increasing:", prevIncreasing)
					return false || isReportSafe2v1(in[1:])
				} else {
					errorOccured = true
					continue
				}
			}
			prevNumber = num
		}
		return true
	}

	result := validReportCount(isReportSafe)
	fmt.Println("Result:", result)

}
