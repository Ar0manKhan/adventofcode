package main

import (
	"fmt"
	"strconv"
)

func Solve03v1() {
	// loop through each line of input
	// loop through each character in line
	// saving the current number
	// create three variables - currentNumber, startCurrentNumber, endCurrentNumber
	// initialize it with zero
	// if current character is a number
	// do this currentNumber += currentNumber * 10 + currentCharValue
	// if startCurrentNumber is -1, set it to current index
	// set endCurrentNumber to current index
	// else if currentNumber is not 0
	// send current numbers position to check if symbol exists around them
	// if it does, add currentNumber to answer
	// and also set currentNumber to 0, startCurrentNumber & endCurrentNumber to -1
	// end if
	// end character loop
	// if currentNumber is not 0
	// send current numbers position to check if symbol exists around them
	// if it does, add currentNumber to answer
	// and also set currentNumber to 0, startCurrentNumber & endCurrentNumber to -1
	// end if
	// end line loop
	input := extractInput("input.txt")
	answer, currentNumber, cnStartIdx, cnEndIdx := 0, 0, -1, -1
	for x, line := range input {
		for y, c := range line {
			if '0' <= c && c <= '9' {
				currentNumber = currentNumber*10 + int(c-'0')
				if cnStartIdx == -1 {
					cnStartIdx = y
				}
				cnEndIdx = y
			} else if currentNumber > 0 {
				if hasSymbolAround(input, x, cnStartIdx, cnEndIdx) {
					answer += currentNumber
				}
				currentNumber, cnStartIdx, cnEndIdx = 0, -1, -1
			}
		}
		if currentNumber > 0 {
			if hasSymbolAround(input, x, cnStartIdx, cnEndIdx) {
				answer += currentNumber
			}
			currentNumber, cnStartIdx, cnEndIdx = 0, -1, -1
		}
	}
	fmt.Println(answer)
}

func hasSymbolAround(input []string, lineNo, start, end int) bool {
	hlen := len(input[lineNo])
	if lineNo > 0 {
		for i := max(0, start-1); i <= min(end+1, hlen-1); i++ {
			if isSymbol(input[lineNo-1][i]) {
				return true
			}
		}
	}
	if (start > 0 && isSymbol(input[lineNo][start-1])) || (end < hlen-1 && isSymbol(input[lineNo][end+1])) {
		return true
	}
	if lineNo < len(input)-1 {
		for i := max(0, start-1); i <= min(end+1, hlen-1); i++ {
			if isSymbol(input[lineNo+1][i]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}

func Solve03v2() {
	input := extractInput("input.txt")
	answer := 0
	for x := range input {
		for y, c := range input[x] {
			if c != '*' {
				continue
			}
			answer += productOfNumbersAround(input, x, y)
		}
	}
	println(answer)
}

func productOfNumbersAround(input []string, row, col int) int {
	numsFound := 0
	result := 1
	isOnLeftEdge := col == 0
	isOnRightEdge := col == len(input[0])-1
	if row > 0 {
		if isByteDigit(input[row-1][col]) {
			result *= getNumberSpanningHorizontally(input[row-1], col)
			numsFound++
		} else {
			if !isOnLeftEdge && isByteDigit(input[row-1][col-1]) {
				result *= getNumberSpanningHorizontally(input[row-1], col-1)
				numsFound++
			}
			if !isOnRightEdge && isByteDigit(input[row-1][col+1]) {
				result *= getNumberSpanningHorizontally(input[row-1], col+1)
				numsFound++
			}
		}
	}
	if !isOnLeftEdge && isByteDigit(input[row][col-1]) {
		result *= getNumberSpanningHorizontally(input[row], col-1)
		numsFound++
	}
	if !isOnRightEdge && isByteDigit(input[row][col+1]) {
		result *= getNumberSpanningHorizontally(input[row], col+1)
		numsFound++
	}
	if row < len(input)+1 {
		if isByteDigit(input[row+1][col]) {
			result *= getNumberSpanningHorizontally(input[row+1], col)
			numsFound++
		} else {
			if !isOnLeftEdge && isByteDigit(input[row+1][col-1]) {
				result *= getNumberSpanningHorizontally(input[row+1], col-1)
				numsFound++
			}
			if !isOnRightEdge && isByteDigit(input[row+1][col+1]) {
				result *= getNumberSpanningHorizontally(input[row+1], col+1)
				numsFound++
			}
		}
	}
	if numsFound == 2 {
		return result
	}
	return 0
}

func getNumberSpanningHorizontally(input string, x int) int {
	var start, end int

	for start = x; start > -1 && isByteDigit(input[start]); start-- {
	}
	for end = x; end < len(input) && isByteDigit(input[end]); end++ {
	}
	val, err := strconv.Atoi(input[start+1 : end])
	if err != nil {
		panic(err)
	}
	return val
}

func isByteDigit(c byte) bool {
	return c > 47 && c < 58
}
