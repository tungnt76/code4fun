// https://leetcode.com/problems/subsets/description/
package recursion

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func subsets(nums []int) [][]int {
	results := [][]int{}
	n := int(math.Pow(2, float64(len(nums))))

	for i := 0; i < n; i++ {
		elems := []int{}
		binary := strconv.FormatInt(int64(i), 2)
		for j := len(binary) - 1; j >= 0; j-- {
			if binary[j] == '1' {
				elems = append(elems, nums[len(binary)-1-j])
			}
		}

		results = append(results, elems)
	}

	return results
}

func subsetsBacktrack(nums []int) [][]int {
	results := [][]int{}
	arr := []int{}

	backtrackS(0, nums, arr, &results)
	return results
}

func backtrackS(cur int, nums []int, arr []int, results *[][]int) {
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	*results = append(*results, tmp)

	for i := cur; i < len(nums); i++ {
		arr = append(arr, nums[i])
		backtrackS(i+1, nums, arr, results)
		arr = arr[:len(arr)-1]
	}
}

func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsetsBacktrack([]int{1, 2, 3}))
}
