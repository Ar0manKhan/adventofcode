package main

import (
	"fmt"
	"slices"
)

func digitLength11(num uint64) int {
	result := 0
	for t := uint64(1); t <= num; t *= 10 {
		result++
	}
	return result
}

func splitNumbers11(num uint64, l int) (uint64, uint64) {
	target := l >> 1
	t := uint64(1)
	for i := 0; i < target; i++ {
		t *= 10
	}
	return num / t, num % t
}

func Solve11v1() {
	tInput := extractInputByLine("input.txt")
	inputs := splitStringToInt(tInput[0])

	repeats := 25

	q := inputs
	for i := 0; i < repeats; i++ {
		newQ := []int{}
		newQ = slices.Grow(newQ, 2*len(q))
		for _, v := range q {
			numLen := digitLength11(uint64(v))
			if numLen == 0 {
				newQ = append(newQ, 1)
			} else if numLen%2 == 1 {
				newQ = append(newQ, v*2024)
			} else {
				a, b := splitNumbers11(uint64(v), numLen)
				newQ = append(newQ, int(a), int(b))
			}
		}
		// fmt.Println("after iteration:", i, "new q:", newQ)
		q = newQ
	}

	fmt.Println("Result:", len(q))
}

type ProblemSet11 struct {
	num       uint64
	remaining int
}

func PredictStones11(problem ProblemSet11, cache *map[ProblemSet11]uint64) uint64 {
	cached, found := (*cache)[problem]
	if found {
		return cached
	}
	if problem.remaining == 0 {
		return 1
	}
	v := problem.num
	localResult := uint64(0)
	numLen := digitLength11(problem.num)
	if v == 0 {
		localResult += PredictStones11(ProblemSet11{1, problem.remaining - 1}, cache)
	} else if numLen&1 == 1 {
		localResult += PredictStones11(ProblemSet11{problem.num * 2024, problem.remaining - 1}, cache)
	} else {
		a, b := splitNumbers11(problem.num, numLen)
		localResult += PredictStones11(ProblemSet11{a, problem.remaining - 1}, cache)
		localResult += PredictStones11(ProblemSet11{b, problem.remaining - 1}, cache)
	}
	(*cache)[problem] = localResult
	return localResult
}

func Solve11v2() {
	tInput := extractInputByLine("input.txt")
	inputs := splitStringToInt(tInput[0])
	repeats := uint8(75)
	result := uint64(0)

	cache := map[ProblemSet11]uint64{}
	for _, input := range inputs {
		result += PredictStones11(ProblemSet11{uint64(input), int(repeats)}, &cache)
	}

	fmt.Println("Result:", result)
	fmt.Println("Cache len:", len(cache))
}

/*
type ProblemSet11 struct {
	num   uint64
	level uint8
}

// stack underflow will panic
func PopPanic[T any](stack []T, i int) (T, int) {
	return stack[i], i - 1
}

func PushPanic[T any](stack []T, val T, i int) ([]T, int) {
	i++
	stack[i] = val
	return stack, i
}

func Solve11v2() {
	tInput := extractInputByLine("input.txt")
	inputs := splitStringToInt(tInput[0])
	repeats := uint8(40)
	result := uint64(0)
	var wg sync.WaitGroup
	maxVal := uint64(0)

	guessCount := func(stack []ProblemSet11) {
		localResult := uint64(0)
		p := 0
		var v ProblemSet11
		for p > -1 {
			v, p = PopPanic(stack, p)
			maxVal = max(v.num, maxVal)
			if v.level >= uint8(repeats) {
				localResult++
				continue
			}
			numLen := digitLength11(v.num)
			if numLen == 0 {
				stack, p = PushPanic(stack, ProblemSet11{1, v.level + 1}, p)
			} else if numLen&1 == 1 {
				stack, p = PushPanic(stack, ProblemSet11{v.num * 2024, v.level + 1}, p)
			} else {
				a, b := splitNumbers11(v.num, numLen)
				stack, p = PushPanic(stack, ProblemSet11{a, v.level + 1}, p)
				stack, p = PushPanic(stack, ProblemSet11{b, v.level + 1}, p)
			}
		}
		atomic.AddUint64(&result, localResult)
		wg.Done()
	}

	for _, input := range inputs {
		stack := make([]ProblemSet11, repeats+10)
		stack[0] = ProblemSet11{uint64(input), 0}
		wg.Add(1)
		guessCount(stack)
	}

	wg.Wait()

	fmt.Println("Result:", result)
	fmt.Println("max val:", maxVal)
}
*/
