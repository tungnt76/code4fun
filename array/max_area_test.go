// https://leetcode.com/problems/container-with-most-water/description/
package array

import "testing"

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	result := maxArea(nums)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	nums = []int{1, 1}
	expected = 1
	result = maxArea(nums)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
