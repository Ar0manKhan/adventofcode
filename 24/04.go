package main

import (
	"fmt"
)

// NOTE: This is not the best code. I completely missed the usage of X and Y axis
// this code would be very confusing and it is my mistake

func getIndexPosition(i, x, y, movX, movY int) (int, int) {
	return x + (i * movX), y + (i * movY)
}

func isUnderLimit(x, y, maxX, maxY int) bool {
	return x > -1 && x < maxX && y > -1 && y < maxY
}

func checkXMas04(grid [][]rune, x, y, movX, movY, maxX, maxY int) bool {
	x1, y1 := getIndexPosition(0, x, y, movX, movY)
	x2, y2 := getIndexPosition(1, x, y, movX, movY)
	x3, y3 := getIndexPosition(2, x, y, movX, movY)
	x4, y4 := getIndexPosition(3, x, y, movX, movY)
	return isUnderLimit(x1, y1, maxX, maxY) && grid[x1][y1] == 'X' &&
		isUnderLimit(x2, y2, maxX, maxY) && grid[x2][y2] == 'M' &&
		isUnderLimit(x3, y3, maxX, maxY) && grid[x3][y3] == 'A' &&
		isUnderLimit(x4, y4, maxX, maxY) && grid[x4][y4] == 'S'
}

func Solve04v1() {
	result := 0
	grid := extractInputInByteGrid()

	h := len(grid)
	w := len(grid[0])

	for i, row := range grid {
		for j := range row {
			fmt.Println("--- checking for ", i, j)

			if checkXMas04(grid, i, j, -1, -1, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, 0, -1, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, 1, -1, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, -1, 0, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, 1, 0, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, -1, 1, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, 0, 1, w, h) {
				result++
			}

			if checkXMas04(grid, i, j, 1, 1, w, h) {
				result++
			}
		}
	}

	fmt.Println("result:", result)
}

func checkXMax04v2(grid [][]rune, y, x, maxY, maxX int) bool {
	tlX, tlY := x-1, y-1
	trX, trY := x+1, y-1
	dlX, dlY := x-1, y+1
	drX, drY := x+1, y+1
	tlValid := isUnderLimit(tlX, tlY, maxX, maxY)
	trValid := isUnderLimit(trX, trY, maxX, maxY)
	dlValid := isUnderLimit(dlX, dlY, maxX, maxY)
	drValid := isUnderLimit(drX, drY, maxX, maxY)

	if !(tlValid && trValid && dlValid && drValid) || grid[y][x] != 'A' {
		return false
	}

	tl, tr, dl, dr := grid[tlY][tlX], grid[trY][trX], grid[dlY][dlX], grid[drY][drX]

	result := 0

	if tl == 'M' && dr == 'S' {
		result++
	}

	if tl == 'S' && dr == 'M' {
		result++
	}

	if tr == 'M' && dl == 'S' {
		result++
	}

	if tr == 'S' && dl == 'M' {
		result++
	}

	return result == 2
}

func Solve04v2() {
	result := 0
	grid := extractInputInByteGrid()

	h := len(grid)
	w := len(grid[0])

	for i, row := range grid {
		for j := range row {
			if checkXMax04v2(grid, i, j, h, w) {
				result++
			}
		}
	}

	fmt.Println("result:", result)
}
