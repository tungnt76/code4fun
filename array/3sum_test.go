// https://leetcode.com/problems/3sum/description/
package array

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	// solution 1
	results := [][]int{}

	// for i, num := range nums {
	// 	new := append(nums[:i], nums[i+1:]...)
	// 	idx := twoSum(new, -num)
	// 	if len(idx) > 0 {
	// 		results = append(results, []int{num, new[idx[0]], new[idx[1]]})
	// 	}
	// }

	// return results

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// solution 2
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return results
}

func Test3Sum(t *testing.T) {
	// -4 -1 -1 0 1 2
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := threeSum(nums)

	if len(result) != len(expected) {
		fmt.Println(result)
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if len(result[i]) != len(expected[i]) {
			fmt.Println(result[i])
			t.Errorf("Expected length %d, got %d", len(expected[i]), len(result[i]))
		}
	}
}

func Test3Sum2(t *testing.T) {
	nums := []int{1, 0, 0, 0, 0}
	expected := [][]int{{0, 0, 0}}
	result := threeSum(nums)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if len(result[i]) != len(expected[i]) {
			fmt.Println(result[i])
			t.Errorf("Expected length %d, got %d", len(expected[i]), len(result[i]))
		}
	}
}
