// https://leetcode.com/problems/product-of-array-except-self/description/
package array

import "testing"

// tags: prefixsum, array

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// 24 2 3 4
// 1 12 3 4
// 1 2 8 4
// 1 2 3 6

// left: 1 1 2 6
// right: 24 12 4 1

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		left[i+1] = left[i] * nums[i]
	}

	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	multiply := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		multiply[i] = left[i] * right[i]
	}

	return multiply
}

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	result := productExceptSelf(nums)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result[i])
		}
	}
}
