// https://leetcode.com/problems/set-matrix-zeroes/description/
package matrix

import (
	"fmt"
	"testing"
)

// op1: use hash map
// op2: use first element on each row and column to mark
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}

	setC := map[int]bool{}
	setR := map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				setR[i] = true
				setC[j] = true
			}
		}
	}

	for row := range setR {
		matrix[row] = make([]int, n)
	}

	for col := range setC {
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
