package main

import (
	"container/heap"
	"fmt"
)

func Solve09v1() {
	inputGrid := extractInputInByteGrid()
	inputBytes := inputGrid[0]
	inputs := splitStringToIntDelimiter(string(inputBytes), "")
	a := -2
	b := len(inputs) + 1
	availableSpace := 0
	needSpace := 0
	idx := 0
	result := 0
	id := 0

	for a < b-2 {
		if availableSpace == 0 {
			a += 2
			id = a >> 1 // id = idx / 2
			for i := 0; i < inputs[a]; i++ {
				result += (id * idx)
				idx++
			}
			if a+1 < len(inputs) {
				availableSpace = inputs[a+1]
			}
		}

		if needSpace == 0 && b > a+2 {
			b -= 2
			if needSpace > -1 {
				needSpace = inputs[b]
			}
		}

		id = b >> 1
		limit := min(needSpace, availableSpace)
		for i := 0; i < limit; i++ {
			result += (id * idx)
			idx++
		}
		availableSpace -= limit
		needSpace -= limit
	}

	id = b >> 1
	if needSpace > 0 {
		for i := 0; i < needSpace; i++ {
			result += (id * idx)
			idx++
		}
	}

	fmt.Println("Result:", result)
}

func Solve09v2() {
	inputGrid := extractInputInByteGrid()
	inputBytes := inputGrid[0]
	inputs := splitStringToIntDelimiter(string(inputBytes), "") // extracting input io []int

	spaceHeaps := make([]MinHeap[int], 10)
	for i := 0; i < 10; i++ {
		h := GetMinHeap[int]()
		spaceHeaps[i] = *h
	}

	// returns space available and index
	getNextEmptyIdx := func(minVal int, maxIdx int) (int, int) {
		space := 10
		possibleIdx := 200000
		for i := 9; i >= minVal; i-- {
			if len(spaceHeaps[i]) == 0 {
				continue
			}
			t := spaceHeaps[i].Top()
			if t < possibleIdx {
				possibleIdx = t
				space = i
			}
		}
		// in case no empty space found
		if space == 10 || possibleIdx >= maxIdx {
			return -1, -1
		}
		val := heap.Pop(&spaceHeaps[space])
		return space, val.(int)
	}

	// building space
	idx := 0
	inputIdx := make([]int, len(inputs))
	t := len(inputs) - 1
	for i := 0; i < t; i++ {
		v := inputs[i]
		inputIdx[i] = idx
		idx += v

		i++
		v = inputs[i]
		inputIdx[i] = idx
		heap.Push(&spaceHeaps[v], idx)
		idx += v
	}
	inputIdx[len(inputs)-1] = idx

	result := 0
	// loop through inputs
	for i := len(inputs) - 1; i > -1; i -= 2 {
		v := inputs[i]
		currentIdx := inputIdx[i]
		// fmt.Println("current idx:", currentIdx, "current v:", v)
		space, idx := getNextEmptyIdx(v, currentIdx)
		if idx == -1 || space == -1 {
			idx = inputIdx[i]
		} else if space > v {
			newSpace := space - v
			newIdx := idx + v
			heap.Push(&spaceHeaps[newSpace], newIdx)
		}

		id := i >> 1
		for i := idx; i < idx+v; i++ {
			// fmt.Printf("placed id:%d at: %d\n", id, i)
			result += i * id
		}
	}

	fmt.Println("Result:", result)
}
