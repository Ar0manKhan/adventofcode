package main

import (
	"fmt"
	"math"
	"slices"
)

func sortedLrInput01() ([]int, []int) {
	left, right := []int{}, []int{}
	for _, v := range extractInputByLine("input.txt") {
		nums := splitStringToInt(v)
		if len(nums) != 2 {
			panic("Some string does not have length of 2")
		}
		left = append(left, nums[0])
		right = append(right, nums[1])
	}

	slices.Sort(left)
	slices.Sort(right)
	return left, right
}

func solve01v1() {
	result := 0
	left, right := sortedLrInput01()
	for i, l := range left {
		r := right[i]
		result += int(math.Abs(float64(r - l)))
	}
	fmt.Println("Result for 01v1:", result)
}

func solve01v2() {
	result := 0
	left, right := sortedLrInput01()

	j := 0

	for _, v := range left {
		// case 1: current number in right is smaller than current number in left
		for ; j < len(right) && right[j] < v; j++ {
		}
		prevJ := j
		// case 2: we found the current number
		for ; j < len(right) && right[j] == v; j++ {
			result += v
		}
		j = prevJ
		// case 3: we did not found the current number
	}

	fmt.Println("result:", result)

}
