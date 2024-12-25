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

// Convert2DSlice converts a 2D slice of type T to a 2D slice of type U.
// AI generated
func Convert2DSlice[T, U any](input [][]T, convert func(T) U) [][]U {
	output := make([][]U, len(input))
	for i, row := range input {
		output[i] = make([]U, len(row))
		for j, val := range row {
			output[i][j] = convert(val)
		}
	}
	return output
}

// DisplayGrid displays a 2D grid with custom cell and row processing.
// `grid` is the 2D slice to display.
// `displayCell` is a function called for each cell.
// `endRow` is a function called at the end of each row.
func Display2DSlice[T any](grid [][]T, displayCell func(cell T), endRow func()) {
	for _, row := range grid {
		for _, cell := range row {
			displayCell(cell)
		}
		endRow()
	}
}
