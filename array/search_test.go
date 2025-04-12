// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
package array

import (
	"testing"
)

// Search in Rotated Sorted Array

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// O(log n)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
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

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4
	result := search(nums, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	target = 3
	expected = -1
	result = search(nums, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
