package main

import "fmt"

func opred(matrix [][]float64) float64 {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0.0
	for col := 0; col < n; col++ {
		minor := make([][]float64, n-1)
		for i := range minor {
			minor[i] = make([]float64, n-1)
		}

		for i := 1; i < n; i++ {
			j2 := 0
			for j := 0; j < n; j++ {
				if j == col {
					continue
				}
				minor[i-1][j2] = matrix[i][j]
				j2++
			}
		}

		sign := 1.0
		if col%2 == 1 {
			sign = -1.0
		}
		det += sign * matrix[0][col] * opred(minor)
	}
	return det
}

func trace(matrix [][]float64) float64 {
	tr := 0.0
	for i := 0; i < len(matrix); i++ {
		tr += matrix[i][i]
	}
	return tr
}

func transpose(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\nОпределитель::", opred(matrix))

	fmt.Println("След:", trace(matrix))

	fmt.Println("\nТранспонированная матрица:")
	transposed := transpose(matrix)
	printMatrix(transposed)
}
