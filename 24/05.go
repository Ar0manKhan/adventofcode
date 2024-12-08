package main

import (
	"fmt"
	"slices"
)

func validate05v1(rules *map[int][]int, seq []int) bool {
	for i := 1; i < len(seq); i++ {
		val := seq[i]
		limits := (*rules)[val]

		for _, limit := range limits {
			if slices.Contains(seq[:i], limit) {
				return false
			}
		}
	}
	return true
}

func extractRuleSeq5() (*map[int][]int, [][]int) {
	loadingRules := true
	rules := map[int][]int{}
	seq := [][]int{}
	for _, row := range extractInputByLine("input.txt") {
		if len(row) == 0 {
			loadingRules = false
			continue
		}

		if loadingRules {
			nums := splitStringToIntDelimiter(row, "|")
			lower := nums[0]
			upper := nums[1]

			val, found := rules[lower]
			if found {
				val = append(val, upper)
				rules[lower] = val
			} else {
				rules[lower] = []int{upper}
			}
		} else {
			seq = append(seq, splitStringToIntDelimiter(row, ","))
		}
	}

	return &rules, seq
}

func Solve05v1() {
	result := 0
	rules, seqs := extractRuleSeq5()
	for _, seq := range seqs {
		if validate05v1(rules, seq) {
			result += seq[len(seq)/2]
		}
	}

	fmt.Println("result:", result)
}

func Solve05v2() {
	result := 0
	rules, seqs := extractRuleSeq5()

	for _, seq := range seqs {
		if !validate05v1(rules, seq) {
			slices.SortFunc(seq, func(a, b int) int {
				aRules := (*rules)[a]
				bRules := (*rules)[b]
				if slices.Contains(aRules, b) {
					return -1
				}
				if slices.Contains(bRules, a) {
					return 1
				}
				return 0
			})
			result += seq[len(seq)/2]
		}
		// if !validate05v1(rules, seq) {
		// 	fmt.Println("something wrong with seq", seq)
		// }
	}

	fmt.Println("result:", result)
}
