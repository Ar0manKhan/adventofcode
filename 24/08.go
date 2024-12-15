package main

import (
	"fmt"
)

func extractRuneOccurance08(grid [][]rune) [][]Coordinates {
	result := make([][]Coordinates, 128)

	for i, r := range grid {
		for j, c := range r {
			cords := result[int(c)]
			cords = append(cords, Coordinates{i, j})
			result[int(c)] = cords
		}
	}

	return result
}

func Solve08v1() {
	grids := extractInputInByteGrid()
	occurance := extractRuneOccurance08(grids)
	maxX, maxY := len(grids), len(grids[0])

	resultGrid := createEmptyGrid(maxX, maxY, func() bool {
		return false
	})

	countPossibleAntinode := func(coords []Coordinates) int {
		count := 0
		for i, a := range coords {
			for j, b := range coords {
				if i == j {
					continue
				}

				dx := b.x - a.x
				dy := b.y - a.y

				nx := b.x + dx
				ny := b.y + dy

				if isUnderLimit(nx, ny, maxX, maxY) {
					// fmt.Println("Possible in:", nx, ny)
					resultGrid[nx][ny] = true
				}
			}
		}
		return count
	}

	for i, v := range occurance {
		if i != '.' && len(v) > 1 {
			countPossibleAntinode(v)
		}
	}

	result := CountOccurrences(resultGrid, true)

	fmt.Println("Result:", result)
}

func Solve08v2() {
	grids := extractInputInByteGrid()
	occurance := extractRuneOccurance08(grids)
	maxX, maxY := len(grids), len(grids[0])

	resultGrid := createEmptyGrid(maxX, maxY, func() bool {
		return false
	})

	countPossibleAntinode := func(coords []Coordinates) {
		for i, a := range coords {
			for j, b := range coords {
				if i == j {
					continue
				}
				// fmt.Println("=== checking for:", a, b)

				dx := b.x - a.x
				dy := b.y - a.y
				dGcd := gcd(abs(dx), abs(dy))
				dx /= dGcd
				dy /= dGcd

				nx, ny := b.x, b.y
				for isUnderLimit(nx, ny, maxX, maxY) {
					// fmt.Println("found:", nx, ny)
					resultGrid[nx][ny] = true
					nx += dx
					ny += dy
					// nx, ny = nx+dx, ny+dy
				}

			}
		}
	}

	for i, v := range occurance {
		if i != '.' && len(v) > 1 {
			countPossibleAntinode(v)
		}
	}

	result := CountOccurrences(resultGrid, true)

	fmt.Println("Result:", result)
}
