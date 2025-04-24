// https://leetcode.com/problems/spiral-matrix/description/
package matrix

import (
	"reflect"
	"testing"
)

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]

// 1 2 3 4
// 5 6 7 8
// 9 10 11 12
// 13 14 15 16

// 0-0|0-1|0-2|1-2|2-2|2-1|2-0

func spiralOrder(matrix [][]int) []int {
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	result := []int{}

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		for i := right; i >= left && top <= bottom; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top && left <= right; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	result := spiralOrder(matrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
