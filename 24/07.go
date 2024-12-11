package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve07v1() {
	andNumbers := []int{}
	for i := 0; i < 16; i++ {
		andNumbers = append(andNumbers, 1<<i)
	}
	result := uint(0)

	possibleCombination := func(input string) {
		splittedInput := strings.Split(input, ": ")
		target, _ := strconv.Atoi(splittedInput[0])
		availableNumbers := splitStringToInt(splittedInput[1])
		// fmt.Println("For:", target, availableNumbers)

		iterations := 1 << (len(availableNumbers) - 1)

		for i := 0; i < iterations; i++ {
			current := availableNumbers[0]
			for j, num := range availableNumbers[1:] {
				if i&andNumbers[j] != 0 {
					current += num
				} else {
					current *= num
				}
			}
			if current == target {
				result += uint(target)
				return
			}
		}
	}

	for _, input := range extractInputByLine("input.txt") {
		possibleCombination(input)
	}

	fmt.Println("Result:", result)
}

func concatNumbers07(a, b int) int {
	t := 1
	for t <= b {
		t *= 10
	}
	return (a * t) + b
}

func Solve07v2() {
	result := uint(0)

	possibleCombination := func(input string) {
		splittedInput := strings.Split(input, ": ")
		target, _ := strconv.Atoi(splittedInput[0])
		availableNumbers := splitStringToInt(splittedInput[1])
		// fmt.Println("for:", input)

		iterations := int(math.Pow(3, float64(len(availableNumbers)-1)))

		for i := 0; i < iterations; i++ {
			j := i
			current := availableNumbers[0]
			for _, num := range availableNumbers[1:] {
				switch j % 3 {
				case 0:
					current *= num
				case 1:
					current += num
				case 2:
					current = concatNumbers07(current, num)
				}
				j /= 3
			}
			if current == target {
				// fmt.Println("worked")
				result += uint(target)
				return
			}
		}
	}

	for _, input := range extractInputByLine("input.txt") {
		possibleCombination(input)
	}

	fmt.Println("Result:", result)
}
