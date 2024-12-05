package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type puzzleInput struct {
	hand string
	bid  int
}

func Solve07v1() {
	input := extractInput("input.txt")
	puzzle := make([]puzzleInput, len(input))
	for i, v := range input {
		x := strings.SplitN(v, " ", 2)
		bid, _ := strconv.Atoi(x[1])
		puzzle[i] = puzzleInput{x[0], bid}
	}
	slices.SortFunc(puzzle, func(a, b puzzleInput) int {
		return compareCards(convertCardsToNumbers(a.hand), convertCardsToNumbers(b.hand))
	})
	result := 0
	puzzleLen := len(puzzle)
	for i, p := range puzzle {
		result += p.bid * (puzzleLen - i)
	}
	fmt.Println(result)
}

type cardConverted = [5]uint8

func convertCardsToNumbers(input string) cardConverted {
	res := cardConverted{0, 0, 0, 0, 0}
	for i, v := range input {
		switch v {
		case 'T':
			res[i] = 10
		case 'J':
			res[i] = 11
		case 'Q':
			res[i] = 12
		case 'K':
			res[i] = 13
		case 'A':
			res[i] = 14
		default:
			res[i] = uint8(v - '0')
		}
	}
	return res
}

const (
	fiveKind  = 7
	fourKind  = 6
	fullHouse = 5
	threeKind = 4
	twoPair   = 2
	onePair   = 1
	highCard  = 0
)

func getHandType(hand cardConverted) int {
	// this will initialize all elements with zero
	appearance := make([]uint8, 15)
	for _, card := range hand {
		appearance[card]++
	}
	twosCount, threesCount := 0, 0
	for _, a := range appearance {
		switch a {

		case 2:
			twosCount++
		case 3:
			threesCount++
		case 4:
			return fourKind
		case 5:
			return fiveKind
		}
	}
	if threesCount == 1 {
		if twosCount == 1 {
			return fullHouse
		}
		return threeKind
	}
	switch twosCount {
	case 1:
		return onePair
	case 2:
		return twoPair
	}
	return highCard
}

func compareCards(a, b cardConverted) int {
	aType, bType := getHandType(a), getHandType(b)
	if aType > bType {
		return -1
	}
	if bType > aType {
		return 1
	}
	for i := 0; i < 5; i++ {
		if a[i] > b[i] {
			return -1
		}
		if a[i] < b[i] {
			return 1
		}
	}
	return 0
}
