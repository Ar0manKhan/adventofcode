package main

import (
	"fmt"
)

func getPuzzleByObstruction06(grid [][]rune) [][]bool {
	puzzle := make([][]bool, len(grid))

	for i, row := range grid {
		r := make([]bool, len(row))
		for j, col := range row {
			r[j] = col == '#'
		}
		puzzle[i] = r
	}
	return puzzle
}

func getGuardPosition06(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, col := range row {
			if col == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func Rotate06(x, y int) (int, int) {
	if x == -1 {
		return 0, 1
	} else if y == 1 {
		return 1, 0
	} else if x == 1 {
		return 0, -1
	} else {
		return -1, 0
	}
}

func CalculateTrueValueGrid(grid [][]bool) int {
	result := 0
	for _, r := range grid {
		for _, c := range r {
			if c {
				result++
			}
		}
	}
	return result
}

func Solve06v1() {
	result := 0

	grid := extractInputInByteGrid()
	// obstructions := getPuzzleByObstruction06(grid)
	gx, gy := getGuardPosition06(grid) // guard position
	gdx, gdy := -1, 0                  // guard movement
	maxX, maxY := len(grid), len(grid[0])
	visitedGrid := makeEmptyBooleanGrid(maxX, maxY)

	for {
		visitedGrid[gx][gy] = true
		nextX, nextY := gx+gdx, gy+gdy
		if !isUnderLimit(nextX, nextY, maxX, maxY) {
			break
		}

		if grid[nextX][nextY] == '#' {
			gdx, gdy = Rotate06(gdx, gdy)
		} else {
			gx, gy = nextX, nextY
		}
	}

	result = CalculateTrueValueGrid(visitedGrid)

	fmt.Println("Result:", result)
}

//	type Path6 struct {
//		actual, possible bool
//	}
type Coordinates struct {
	x, y int
}

// NOTE: not able to solve
func Solve06v2() {
	directions := []Coordinates{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	walkDirection := []uint8{1, 2, 4, 8}

	grid := extractInputInByteGrid()
	maxX, maxY := len(grid), len(grid[0])
	walkedGrid := createEmptyGrid(maxX, maxY, func() uint8 { return 0 })
	result := createEmptyGrid(maxX, maxY, func() bool { return false })

	isLooping := func(walkedGrid [][]uint8, start, extraObstacle Coordinates, gd int) bool {
		curr := start
		next := curr
		for {
			next = curr
			if walkedGrid[curr.x][curr.y]&walkDirection[gd] == 1 {
				result[extraObstacle.x][extraObstacle.y] = true
				return true
			}
			walkedGrid[curr.x][curr.y] |= walkDirection[gd]
			next.x += directions[gd].x
			next.y += directions[gd].y
			if !isUnderLimit(next.x, next.y, maxX, maxY) {
				break
			}
			if grid[next.x][next.y] == '#' || (next.x == extraObstacle.x && next.y == extraObstacle.y) {
				gd++
				gd %= 4
			} else {
				curr.x = next.x
				curr.y = next.y
			}
		}
		return false
	}

	gd := 0
	x, y := getGuardPosition06(grid)
	curr := Coordinates{x, y}
	next := curr
	for {
		next.x, next.y = curr.x+directions[gd].x, curr.y+directions[gd].y
		if !isUnderLimit(next.x, next.y, maxX, maxY) {
			break
		}
		if grid[next.x][next.y] == '#' {
			walkedGrid[curr.x][curr.y] |= walkDirection[gd]
			gd++
			gd %= 4
		} else {
			if !(next.x == x && next.y == y) && walkedGrid[next.x][next.y] == 0 {
				cloneWGrid := DeepClone2DSlice(walkedGrid)
				isLooping(cloneWGrid, curr, next, gd)
			}
			walkedGrid[curr.x][curr.y] |= walkDirection[gd]
			curr = next
		}
	}

	resCount := 0
	for _, r := range result {
		for _, c := range r {
			if c {
				resCount++
			}
		}
	}

	fmt.Println("Result:", resCount)
}

/*
func Solve06v2() {
	result := 0

	directions := []Coordinates{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	grid := extractInputInByteGrid()
	// obstructions := getPuzzleByObstruction06(grid)
	guard := Coordinates{}
	gx, gy := getGuardPosition06(grid) // guard position
	guard.x = gx
	guard.y = gy
	ig := guard
	gd := 0
	maxX, maxY := len(grid), len(grid[0])
	pathGrid := createEmptyGrid(
		maxX,
		maxY,
		func() []Path6 { return make([]Path6, 4) },
	)
	resultGrid := createEmptyGrid(maxX, maxY, func() bool { return false })

	markBackword := func() {
		current := guard
		movement := directions[gd]
		for {
			current.x, current.y = current.x-movement.x, current.y-movement.y
			if !isUnderLimit(current.x, current.y, maxX, maxY) || grid[current.x][current.y] == '#' {
				break
			}

			pathGrid[current.x][current.y][gd].possible = true
		}
	}

	markBackword()
	nextGd := (gd + 1) % 4
	for {
		pathGrid[guard.x][guard.y][gd].actual = true
		next := guard
		movement := directions[gd]

		next.x += movement.x
		next.y += movement.y

		if !isUnderLimit(next.x, next.y, maxX, maxY) {
			break
		}

		if grid[next.x][next.y] == '#' {
			// do something
			gd = nextGd
			nextGd = (gd + 1) % 4
			markBackword()
		} else {
			nextHasPath := false
			for _, v := range pathGrid[next.x][next.y] {
				nextHasPath = nextHasPath || v.actual
			}
			fmt.Println("checking for:", guard, pathGrid[guard.x][guard.y][nextGd], string(grid[guard.x][guard.y]), next, pathGrid[next.x][next.y], string(grid[next.x][next.y]))
			if (pathGrid[guard.x][guard.y][nextGd].actual || pathGrid[guard.x][guard.y][nextGd].possible) && !nextHasPath {
				fmt.Println("adding result")
				resultGrid[next.x][next.y] = true
			}
			guard = next
		}
	}

	for _, r := range resultGrid {
		for _, c := range r {
			if c {
				result++
			}
		}
	}

	if resultGrid[ig.x][ig.y] {
		result--
	}

	fmt.Println("Result:", result)
}
*/
