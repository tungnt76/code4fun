// https://leetcode.com/problems/combinations/description/

package recursion

import (
	"fmt"
	"testing"
)

func combine(n int, k int) [][]int {
	results := [][]int{}
	arr := []int{}
	collectN(0, n, k, arr, &results)

	return results
}

//
// {1,2} {1,}

// 4 2
// 1 2 3 4
// 1 2 3 4
// 1 2

func collectN(cur, n, k int, arr []int, results *[][]int) {
	fmt.Println(cur, n, k, arr, results)

	if len(arr) == k {
		result := make([]int, k)
		copy(result, arr)
		fmt.Println(arr, result)
		*results = append(*results, result)

		// arr = []int{}
		return
	}

	for i := cur; i < n; i++ {
		arr = append(arr, i+1)
		collectN(i+1, n, k, arr, results)
		arr = arr[:len(arr)-1]
	}
}

func TestCombination(t *testing.T) {
	n := 4
	k := 2

	results := combine(n, k)
	fmt.Println(results)
}
