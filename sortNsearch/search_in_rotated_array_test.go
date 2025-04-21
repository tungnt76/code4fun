// https://leetcode.com/problems/search-in-rotated-sorted-array/
package sortnsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 4, 5, 6, 7, 0, 1, 2
// target = 0
// mid = 7
// left = 4, right = 2

func searchInSorted(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		// op2:
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func TestSearchInSorted(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1

	assert.Equal(t, expected, searchInSorted(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	expected = 4
	assert.Equal(t, expected, searchInSorted(nums, target))
}
