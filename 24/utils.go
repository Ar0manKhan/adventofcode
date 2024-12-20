package main

func makeEmptyBooleanGrid(row, col int) [][]bool {
	res := make([][]bool, row)
	for i := 0; i < row; i++ {
		res[i] = make([]bool, col)
	}
	return res
}

// AI generated function
func createEmptyGrid[T any](m, n int, generator func() T) [][]T {
	matrix := make([][]T, m)
	for i := 0; i < m; i++ {
		row := make([]T, n)
		for j := 0; j < n; j++ {
			row[j] = generator()
		}
		matrix[i] = row
	}
	return matrix
}

// DeepClone2DSlice clones a 2D slice with primary data types (e.g., int, float64, string).
// Chatgpt
func DeepClone2DSlice[T any](original [][]T) [][]T {
	if original == nil {
		return nil
	}
	clone := make([][]T, len(original))
	for i, row := range original {
		// Allocate a new slice for each row and copy the elements
		clone[i] = append([]T(nil), row...)
	}
	return clone
}

// AI generated
func CountOccurrences[T comparable](grid [][]T, value T) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == value {
				count++
			}
		}
	}
	return count
}

func gcd(a, b int) int {
	if b < a {
		a, b = b, a
	}
	rem := b % a
	switch rem {
	case 0:
		return a
	case 1:
		return 1
	default:
		return gcd(rem, a)
	}
}

func abs(a int) int {
	return max(a, -a)
}
