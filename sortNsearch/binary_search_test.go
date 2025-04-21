// https://leetcode.com/problems/binary-search/
package sortnsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4

	assert.Equal(t, search(nums, target), expected)
}
