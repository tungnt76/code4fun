// https://leetcode.com/problems/maximum-subarray/description/
package array

import "testing"

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// -2,-1,-3,4,3,5,6,1,5
// 6
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	result := maxSubArray(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
