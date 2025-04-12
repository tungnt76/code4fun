// https://leetcode.com/problems/maximum-product-subarray/description/
package array

import "testing"

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	currentMax := nums[0]
	currentMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = max(nums[i], currentMax*nums[i])
		currentMin = min(nums[i], currentMin*nums[i])
		maxProduct = max(maxProduct, currentMax)
	}

	return maxProduct
}

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, -3, -1, 4}
	expected := 36
	result := maxProduct(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	nums = []int{-2, 0, -1}
	expected = 0
	result = maxProduct(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
