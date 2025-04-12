// https://leetcode.com/problems/sliding-window-maximum/
package array

import (
	"fmt"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	max := 0
	result := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		if i == 0 {
			max = nums[i]
			for j := i + 1; j < i+k; j++ {
				if max < nums[j] {
					max = nums[j]
				}
			}
			result = append(result, max)
			
			continue
		}else {
			if nums[i-1] == max {
				max = nums[i]
				for j := i; j < i+k; j++ {
					if max < nums[j] {
						max = nums[j]
					}
				}
			} else {
				if max < nums[i+k-1] {
					max = nums[i+k-1]
				}
			}
		}
		
		result = append(result, max)
	}

	return result
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expected := []int{3, 3, 5, 5, 6, 7}
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}
