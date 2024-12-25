package main

import "fmt"

type CoordinateSet map[Coordinates]bool

func calcRoute10(cord Coordinates, expected uint8, inputs [][]uint8, routeTable [][]*CoordinateSet, visited [][]bool, moveDirections []Coordinates, limit Coordinates) *CoordinateSet {
	if !isUnderLimit(cord.x, cord.y, limit.x, limit.y) {
		return nil
	}
	v := inputs[cord.x][cord.y]
	if v != expected {
		// fmt.Println("returning for bad value", v)
		return nil
	}
	visitedCell := visited[cord.x][cord.y]
	if visitedCell {
		// fmt.Printf("cord: %v, expected: %d, current: %d, returning memoized value: %d\n", cord, expected, v, routeTable[cord.x][cord.y])
		return routeTable[cord.x][cord.y]
	}
	visited[cord.x][cord.y] = true
	if v == 0 {
		// fmt.Printf("cord: %v, expected: %d, current: %d, returning default value \n", cord, expected, v)
		routeTable[cord.x][cord.y] = &CoordinateSet{cord: true}
		return routeTable[cord.x][cord.y]
	}
	routes := CoordinateSet{}
	for _, direction := range moveDirections {
		result := calcRoute10(Coordinates{x: cord.x + direction.x, y: cord.y + direction.y}, expected-1, inputs, routeTable, visited, moveDirections, limit)
		if result == nil {
			continue
		}
		for k := range *result {
			routes[k] = true
		}
	}
	// fmt.Println("for:", cord, "expected:", expected, "result:", result)
	routeTable[cord.x][cord.y] = &routes
	return &routes
}

func calcRoute10v2(cord Coordinates, expected uint8, inputs [][]uint8, routeTable [][]int, moveDirections []Coordinates, limit Coordinates) int {
	if !isUnderLimit(cord.x, cord.y, limit.x, limit.y) {
		return 0
	}
	v := inputs[cord.x][cord.y]
	if v != expected {
		// fmt.Println("returning for bad value", v)
		return 0
	}
	if routeTable[cord.x][cord.y] != -1 {
		// fmt.Printf("cord: %v, expected: %d, current: %d, returning memoized value: %d\n", cord, expected, v, routeTable[cord.x][cord.y])
		return routeTable[cord.x][cord.y]
	}
	if v == 0 {
		// fmt.Printf("cord: %v, expected: %d, current: %d, returning default value \n", cord, expected, v)
		routeTable[cord.x][cord.y] = 1
		return routeTable[cord.x][cord.y]
	}
	result := 0
	for _, direction := range moveDirections {
		result += calcRoute10v2(Coordinates{x: cord.x + direction.x, y: cord.y + direction.y}, expected-1, inputs, routeTable, moveDirections, limit)
	}
	// fmt.Println("for:", cord, "expected:", expected, "result:", result)
	routeTable[cord.x][cord.y] = result
	return result
}
func Solve10v1() {
	inputGridRune := extractInputInByteGrid()
	inputs := Convert2DSlice(inputGridRune, func(v rune) uint8 { return uint8(v - '0') })
	limit := Coordinates{len(inputs), len(inputs[0])}
	routeTable := createEmptyGrid(limit.x, limit.y, func() *CoordinateSet { return nil })
	visitedTable := createEmptyGrid(limit.x, limit.y, func() bool { return false })
	moveDirections := []Coordinates{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	result := 0
	for i, r := range inputs {
		for j, v := range r {
			if v == 9 {
				result += len(*calcRoute10(Coordinates{i, j}, 9, inputs, routeTable, visitedTable, moveDirections, limit))
			}
		}
	}
	fmt.Println("Result:", result)
}

func Solve10v2() {
	inputGridRune := extractInputInByteGrid()
	inputs := Convert2DSlice(inputGridRune, func(v rune) uint8 { return uint8(v - '0') })
	limit := Coordinates{len(inputs), len(inputs[0])}
	routeTable := createEmptyGrid(limit.x, limit.y, func() int { return -1 })
	moveDirections := []Coordinates{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	result := 0
	for i, r := range inputs {
		for j, v := range r {
			if v == 9 {
				result += calcRoute10v2(Coordinates{i, j}, 9, inputs, routeTable, moveDirections, limit)
			}
		}
	}
	fmt.Println("Result:", result)
}
